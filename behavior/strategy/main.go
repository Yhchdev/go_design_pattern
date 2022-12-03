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
