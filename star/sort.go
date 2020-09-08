package star

func quicksort(list []root, dim int, l int, r int) {
	i := l
	j := r + 1
	var tmp int
	pivot := points[dim + (list[l].id) * s]

	for i < j {
		for {
			i++
			if !(i < r && points[dim + (list[i].id) * s] < pivot) {
				break
			}
		}

		for {
			j--
			if !(pivot < points[dim + (list[j].id) * s]) {
				break
			}
		}

		if i < j {
			tmp = list[i].id
			list[i].id = list[j].id
			list[j].id = tmp
		}
	}

	tmp = list[l].id
	list[l].id = list[j].id
	list[j].id = tmp

	if l < (j - 1) {
		quicksort(list, dim, l, j - 1)
	}

	if (j + 1) < r {
		quicksort(list, dim, j+1, r)
	}
}