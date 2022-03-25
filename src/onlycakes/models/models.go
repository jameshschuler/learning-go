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

type Order struct {
	Items []Item
	Total float32
}

func (o *Order) AddItem(item Item) {
	o.Items = append(o.Items, item)

	o.calcOrderTotal()
}

func (o *Order) calcOrderTotal() {
	for _, item := range o.Items {
		o.Total += item.Total
	}
}
