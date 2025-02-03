package t00599_minimum_Index_sum_of_two_lists

import (
	"cmp"
	"slices"
)

type ind struct {
	first  int
	second int
	key    string
}

func compareIndex(a, b ind) int {
	return cmp.Compare(a.first+a.second, b.first+b.second)
}

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int, len(list1))

	for index, restaurant := range list1 {
		m[restaurant] = index
	}

	inds := make([]ind, 0)

	for second, restaurant := range list2 {
		first, ok := m[restaurant]

		if !ok {
			continue
		}

		inds = append(inds, ind{
			first:  first,
			second: second,
			key:    restaurant,
		})
	}

	slices.SortFunc(inds, compareIndex)

	if len(inds) == 0 {
		return []string{}
	}

	zero := inds[0]
	res := make([]string, 0)
	res = append(res, zero.key)

	for i := 1; i < len(inds); i++ {
		if compareIndex(inds[i], zero) != 0 {
			break
		}
		res = append(res, inds[i].key)
	}

	return res
}
