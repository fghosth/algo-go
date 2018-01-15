package algo_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/RoaringBitmap/roaring"
)

func init() {

}

func TestRoaringBitmap(t *testing.T) {
	// example inspired by https://github.com/fzandona/goroar
	fmt.Println("==roaring==")
	rb1 := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
	fmt.Println(rb1.String())

	rb2 := roaring.BitmapOf(3, 4, 1000)
	fmt.Println(rb2.String())

	rb3 := roaring.NewBitmap()
	fmt.Println(rb3.String())

	fmt.Println("Cardinality: ", rb1.GetCardinality())

	fmt.Println("Contains 3? ", rb1.Contains(3))

	rb1.And(rb2)

	rb3.Add(1)
	rb3.Add(5)

	rb3.Or(rb1)

	// computes union of the three bitmaps in parallel using 4 workers
	ParOr(4, rb1, rb2, rb3)
	// computes intersection of the three bitmaps in parallel using 4 workers
	ParAnd(4, rb1, rb2, rb3)

	// prints 1, 3, 4, 5, 1000
	i := rb3.Iterator()
	for i.HasNext() {
		fmt.Println(i.Next())
	}
	fmt.Println()

	// next we include an example of serialization
	buf := new(bytes.Buffer)
	rb1.WriteTo(buf) // we omit error handling
	newrb := roaring.NewBitmap()
	newrb.ReadFrom(buf)
	if rb1.Equals(newrb) {
		fmt.Println("I wrote the content to a byte stream and read it back.")
	}
}
