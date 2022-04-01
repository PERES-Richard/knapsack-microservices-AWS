package entities

type Item struct {
	Id    int
	Size  int
	Value int
}

type KnapSet struct {
	BagSize int
	Items   []Item
}
