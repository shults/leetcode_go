package t00022

import (
	"maps"
	"slices"
)

func generateParenthesis(n int) []string {
	cache := make(map[int][]string)
	return generateParenthesisInternal(n, cache)
}

func generateParenthesisInternal(
	n int,
	cache map[int][]string,
) []string {
	if cached, ok := cache[n]; ok {
		return cached
	}

	if n <= 0 {
		return nil
	}

	if n == 1 {
		cache[n] = []string{"()"}
		return cache[n]
	}

	uniqCases := make(map[string]struct{})
	for _, baseCase := range generateParenthesisInternal(n-1, cache) {
		generateFromBase(baseCase, uniqCases)
	}

	cache[n] = slices.Collect(maps.Keys(uniqCases))
	return cache[n]
}

func generateFromBase(base string, uniqCases map[string]struct{}) {
	for i, _ := range base {
		res := base[:i] + "()" + base[i:]
		uniqCases[res] = struct{}{}
	}

}
