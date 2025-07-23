// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Attachment is the golang structure for table attachment.
type Attachment struct {
	Id        int64       `json:"id"        orm:"id"         ` // 文件ID
	AppId     string      `json:"appId"     orm:"app_id"     ` // 应用ID
	MemberId  int64       `json:"memberId"  orm:"member_id"  ` // 管理员ID
	CateId    uint64      `json:"cateId"    orm:"cate_id"    ` // 上传分类
	Drive     string      `json:"drive"     orm:"drive"      ` // 上传驱动
	Name      string      `json:"name"      orm:"name"       ` // 文件原始名
	Kind      string      `json:"kind"      orm:"kind"       ` // 上传类型
	MimeType  string      `json:"mimeType"  orm:"mime_type"  ` // 扩展类型
	NaiveType string      `json:"naiveType" orm:"naive_type" ` // NaiveUI类型
	Path      string      `json:"path"      orm:"path"       ` // 本地路径
	FileUrl   string      `json:"fileUrl"   orm:"file_url"   ` // url
	Size      int64       `json:"size"      orm:"size"       ` // 文件大小
	Ext       string      `json:"ext"       orm:"ext"        ` // 扩展名
	Md5       string      `json:"md5"       orm:"md5"        ` // md5校验码
	Status    int         `json:"status"    orm:"status"     ` // 状态
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 修改时间
}
