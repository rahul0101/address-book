package impl

type IndexMap struct {
	contactIds map[string][]int
}

func NewIndexMap() *IndexMap {
	return &IndexMap{contactIds: make(map[string][]int)}
}

func (im *IndexMap) Insert(word string, id int) {
	im.contactIds[word] = append(im.contactIds[word], id)
}

func (im *IndexMap) Search(word string) []int {
	contactIds, ok := im.contactIds[word]
	if !ok {
		return nil
	}
	return contactIds
}
