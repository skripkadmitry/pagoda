package series

import (
	"fmt"
	"sort"
)

type Series interface {
	Implement(x interface{}) error

	Len() int
	Type() _Type
	typePrec() int

	fmt.Stringer
	sort.Interface

	Mask(expr string, functions ...MapFunction) boolSeries //return bool series
	//s.Mask(%self% > 5) 1 5 6 4 2 3 -> 0 0 1 0 0 0
	//s.Mask(%1% %self% > 2, MapSqrt)
	IsNa() boolSeries
	NotNa() boolSeries
	//Not() Series //inverse of bool
	IsFinite() boolSeries
	Filter(mask boolSeries) Series //filter values according to boolean mask

	More(series Series) (boolSeries, error)
	MoreEq(series Series) (boolSeries, error)
	Eq(series Series) (boolSeries, error)
	LessEq(series Series) (boolSeries, error)
	Lesser(series Series) (boolSeries, error) //to avoid clash with sort.Interface
	//And(series Series) (Series, error)
	//Or(series Series) (Series, error)
	//Xor(series Series) (Series, error)
	//Add(series Series) (Series, error)

	Convert(dsttype _Type) error //convert between types, not to factors
	//StringsToCat() error //rework series inplace

	Sort(rev bool) Series //sorts Series inplace and returns permutation
	SortbyPermute(perm Series)

	//ReorderCat([]string) error

	//Mean() float64
	//Std() float64
	//Max() Element
	//Min() Element
	//Med() Element
}

type MapFunction interface {
	func(element Element) Element
}
