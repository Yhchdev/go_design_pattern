package main

import "fmt"

type IPizza interface {
	getPrice() int
}

// 基类: 素食披萨
type Vegetable struct {
}

func (v Vegetable) getPrice() int {
	return 10
}

// 装饰器1: 奶酪装饰器
type Cheese struct {
	pizza IPizza
}

func (c Cheese) getPrice() int {
	return c.pizza.getPrice() + 3
}

// 装饰器2:
type Tomato struct {
	pizza IPizza
}

func (c Tomato) getPrice() int {
	return c.pizza.getPrice() + 4
}

func main() {

	vegetablePizza := Vegetable{}

	cheeseVegePizza := Cheese{vegetablePizza}

	tomatoCheeseVegePizza := Tomato{cheeseVegePizza}

	fmt.Printf("加了番茄和奶酪的披萨最终价格:%d\n", tomatoCheeseVegePizza.getPrice())

}


// output
// 加了番茄和奶酪的披萨最终价格:17