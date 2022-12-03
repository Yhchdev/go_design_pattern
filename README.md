# go_design_pattern
go设计模式

### 结构型模式
#### 1. 结构型模式解决什么问题
结构模式关注类和对象的组合，解决如何将类和对象组装成较大结构的同时，保持结构的灵活和可复用性


#### 2.策略模式
策略模式将一组行为分别封装成不同对象，使得这些对象可以根据需要任意替换，而不影响原有代码的逻辑流程。其本质是通过接口，解耦行为和调用该行为的对象。
最常见的应用场景就是缓存库的设计，需要根据实际需要自行选择和灵活替换缓存淘汰算法（常见的有LRU,FIFO,LFU），甚至自定义缓存淘汰算法，而我们只需要实现缓存淘汰算法所规定的方法
就可以替换不用的策略，需要注意的是为了方便替换，我们往往会设置下属代码中诸如`setEvictStrategy()`的方法进行策略替换。

```go
package main

import "fmt"

// 策略接口
type EvictionAlgorithm interface {
	evict()
}

type Lru struct {
}

// lru 的实现
func (c Lru) evict() {
	fmt.Println("evicting by lru strategy")
}

type Fifo struct {
}

// fifo 的实现
func (c Fifo) evict() {
	fmt.Println("evicting by fifo strategy")
}

type Cache struct {
	storage           map[string]string
	evictionAlgorithm EvictionAlgorithm
	capacity          int
	maxCapacity       int
}

func NewCache(e EvictionAlgorithm) *Cache {
	return &Cache{
		storage:           make(map[string]string, 0),
		evictionAlgorithm: e,
		capacity:          0,
		maxCapacity:       0,
	}
}

// 设置策略
func (c *Cache) setEvictStrategy(e EvictionAlgorithm) {
	c.evictionAlgorithm = e
}

func (c *Cache) evict(){
	c.evictionAlgorithm.evict()
	c.capacity--
}

func (c *Cache) Add(k, v string) {
	if c.capacity >= c.maxCapacity {
		c.evict()
	}

	c.storage[k] = v
	c.capacity++
}

func main() {
	CacheClient := NewCache(nil)

	lurStrategy := Lru{}
	CacheClient.setEvictStrategy(lurStrategy)

	CacheClient.Add("a","1")
	CacheClient.Add("b","2")


	fifoStrategy := Fifo{}
	CacheClient.setEvictStrategy(fifoStrategy)

	CacheClient.Add("a","1")
	CacheClient.Add("b","2")


}

```
#### 3.模版模式
模版模式在基类中定义了一系列的逻辑（算法，业务逻辑）的框架，可以通过子类重写逻辑的特定步骤，而不修改原有结构。其本质是将公共的方法放到抽象类中，而通同接口将
不能通用的方法定义为接口，让实现了接口的子类去实现这部分差异的方法。下面代码演示了一种场景，给用户发送验证码，可以有短信和邮件两种方式，而在这之前的
业务逻辑是共用的，例如第一步生成验证码，第二步保存验证码...,因此像这种操作步骤的流程是相同的只是某几个实现方式不同的场景就可以使用模版模式。
```go
package main

import "fmt"

type IVerificationCode interface {
	genCode() string
	saveCode(code string) error
	getMsg(code string) string
	sendMsg(msg string) error
}

func genAndSendCode(opt IVerificationCode) error {
	code := opt.genCode()

	if err := opt.saveCode(code); err != nil {
		return err
	}
	msg := opt.getMsg(code)

	if err := opt.sendMsg(msg); err != nil {
		return err
	}

	return nil
}

// 公共部分抽象
type CodePart struct {
}

func (c CodePart) genCode() string {
	return "1234"
}

func (c CodePart) getMsg(code string) string {
	return fmt.Sprintf("你的验证码是：%s", code)

}

func (c CodePart) saveCode(code string) error {
	fmt.Printf("服务端保存了验证码：%s\n", code)
	return nil
}

type SmsCode struct {
	CodePart
}

func (s SmsCode) sendMsg(code string) error {
	fmt.Println("通过短信的方式发送了验证码")
	return nil
}

type EmailCode struct {
	CodePart
}

func (e EmailCode) sendMsg(code string) error {
	fmt.Println("通过邮件的方式发送了验证码")
	return nil
}

func main() {

	sms := SmsCode{}

	if err := genAndSendCode(sms); err != nil {
		return
	}

	email := EmailCode{}

	if err := genAndSendCode(email); err != nil {
		return
	}

}

```

#### 4.观察者模式
观察者模式提供了一种把一个对象其状态的变更，通知给实现了订阅者接口的对象(观察者)的机制，同时其他对象(观察者)可以此对象(被观察的对象)进行订阅和取消订阅。其本质是通过接口解耦通知对象和
被通知对象这种一对多的关系，使得通知对象和接口，接口和被通知对象简化为简单的一对一关系。
下面的代码实现了当某个商品有库存的时候，通知订阅了这个商品上架提醒的顾客
```go
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

```

#### 5.总结
下面是分别是这3种设计模式的应用场景:

| 设计模式  | 常见应用场景                         |
|-------|--------------------------------|
| 策略模式  | 按照实际需求要对系统的算法做任意替换，而不影响原有代码    |
| 模版模式  | 固定的流程和逻辑，但不同对象在某些步骤上的实现方式有差别   |
| 观察者模式 | 一个对象(被观察者)需要将其状态的变化通知其他对象(观察者) |
