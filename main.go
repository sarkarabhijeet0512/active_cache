package main

import (
	activecache "audigent_task/activeCache"
	"fmt"
	"time"
)

func main() {
	cache := activecache.NewCache()
	// Calling Set and Get using the Cache interface
	cache.Set([]byte("key3"), []byte("value3"), 10*time.Second)
	value, duration := cache.Get([]byte("key3"))
	if value != nil {
		fmt.Println("Value for key3:", string(value), duration)
	} else {
		fmt.Println("Key3 not found")
	}
	time.Sleep(6 * time.Second)
	value, duration = cache.Get([]byte("key3"))
	if value != nil {
		fmt.Println("Value for key3:", string(value), duration)
	} else {
		// byte is a numeric type, actually an alias for uint8.
		// so means it has default zero value of 0 and cant be compare to nil or set to nil directly
		fmt.Println(nil)
	}
	time.Sleep(11 * time.Second)
	value, duration = cache.Get([]byte("key3"))
	if value != nil {
		fmt.Println("Value for key3:", string(value), duration)
	} else {
		// byte is a numeric type, actually an alias for uint8.
		// so means it has default zero value of 0 and cant be compare to nil or set to nil directly
		fmt.Println(nil)
	}
}
