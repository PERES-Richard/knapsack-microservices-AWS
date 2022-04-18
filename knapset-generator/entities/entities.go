package entities

type Item struct {
	Id    int `json:"id"`
	Size  int `json:"size"`
	Value int `json:"value"`
}

type KnapSet struct {
	BagSize int    `json:"bagSize"`
	Items   []Item `json:"items"`
}
