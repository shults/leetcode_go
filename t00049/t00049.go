package t00049

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	col := newAnagramCollection()

	for _, str := range strs {
		aItem := col.findOrCreate(str)
		aItem.addItem(str)
	}

	return col.prepareResult()
}

func newAnagram() *anagram {
	return &anagram{
		capacity: 0,
		items:    make(map[string]int),
	}
}

func (a *anagram) addItem(word string) {
	cur, _ := a.items[word]
	a.items[word] = cur + 1
	a.capacity++
}

type anagram struct {
	items    map[string]int
	capacity int
}

type anagramCollection struct {
	aMap map[string]*anagram
}

func newAnagramCollection() *anagramCollection {
	return &anagramCollection{
		aMap: make(map[string]*anagram),
	}
}

func (a *anagramCollection) findOrCreate(word string) *anagram {
	bWord := []byte(word)
	slices.Sort(bWord)

	if anagramItem, ok := a.aMap[string(bWord)]; ok {
		return anagramItem
	}

	anagramItem := newAnagram()

	a.aMap[string(bWord)] = anagramItem

	return anagramItem
}

func (a *anagramCollection) prepareResult() [][]string {
	var res = make([][]string, 0, len(a.aMap))

	for _, aItem := range a.aMap {
		var aItems = make([]string, 0, aItem.capacity)

		for key, value := range aItem.items {
			for i := 0; i < value; i++ {
				aItems = append(aItems, key)
			}
		}

		res = append(res, aItems)

	}

	return res
}
