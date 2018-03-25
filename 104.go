package main

import "fmt"

/*
104. 二叉树的最大深度

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶节点的最长路径上的节点数。

案例：
给出二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回最大深度为 3 。


思路： 我采用中序遍历的方式，然后递归就可以了，然后注意几个判断为空的地方就 ok 了
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{1, &TreeNode{2, nil, nil}, nil}
	maxDepth := maxDepth(&root)
	fmt.Println("maxDepth:", maxDepth)
}

func maxDepth(root *TreeNode) int {
	maxDepth := getDepth(root, 0)
	return maxDepth
}

func getDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth = depth + 1
	if root.Left == nil && root.Right == nil {
		return depth
	}

	leftDepth := getDepth(root.Left, depth)
	rithDepth := getDepth(root.Right, depth)
	if leftDepth > rithDepth {
		return leftDepth
	} else {
		return rithDepth
	}
}

