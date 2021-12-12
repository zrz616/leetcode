/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var sum int

func postOrder(root *TreeNode) {
    if root == nil {
        return
    }
    postOrder(root.Right)
    root.Val += sum
    sum = root.Val
    postOrder(root.Left)
}

func convertBST(root *TreeNode) *TreeNode {
    // 8 7  6  5  4  3  2  1  0
    // 8 15 21 26 30 33 35 36 36
    sum = 0
    postOrder(root)
    return root
}
