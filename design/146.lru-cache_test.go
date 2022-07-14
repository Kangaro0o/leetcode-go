package design

import (
	"fmt"
	"testing"
)

func TestConstructor1(t *testing.T) {
	cache := Constructor1(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	v1 := cache.Get(1)
	fmt.Println("v1", v1)
	cache.Put(3, 3)
	v2 := cache.Get(2)
	fmt.Println("v2", v2)
	cache.Put(4, 4)
	v3 := cache.Get(1)
	fmt.Println("v3", v3)
	v4 := cache.Get(3)
	fmt.Println("v4", v4)
	v5 := cache.Get(4)
	fmt.Println("v5", v5)
}

func TestCache(t *testing.T) {
	cache := Constructor1(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	v1 := cache.Get(1)
	fmt.Println("v1", v1)
	v2 := cache.Get(2)
	fmt.Println("v2", v2)
}
