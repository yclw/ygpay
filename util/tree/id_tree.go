package tree

import (
	"fmt"
)

// T 基础的ID-PID结构，用于接收表数据
type T struct {
	Id  int64 `json:"id"`
	Pid int64 `json:"pid"`
}

// TreeNode 树节点结构
type TreeNode struct {
	Id       int64       `json:"id"`                 // 节点ID
	Title    string      `json:"title,omitempty"`    // 节点标题
	Pid      int64       `json:"pid"`                // 父节点ID
	Children []*TreeNode `json:"children,omitempty"` // 子节点列表
	Data     interface{} `json:"-"`                  // 附加数据，可以存储完整的业务对象
}

// IdTree ID树结构
type IdTree struct {
	Root    *TreeNode           `json:"root"`    // 根节点
	NodeMap map[int64]*TreeNode `json:"nodeMap"` // 节点映射表，用于快速查找
	RootId  int64               `json:"rootId"`  // 根节点ID
}

// NewIdTree 创建新的ID树
func NewIdTree(rootId int64) *IdTree {
	return &IdTree{
		NodeMap: make(map[int64]*TreeNode),
		RootId:  rootId,
	}
}

// BuildTree 构建树结构
func BuildTree(data []T, rootId int64) (*IdTree, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("数据为空")
	}

	// 第一步：创建所有原始数据的映射表
	allDataMap := make(map[int64]T)
	for _, item := range data {
		allDataMap[item.Id] = item
	}

	// 第二步：检查根节点是否存在
	rootData, exists := allDataMap[rootId]
	if !exists {
		return nil, fmt.Errorf("未找到根节点 ID: %d", rootId)
	}

	tree := NewIdTree(rootId)

	// 第三步：从根节点开始递归构建子树，只包含属于该子树的节点
	tree.Root = tree.buildSubTree(rootData, allDataMap)

	return tree, nil
}

// buildSubTree 递归构建子树
func (tree *IdTree) buildSubTree(nodeData T, allDataMap map[int64]T) *TreeNode {
	// 创建当前节点
	node := &TreeNode{
		Id:       nodeData.Id,
		Pid:      nodeData.Pid,
		Children: make([]*TreeNode, 0),
		Data:     nodeData,
	}

	// 将节点添加到映射表中
	tree.NodeMap[nodeData.Id] = node

	// 查找并递归构建所有子节点
	for _, item := range allDataMap {
		if item.Pid == nodeData.Id {
			childNode := tree.buildSubTree(item, allDataMap)
			node.Children = append(node.Children, childNode)
		}
	}

	return node
}

// BuildTreeFromNode 从指定节点构建新的子树
func BuildTreeFromNode(rootNode *TreeNode) *IdTree {
	if rootNode == nil {
		return nil
	}

	// 创建新的树实例
	newTree := NewIdTree(rootNode.Id)
	newTree.Root = rootNode

	// 遍历节点并添加到映射表
	newTree.addToNodeMap(rootNode)

	return newTree
}

// addToNodeMap 递归添加节点到映射表
func (tree *IdTree) addToNodeMap(node *TreeNode) {
	tree.NodeMap[node.Id] = node
	for _, child := range node.Children {
		tree.addToNodeMap(child)
	}
}

// GetNode 根据ID获取节点
func (tree *IdTree) GetNode(id int64) *TreeNode {
	return tree.NodeMap[id]
}

// GetChildren 获取指定节点的所有子节点
func (tree *IdTree) GetChildren(id int64) []*TreeNode {
	if node := tree.GetNode(id); node != nil {
		return node.Children
	}
	return nil
}

// GetAllNodes 获取树中的所有节点（平铺）
func (tree *IdTree) GetAllNodes() []*TreeNode {
	nodes := make([]*TreeNode, 0, len(tree.NodeMap))
	for _, node := range tree.NodeMap {
		nodes = append(nodes, node)
	}
	return nodes
}

// GetPath 获取从根节点到指定节点的路径
func (tree *IdTree) GetPath(targetId int64) []*TreeNode {
	path := make([]*TreeNode, 0)
	node := tree.GetNode(targetId)

	// 从目标节点向上追溯到根节点
	for node != nil && node.Id != tree.RootId {
		path = append([]*TreeNode{node}, path...) // 在开头插入
		node = tree.GetNode(node.Pid)
	}

	// 添加根节点
	if node != nil && node.Id == tree.RootId {
		path = append([]*TreeNode{node}, path...)
	}

	return path
}

// ToFlat 将树结构转换为平铺列表
func (tree *IdTree) ToFlat() []T {
	result := make([]T, 0, len(tree.NodeMap))
	for _, node := range tree.NodeMap {
		result = append(result, T{
			Id:  node.Id,
			Pid: node.Pid,
		})
	}
	return result
}
