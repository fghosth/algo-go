package algo_test

import (
	"fmt"
	"testing"
	"time"
)

var hashmap map[int]int
var count int

func init() {
	count = 1000000
	hashmap = make(map[int]int, count)
}

func BenchmarkHashmap(b *testing.B) {
	for i := 0; i < count; i++ {
		hashmap[i] = i
	}
}

func BenchmarkHashmap2(b *testing.B) {
	hashmap2 := make(map[int]int)
	for i := 0; i < count; i++ {
		hashmap2[i] = i
	}
}

func TestHashmap(t *testing.T) {
	start := time.Now().UnixNano()
	for i := 0; i < count; i++ {
		hashmap[i] = i
	}
	end := time.Now().UnixNano()
	fmt.Println("所用时间(纳秒):", end-start)
}

func TestHashmap2(t *testing.T) {
	start := time.Now().UnixNano()
	hashmap2 := make(map[int]int)
	for i := 0; i < count; i++ {
		hashmap2[i] = i
	}
	end := time.Now().UnixNano()
	fmt.Println("所用时间(纳秒):", end-start)
}
