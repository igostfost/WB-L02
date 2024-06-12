package main

import "fmt"

/*
Стратегия — это поведенческий паттерн,
выносит набор алгоритмов в собственные классы и делает их взаимозаменимыми.

Другие объекты содержат ссылку на объект-стратегию и делегируют ей работу.
Программа может подменить этот объект другим, если требуется иной способ решения задачи.
*/

// EvictionAlgorithm - интерфейс алгоритма освобождения кеша (стратегии)
type EvictionAlgorithm interface {
	evict(c *Cache)
}

type Fifo struct {
}

// Убирает запись, которая была создана раньше остальных
func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strategy")
}

// Убирает запись, которая использовалась наиболее давно.
type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
}

// Убирает запись, которая использовалась наименее часто.
type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strategy")
}

type Cache struct {
	storage      map[string]int
	evictionAlgo EvictionAlgorithm
	capacity     int
	maxCapacity  int
}

func initCache(evictionAlgo EvictionAlgorithm) *Cache {
	return &Cache{
		storage:      make(map[string]int),
		evictionAlgo: evictionAlgo,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvicAlgorithm(newEvictionAlgo EvictionAlgorithm) {
	c.evictionAlgo = newEvictionAlgo
}

func (c *Cache) add(key string, value int) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.storage[key] = value
	c.capacity++
}

func (c *Cache) del(key string) {
	delete(c.storage, key)
	c.capacity--
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// Клиентский код
func main() {
	lruStrategy := Lru{}
	cache := initCache(&lruStrategy)

	cache.add("key1", 1)
	cache.add("key2", 2)
	cache.add("key3", 3)

	// меняем стратегию
	fifoStrategy := Fifo{}
	cache.setEvicAlgorithm(&fifoStrategy)
	cache.add("key4", 4)

	// меняем стратегию
	lfuStrategy := Lfu{}
	cache.setEvicAlgorithm(&lfuStrategy)
	cache.add("key5", 5)

}
