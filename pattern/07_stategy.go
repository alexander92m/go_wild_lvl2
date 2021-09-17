package main

import "fmt"


type cache struct {
    storage      map[string]string
    evictionAlgo evictionAlgo
    capacity     int
    maxCapacity  int
}

type evictionAlgo interface {
    evict(c *cache)
}

type fifo struct {
}

type lfu struct {
}

type lru struct {
}

func main() {
    lfu := &lfu{}
    cache := initCache(lfu)

    cache.add("a", "1")
    cache.add("b", "2")

    cache.add("c", "3")

    lru := &lru{}
    cache.setEvictionAlgo(lru)

    cache.add("d", "4")

    fifo := &fifo{}
    cache.setEvictionAlgo(fifo)

    cache.add("e", "5")
    
}

//initCache создание кэша
func initCache(e evictionAlgo) *cache {
    storage := make(map[string]string)
    return &cache{
        storage:      storage,
        evictionAlgo: e,
        capacity:     0,
        maxCapacity:  2,
    }
}

//setEvictionAlgo выбор алгоритма управления кешем
func (c *cache) setEvictionAlgo(e evictionAlgo) {
    c.evictionAlgo = e
}

//add добавить
func (c *cache) add(key, value string) {
    if c.capacity == c.maxCapacity {
        c.evict()
		fmt.Println(key, value, c.capacity)
    }
    c.capacity++
    c.storage[key] = value
}

//get получить
func (c *cache) get(key string) {
    delete(c.storage, key)
}

//evict освободить аким либо алгоритмом
func (c *cache) evict() {
    c.evictionAlgo.evict(c)
    c.capacity--
}

//evict освободить алоритмом fifo
func (l *fifo) evict(c *cache) {
    fmt.Println("Evicting by fifo strtegy")
}

//evict освободить алоритмом lfu
func (l *lfu) evict(c *cache) {
    fmt.Println("Evicting by lfu strtegy")
}

//evict освободить алоритмом lru
func (l *lru) evict(c *cache) {
    fmt.Println("Evicting by lru strtegy")
}