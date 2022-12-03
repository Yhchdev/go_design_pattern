package main

import "fmt"

type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func NewItem(name string) Item {
	return Item{name: name}
}

func (i *Item) register(observer Observer) {
	i.observerList = append(i.observerList, observer)
}

func (i *Item) deregister(observer Observer) {
	// todo remove observer from observerList
}

func (i *Item) notifyAll() {
	for _, v := range i.observerList {
		v.update(i.name)
	}
}

func (i Item) updateAvailability() {
	fmt.Printf("item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

type Observer interface {
	update(string)
}

type Customs struct {
	id string
}

func (c Customs) update(name string) {
	fmt.Printf("send email to customer %s for item %s\n", c.id, name)
}

func main() {
	book := NewItem("《设计模式：可复用面向对象软件的基础》")

	book.register(&Customs{id: "a@qq.com"})
	book.register(&Customs{id: "b@qq.com"})

	book.updateAvailability()

}
