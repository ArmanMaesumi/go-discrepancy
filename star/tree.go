package star

func supertree() {
	suptree = make([]root, n + 1)

	for i := 1; i <= n; i++ {
		suptree[i].id = i - 1
	}

	suptree[0].id = n
	quicksort(suptree, 0, 1, n)
}

func freetree(node root) {
	max := node.pt[0].id

	for i := 1; i <= max; i++ {
		if node.pt[i].pt != nil {
			freetree(node.pt[i])
		}
	}

	node.pt = nil
}