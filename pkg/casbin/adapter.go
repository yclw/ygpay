package casbin

// copy from https://github.com/hailaz/gf-casbin-adapter/blob/main/adapter.go

import (
	"context"
	"runtime"
	"strings"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FieldName 字段名
type FieldName struct {
	Id    string
	PType string
	V     []string
	VLen  int
}

const (
	CASBINRULE_TABLE_NAME = "t_casbin_rule"
	CASBIN_V_LEN          = 6
)

var (
	// 判断是否实现了persist.*接口
	_ persist.Adapter      = new(Adapter)
	_ persist.BatchAdapter = new(Adapter)
	// _ persist.Dispatcher   = new(Adapter)

	DefaultFieldName = FieldName{
		Id:    "id",
		PType: "ptype",
		V:     []string{"v0", "v1", "v2", "v3", "v4", "v5"},
		VLen:  CASBIN_V_LEN,
	}
)

// Options 输入配置
type Options struct {
	Ctx       context.Context
	GDB       gdb.DB // gdb
	TableName string // 表名
	FieldName *FieldName
}

// Adapter represents the Xorm adapter for policy storage.
type Adapter struct {
	ctx       context.Context
	o         gdb.DB
	tableName string
	fieldName *FieldName
}

func NewAdapter(opts Options) *Adapter {
	fieldName := DefaultFieldName

	a := &Adapter{
		o:         opts.GDB,
		tableName: CASBINRULE_TABLE_NAME,
		fieldName: &fieldName,
	}

	if opts.Ctx != nil {
		a.ctx = opts.Ctx
	} else {
		a.ctx = context.Background()
	}

	if opts.TableName != "" {
		a.tableName = opts.TableName
	}

	a.SetFieldName(opts.FieldName)

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a
}

// NewAdapterWithTableName 设置表名
//
// createTime: 2022-03-04 17:04:32
//
// author: hailaz
func NewAdapterWithTableName(gdb gdb.DB, tableName string) *Adapter {
	return NewAdapter(Options{GDB: gdb, TableName: tableName})
}

// finalizer is the destructor for Adapter.
func finalizer(a *Adapter) {
}

func (a *Adapter) SetFieldName(fieldName *FieldName) {
	if fieldName == nil {
		return
	}
	if fieldName.Id != "" {
		a.fieldName.Id = fieldName.Id
	}
	if fieldName.PType != "" {
		a.fieldName.PType = fieldName.PType
	}
	vLen := len(fieldName.V)
	if vLen >= CASBIN_V_LEN {
		a.fieldName.VLen = vLen
		a.fieldName.V = make([]string, vLen)
		copy(a.fieldName.V, fieldName.V)
	}
}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	res, err := a.o.Model(a.tableName).All()
	if err != nil {
		return err
	}

	for _, line := range res {
		vList := make([]string, 0, a.fieldName.VLen+1)
		vList = append(vList, line[a.fieldName.PType].String())
		for index := 0; index < a.fieldName.VLen; index++ {
			vList = append(vList, line[a.fieldName.V[index]].String())
		}
		lineText := strings.Join(vList, ",")
		lineText = strings.TrimRight(lineText, ",")
		persist.LoadPolicyLine(lineText, model)
	}

	return nil
}

// toData description
func (a *Adapter) toData(ptype string, rule []string) g.Map {
	data := g.Map{
		a.fieldName.PType: ptype,
	}

	vLen := len(a.fieldName.V)

	for index := 0; index < len(rule); index++ {
		if index >= vLen {
			break
		}
		data[a.fieldName.V[index]] = rule[index]
	}

	return data
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) error {

	var lines g.List

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := a.toData(ptype, rule)
			lines = append(lines, line)
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := a.toData(ptype, rule)
			lines = append(lines, line)
		}
	}

	_, err := a.o.Ctx(a.ctx).Model(a.tableName).Data(lines).Insert()
	return err
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	_, err := a.o.Ctx(a.ctx).Model(a.tableName).Data(a.toData(ptype, rule)).Insert()
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	qs := a.o.Model(a.tableName).Safe()
	qs = qs.Where(a.fieldName.PType, ptype)
	for index := 0; index < len(rule); index++ {
		qs = qs.Where(a.fieldName.V[index], rule[index])
	}
	_, err := qs.Delete()
	return err

}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	qs := a.o.Model(a.tableName).Safe()
	qs = qs.Where(a.fieldName.PType, ptype)
	for index := 0; index <= CASBIN_V_LEN-1; index++ {
		if fieldIndex <= index && index < fieldIndex+len(fieldValues) {
			qs = qs.Where(a.fieldName.V[index], fieldValues[index-fieldIndex])
		}
	}
	_, err := qs.Delete()
	return err
}

func (a *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	var lines g.List
	for _, rule := range rules {
		line := a.toData(ptype, rule)
		lines = append(lines, line)
	}
	_, err := a.o.Ctx(a.ctx).Model(a.tableName).Data(lines).Insert()
	return err
}

func (a *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {
	return a.RemovePoliciesCtx(a.ctx, sec, ptype, rules)
}

// RemovePoliciesCtx removes multiple policy rules from the storage.
func (a *Adapter) RemovePoliciesCtx(ctx context.Context, sec string, ptype string, rules [][]string) error {
	return a.o.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, rule := range rules {
			_, err := tx.Model(a.tableName).Safe().Ctx(ctx).Where(a.toData(ptype, rule)).Delete()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// // ClearPolicy clears all current policy in all instances
// func (a *Adapter) ClearPolicy() error {
// 	_, err := a.o.Model(a.tableName).WhereGT("id", 0).Delete()
// 	return err
// }

// // UpdatePolicy updates policy rule from all instance.
// func (a *Adapter) UpdatePolicy(sec string, ptype string, oldRule, newRule []string) error {
// 	return nil
// }

// // UpdatePolicies updates some policy rules from all instance
// func (a *Adapter) UpdatePolicies(sec string, ptype string, oldrules, newRules [][]string) error {
// 	return nil
// }

// // UpdateFilteredPolicies deletes old rules and adds new rules.
// func (a *Adapter) UpdateFilteredPolicies(sec string, ptype string, oldRules [][]string, newRules [][]string) error {
// 	return nil
// }
