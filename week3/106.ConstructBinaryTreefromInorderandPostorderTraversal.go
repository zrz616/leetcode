func buildTree(inorder []int, postorder []int) *TreeNode {
	r1 := len(inorder) - 1
	r2 := len(postorder) - 1
	var buildSubTree func(l1 int, r1 int, l2 int, r2 int) *TreeNode
	buildSubTree = func(l1 int, r1 int, l2 int, r2 int) *TreeNode {
		if l2 > r2 {
			return nil
		}
		head := &TreeNode{Val: postorder[r2]}
		//fmt.Println(postorder[r2])
		mid := l1
		for inorder[mid] != head.Val {
			mid++
		}
		// inorder left [l1, mid - 1]
		// inorder right [mid + 1, r1]
		// postorder left [l2, l2 + mid - l1]
		// postorder right [l2 - mid - l1, r2 - 1]
		head.Left = buildSubTree(l1, mid-1, l2, l2+mid-l1-1)
		head.Right = buildSubTree(mid+1, r1, l2+mid-l1, r2-1)
		return head
	}
	return buildSubTree(0, r1, 0, r2)
}
