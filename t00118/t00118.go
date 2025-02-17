package t00118

import "math"

var gCache = newCache()

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)

	for i := 1; i <= numRows; i++ {
		res = append(res, generateRow(i, gCache))
	}

	return res
}

func generateRow(numRows int, cache *Cache) []int {
	if res, ok := cache.get(numRows); ok {
		return res
	}

	var lastRow = make([]int, numRows)
	lastRow[0] = 1
	lastRow[len(lastRow)-1] = 1

	if numRows > 2 {
		prev := generateRow(numRows-1, gCache)
		upto := int(math.Ceil(float64(numRows) / 2))

		for i := 1; i < upto; i++ {
			lastRow[i] = prev[i-1] + prev[i]
			lastRow[len(lastRow)-1-i] = lastRow[i]
		}
	}

	cache.put(numRows, lastRow)
	return lastRow
}

type Cache struct {
	data map[int][]int
}

func newCache() *Cache {
	return &Cache{
		data: make(map[int][]int, 30),
	}
}

func (c *Cache) get(row int) ([]int, bool) {
	res, ok := c.data[row]
	return res, ok
}

func (c *Cache) put(row int, val []int) {
	c.data[row] = val
}
