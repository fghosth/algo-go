package algo_test

import (
	"fmt"
	"testing"
	"time"

	"jvole.com/algo-go/gcache"
)

func init() {

}

func TestGcache(t *testing.T) {
	gc := gcache.New(20).
		LRU(). //LFU()  ARC()
		Build()
	gc.Set("key", "ok")
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}
func TestLRU(t *testing.T) {
	gc := gcache.New(20).
		LRU().
		Build()
	gc.Set("key", "ok")
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}

func TestGcache2(t *testing.T) {
	var evictCounter, loaderCounter, purgeCounter int
	gc := gcache.New(20).
		LRU().
		LoaderExpireFunc(func(key interface{}) (interface{}, *time.Duration, error) {
			loaderCounter++
			expire := 1 * time.Second
			return "ok", &expire, nil
		}).
		EvictedFunc(func(key, value interface{}) {
			evictCounter++
			fmt.Println("evicted key:", key)
		}).
		PurgeVisitorFunc(func(key, value interface{}) {
			purgeCounter++
			fmt.Println("purged key:", key)
		}).
		Build()
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
	time.Sleep(1 * time.Second)
	value, err = gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
	gc.Purge()
	if loaderCounter != evictCounter+purgeCounter {
		panic("bad")
	}

}
