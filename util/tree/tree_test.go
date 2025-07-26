package tree

import (
	"testing"
)

func TestBuildTreeOnlyIncludesSubTree(t *testing.T) {
	// 模拟表中有多个不相关的子树
	data := []T{
		// 第一棵树：以ID=1为根
		{Id: 1, Pid: 0}, // 根节点1
		{Id: 2, Pid: 1}, // 子节点
		{Id: 3, Pid: 1}, // 子节点
		{Id: 4, Pid: 2}, // 孙节点

		// 第二棵树：以ID=10为根（独立的树）
		{Id: 10, Pid: 0},  // 根节点2
		{Id: 11, Pid: 10}, // 子节点
		{Id: 12, Pid: 10}, // 子节点

		// 孤立节点
		{Id: 20, Pid: 999}, // 父节点不存在的孤立节点
	}

	// 测试1：构建以ID=1为根的树，应该只包含4个节点
	tree1, err := BuildTree(data, 1)
	if err != nil {
		t.Fatalf("构建树1失败: %v", err)
	}

	// 验证只包含预期的节点
	expectedNodes1 := map[int64]bool{1: true, 2: true, 3: true, 4: true}
	if len(tree1.NodeMap) != 4 {
		t.Errorf("树1期望包含4个节点，实际包含%d个", len(tree1.NodeMap))
	}

	for id := range tree1.NodeMap {
		if !expectedNodes1[id] {
			t.Errorf("树1包含了不应该存在的节点 ID: %d", id)
		}
	}

	// 测试2：构建以ID=10为根的树，应该只包含3个节点
	tree2, err := BuildTree(data, 10)
	if err != nil {
		t.Fatalf("构建树2失败: %v", err)
	}

	expectedNodes2 := map[int64]bool{10: true, 11: true, 12: true}
	if len(tree2.NodeMap) != 3 {
		t.Errorf("树2期望包含3个节点，实际包含%d个", len(tree2.NodeMap))
	}

	for id := range tree2.NodeMap {
		if !expectedNodes2[id] {
			t.Errorf("树2包含了不应该存在的节点 ID: %d", id)
		}
	}

	// 测试3：构建以孤立节点为根的树，应该只包含1个节点
	tree3, err := BuildTree(data, 20)
	if err != nil {
		t.Fatalf("构建树3失败: %v", err)
	}

	if len(tree3.NodeMap) != 1 {
		t.Errorf("树3期望包含1个节点，实际包含%d个", len(tree3.NodeMap))
	}

	if tree3.Root.Id != 20 {
		t.Errorf("树3根节点期望为20，实际为%d", tree3.Root.Id)
	}

	// 测试4：尝试构建不存在的根节点，应该失败
	_, err = BuildTree(data, 999)
	if err == nil {
		t.Error("期望构建不存在的根节点时返回错误")
	}

	// 测试5：构建以ID=2为根的树，应该只包含2个节点
	tree4, err := BuildTree(data, 2)
	if err != nil {
		t.Fatalf("构建树4失败: %v", err)
	}

	if len(tree4.NodeMap) != 2 {
		t.Errorf("树4期望包含2个节点，实际包含%d个", len(tree4.NodeMap))
	}

	expectedNodes4 := map[int64]bool{2: true, 4: true}
	for id := range tree4.NodeMap {
		if !expectedNodes4[id] {
			t.Errorf("树4包含了不应该存在的节点 ID: %d", id)
		}
	}
}

func TestBuildTreeStructure(t *testing.T) {
	data := []T{
		{Id: 1, Pid: 0},
		{Id: 2, Pid: 1},
		{Id: 3, Pid: 1},
		{Id: 4, Pid: 2},
		{Id: 5, Pid: 2},
	}

	tree, err := BuildTree(data, 1)
	if err != nil {
		t.Fatalf("构建树失败: %v", err)
	}

	// 验证树结构正确
	if tree.Root.Id != 1 {
		t.Errorf("根节点ID期望为1，实际为%d", tree.Root.Id)
	}

	if len(tree.Root.Children) != 2 {
		t.Errorf("根节点期望有2个子节点，实际有%d个", len(tree.Root.Children))
	}

	// 验证子节点
	node2 := tree.GetNode(2)
	if node2 == nil {
		t.Fatal("未找到节点2")
	}

	if len(node2.Children) != 2 {
		t.Errorf("节点2期望有2个子节点，实际有%d个", len(node2.Children))
	}

	node3 := tree.GetNode(3)
	if node3 == nil {
		t.Fatal("未找到节点3")
	}

	if len(node3.Children) != 0 {
		t.Errorf("节点3期望是叶子节点，实际有%d个子节点", len(node3.Children))
	}
}
