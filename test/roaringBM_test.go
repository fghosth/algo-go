package algo_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/RoaringBitmap/roaring"
)

func init() {

}

func read3(path string) []byte {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return fd
}

func TestRoaringB(t *testing.T) {
	rb := roaring.NewBitmap()
	rb2 := roaring.NewBitmap()
	// count := uint32(4000000000)
	// for i := uint32(0); i < count; i++ {
	// 	rb.Add(i)
	// 	rb2.Add(i + 2)
	// }
	// buf := new(bytes.Buffer)
	// rb.WriteTo(buf)                                     // we omit error handling
	// err2 := ioutil.WriteFile("./rb", buf.Bytes(), 0666) //写入文件(字节数组)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	//
	// buf2 := new(bytes.Buffer)
	// rb2.WriteTo(buf2)                                     // we omit error handling
	// err3 := ioutil.WriteFile("./rb2", buf2.Bytes(), 0666) //写入文件(字节数组)
	// if err3 != nil {
	// 	fmt.Println(err3)
	// }

	buff := new(bytes.Buffer)
	buff.Write(read3("./rb"))
	rb.ReadFrom(buff)
	// newrb := roaring.NewBitmap()
	// newrb.ReadFrom(buff)
	// if rb.Equals(newrb) {
	// 	fmt.Println("I wrote the content to a byte stream and read it back.")
	// }

	buff2 := new(bytes.Buffer)
	buff2.Write(read3("./rb2"))
	rb2.ReadFrom(buff2)
	t1 := time.Now() // get current time
	roaring.ParOr(8, rb, rb2)
	// fmt.Println("rb1.And(rb2)", c.String())
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}

// func TestRoaringBitmap(t *testing.T) {
// 	// example inspired by https://github.com/fzandona/goroar
// 	fmt.Println("==roaring==")
// 	rb1 := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
// 	fmt.Println(rb1.String())
// 	fmt.Println(rb1.Select(2))
// 	rb2 := roaring.BitmapOf(3, 4, 1000)
// 	fmt.Println(rb2.String())
//
// 	rb3 := roaring.NewBitmap()
// 	fmt.Println(rb3.String())
//
// 	fmt.Println("Cardinality: ", rb1.GetCardinality())
//
// 	fmt.Println("Contains 3? ", rb1.Contains(3))
//
// 	rb1.And(rb2)
// 	fmt.Println("rb1.And(rb2)", rb1.String())
// 	rb3.Add(1)
// 	rb3.Add(5)
//
// 	rb3.Or(rb1)
// 	fmt.Println("rb3.Or(rb1)", rb3.String())
// 	// computes union of the three bitmaps in parallel using 4 workers
// 	c := roaring.ParOr(4, rb1, rb2, rb3)
// 	fmt.Println("roaring.ParOr(4, rb1, rb2, rb3)", c.String())
// 	// computes intersection of the three bitmaps in parallel using 4 workers
// 	d := roaring.ParAnd(8, rb1, rb2, rb3)
// 	fmt.Println("roaring.ParAnd(4, rb1, rb2, rb3)", d.String())
// 	// prints 1, 3, 4, 5, 1000
// 	i := rb3.Iterator()
// 	for i.HasNext() {
// 		fmt.Println(i.Next())
// 	}
// 	fmt.Println()
//
// 	// next we include an example of serialization
// 	buf := new(bytes.Buffer)
// 	rb1.WriteTo(buf) // we omit error handling
// 	newrb := roaring.NewBitmap()
// 	newrb.ReadFrom(buf)
// 	if rb1.Equals(newrb) {
// 		fmt.Println("I wrote the content to a byte stream and read it back.")
// 	}
// }
