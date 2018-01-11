package algo_test

import (
	"bytes"
	"fmt"
	"testing"

	bloom "jvole.com/algo-go/BloomFilters"
)

func init() {

}

func TestMinHash(t *testing.T) {
	bag1 := []string{"bill", "alice", "frank", "bob", "sara", "tyler", "james", "wr"}
	bag2 := []string{"bill", "alice", "frank", "bob", "sara", "adsfasdfeae", "adsfaewr"}
	fmt.Println("similarity", bloom.MinHash(bag1, bag2))
}

func TestTopK(t *testing.T) {
	topk := bloom.NewTopK(0.001, 0.99, 3)

	topk.Add([]byte(`bob`)).Add([]byte(`bob`)).Add([]byte(`bob`))
	topk.Add([]byte(`tyler`)).Add([]byte(`tyler`)).Add([]byte(`tyler`)).Add([]byte(`tyler`))
	topk.Add([]byte(`fred`))
	topk.Add([]byte(`alice`)).Add([]byte(`alice`)).Add([]byte(`alice`)).Add([]byte(`alice`))
	topk.Add([]byte(`james`))
	topk.Add([]byte(`fred`))
	topk.Add([]byte(`sara`)).Add([]byte(`sara`))
	topk.Add([]byte(`bill`))

	for i, element := range topk.Elements() {
		fmt.Println(i, string(element.Data), element.Freq)
	}

	// Restore to initial state.
	topk.Reset()
}

func TestContMinSketch(t *testing.T) {
	cms := bloom.NewCountMinSketch(0.001, 0.99)

	for i := 0; i < 1000; i++ {
		cms.Add([]byte(`alice`)).Add([]byte(`bob`)).Add([]byte(`bob`)).Add([]byte(`frank`))
	}
	fmt.Println("frequency of alice", cms.Count([]byte(`alice`)))
	fmt.Println("frequency of bob", cms.Count([]byte(`bob`)))
	fmt.Println("frequency of frank", cms.Count([]byte(`frank`)))

	// Serialization example
	buf := new(bytes.Buffer)
	n, err := cms.WriteDataTo(buf)
	if err != nil {
		fmt.Println(err, n)
	}

	// Restore to initial state.
	cms.Reset()

	newCMS := bloom.NewCountMinSketch(0.001, 0.99)
	n, err = newCMS.ReadDataFrom(buf)
	if err != nil {
		fmt.Println(err, n)
	}

	fmt.Println("frequency of frank", newCMS.Count([]byte(`frank`)))
}

func TestCuckooFilter(t *testing.T) {
	cf := bloom.NewCuckooFilter(1000, 0.01)

	cf.Add([]byte(`a`))
	if cf.Test([]byte(`a`)) {
		fmt.Println("contains a")
	}

	if contains, _ := cf.TestAndAdd([]byte(`b`)); !contains {
		fmt.Println("doesn't contain b")
	}

	if cf.TestAndRemove([]byte(`b`)) {
		fmt.Println("removed b")
	}

	// Restore to initial state.
	cf.Reset()
}
