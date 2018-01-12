package algo_test

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"testing"
)

func init() {

}
func TestRBTree(t *testing.T) {
	num := uint32(1000000000)
	var id [1000000000]uint32
	for i := uint32(0); i < num; i++ {
		id[i] = i
	}
	fmt.Println(id[3000000])
	fm, err := os.OpenFile("./tmp/mem.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(fm)
	fm.Close()
}
