package star

func explore(list []root, pave []float64, dim int) int {
	min := 1
	var max int
	var next int
	var total int

	if pave[dim] <= points[dim+(list[1].id)*s] {
	  return 0
	}
  
	if list[0].id == 1 {
	  total = 1
	  next = list[1].id;
	  for i := dim; i < s; i++ {
		if pave[i] <= points[i+next*s] {
		  total = 0
		  break
		}
	  }
	} else {
	  total = 0
	  max = list[0].id
  
	  if dim == s-1 {
		if points[dim+(list[max].id)*s] < pave[dim] {
		  total = max
		} else {
		  for min <= max {
			next = ( min + max + 1 ) / 2
			if points[dim+(list[next].id)*s] < pave[dim]	{    
			  total = total + next - min + 1
			  min = next + 1
			} else {
			  max = next - 1
			}
		  }
		}
	  } else {
		for min <= max {
		  next = ( ( 1 + min ) * num2 + max * num ) / den
		  if points[dim+(list[next].id)*s] < pave[dim]	{
			if list[next].pt == nil {
			  list[next].pt = subtree(list, min, next, dim+1)
			}
			total = total + explore(list[next].pt, pave, dim+1)
			min = next + 1
		  } else {
			max = next - 1
		  }
		}
	  }
	}
	return total;
}

func fastexplore(pave []float64, rng int, maxsize []int, lexsize []int, lex [][]root, subtotal *int) int {
	var min int
	var max int
	var next int
	var start int
	size := lexsize[rng]
	var right int
	total := 0
	seuil := pave[rng]
	var refnode root
	var node []root

	if rng == s - 1 {
		for i := size - 1; 0 <= i; i-- {
			refnode = lex[rng][i]
			node = refnode.pt
			min = refnode.id
			max = node[0].id

			if max < min {
				lexsize[rng]--
				lex[rng][i] = lex[rng][lexsize[rng]]
				*subtotal += min - 1
			} else {
				total += min - 1
				right = 1
				for min <= max {
					next = (min + max + 1) / 2
					if points[rng+(node[next].id)*s] < seuil {
						total += next - min + 1
						min = next + 1
						if right == 1 {
							lex[rng][i].id = min
						}
					} else {
						right = 0
						max = next - 1
					}
				}
			}
		}
		total += *subtotal
	} else {
		*subtotal = 0
		lexsize[rng + 1] = 0
		for i := 0; i < size; i++ {
			refnode = lex[rng][i]
			node = refnode.pt
			start = refnode.id
			min = 1
			max = node[0].id
			for min != start {
				next = (( 1 + min ) * num2 + max * num) / den
				insertlex(node[next].pt, rng+1, maxsize, lexsize, lex)
				total += explore(node[next].pt, pave, rng + 1)
				min = next + 1
			}
			right = 1
			for min <= max {
				next = ( ( 1 + min ) * num2 + max * num ) / den
				if points[rng+(node[next].id)*s] < seuil {
				  if node[next].pt == nil {
					node[next].pt = subtree(node, min, next, rng+1)
				  }
				  insertlex(node[next].pt, rng+1, maxsize, lexsize, lex)
				  total = total + explore(node[next].pt, pave, rng+1)
				  min = next + 1
				  if right == 1 {
					if rng == 0 {
					  if &lex == &lexalpha {
						for j := lex[rng][i].id; j < next; j++ {
						  if node[j].pt != nil {
							freetree(node[j]);
						  }
						}
					  }
					}
					lex[rng][i].id = min
				  }
				} else {
				  right = 0
				  max = next - 1
				}
			}
		}
	}

	return total
}