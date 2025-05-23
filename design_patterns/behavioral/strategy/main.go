/* Strategy is a behavioral design pattern that turns a set of behaviors into objects and makes them interchangeable inside original context object. */

package main

func main() {
    lfu := &Lfu{}
    cache := initCache(lfu)

    cache.add("a", "1")
    cache.add("b", "2")

    cache.add("c", "3")

    lru := &Lru{}
    cache.setEvictionAlgo(lru)

    cache.add("d", "4")

    fifo := &Fifo{}
    cache.setEvictionAlgo(fifo)

    cache.add("e", "5")

}