package main

import (
	"fmt"

	cache "github.com/hitesh22rana/speedy-cache/lib/cache"
)

func main() {
	cache := cache.NewLRUCache(2)
	cache.Set("a", "1")
	cache.Set("b", 2)

	val1, _ := cache.Get("a")
	fmt.Println(val1)

	val2, _ := cache.Get("b")
	fmt.Println(val2)

	cache.Set("c", 3)
	val3, _ := cache.Get("c")
	fmt.Println(val3)

	val4, _ := cache.Get("b")
	fmt.Println(val4)

	cache.Delete("b")

	val5, _ := cache.Get("b")
	fmt.Println(val5)

	err := cache.Delete("b")
	fmt.Println(err)
}
