package main

import (
	"fmt"
	"sort"
)

type Human struct {
	name string
	age  int
}

type AgeFactor []Human

func (a AgeFactor) Len() int           { return len(a) }
func (a AgeFactor) Less(i, j int) bool { return a[i].age > a[j].age }
func (a AgeFactor) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

//func (a AgeFactor) Len() int  {return len(a)}
func main() {
	audience := []Human{
		{"Alice", 35},
		{"Bob", 45},
		{"James", 25},
	}
	sort.Sort(AgeFactor(audience))
	fmt.Println(audience)
}
