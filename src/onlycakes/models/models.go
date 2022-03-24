package models

type Option struct {
	Id    uint8
	Name  string
	Price float32
}

type Item struct {
	CakeType      Option
	CakeFlavor    Option
	ToppingFlavor Option
	Total         float32
}
