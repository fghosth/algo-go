package lists

import (
	"github.com/emirpasic/gods/containers"
	"github.com/emirpasic/gods/utils"
)

// List interface that all lists implement
type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(values ...interface{})
	Contains(values ...interface{}) bool
	Sort(comparator utils.Comparator)
	Swap(index1, index2 int)
	Insert(index int, values ...interface{})

	containers.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
