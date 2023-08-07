package indexDataStructures

type IIndexDS interface {
	Insert(word string, id int)
	Search(word string) []int
}
