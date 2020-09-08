package star

func initlex() {
	maxsizealpha = make([]int, s)

	for i := 0; i < s; i++ {
		maxsizealpha[i] = 1
	}

	lexsizealpha = make([]int, s)
	lexsizealpha[0] = 1
	lexalpha = make([][]root, s)

	for i := 0; i < s; i++ {
		lexalpha[i] = make([]root, maxsizealpha[i])
	}

	lexalpha[0][0].id = 1
	lexalpha[0][0].pt = suptree

	maxsizebeta = make([]int, s)

	for i := 0; i < s; i++ {
		maxsizebeta[i] = 1
	}

	lexsizebeta = make([]int, s)
	lexsizebeta[0] = 1
	lexbeta = make([][]root, s)

	for i := 0; i < s; i++ {
		lexbeta[i] = make([]root, maxsizebeta[i])
	}

	lexbeta[0][0].id = 1
	lexbeta[0][0].pt = suptree
}

func insertlex(node []root, rng int, maxsize []int, lexsize []int, lex [][]root) {
	i := lexsize[rng]
	if i == maxsize[rng] {
		maxsize[rng] *= 2
		new_lex := make([]root, maxsize[rng])
		for i := 0; i < len(lex[rng]); i++ {
			new_lex[i] = lex[rng][i]
		}
		lex[rng] = new_lex
	}

	lex[rng][i].pt = node
	lex[rng][i].id = 1
	lexsize[rng] = i + 1
}

func subtree(list []root, min int, next int, dim int) []root {
	var aux int
	var n2 int
	var newtree []root

	aux = min - 1
	n2 = next - min + 1
	newtree = make([]root, n2 + 1)

	for i := 1; i <= n2; i++ {
		newtree[i].id = list[i + aux].id
	}

	newtree[0].id = n2

	if 1 < n2 {
		quicksort(newtree, dim, 1, n2)
	}

	return newtree
}