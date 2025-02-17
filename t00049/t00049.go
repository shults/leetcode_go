package t00049

type anagramKey [26]byte

func newAnagramKey(word string) anagramKey {
	var res anagramKey

	for _, b := range word {
		res[b-'a']++
	}

	return res
}

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
		items:    make(map[string]byte),
	}
}

func (a *anagram) addItem(word string) {
	cur, _ := a.items[word]
	a.items[word] = cur + 1
	a.capacity++
}

type anagram struct {
	items    map[string]byte
	capacity int
}

type anagramCollection struct {
	aMap map[anagramKey]*anagram
}

func newAnagramCollection() *anagramCollection {
	return &anagramCollection{
		aMap: make(map[anagramKey]*anagram),
	}
}

func (a *anagramCollection) findOrCreate(word string) *anagram {
	key := newAnagramKey(word)

	if anagramItem, ok := a.aMap[key]; ok {
		return anagramItem
	}

	anagramItem := newAnagram()
	a.aMap[key] = anagramItem
	return anagramItem
}

func (a *anagramCollection) prepareResult() [][]string {
	var res = make([][]string, 0, len(a.aMap))

	for _, aItem := range a.aMap {
		var aItems = make([]string, 0, aItem.capacity)

		for key, value := range aItem.items {
			for i := byte(0); i < value; i++ {
				aItems = append(aItems, key)
			}
		}

		res = append(res, aItems)

	}

	return res
}
