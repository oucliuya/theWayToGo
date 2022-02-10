package main

import (
	"fmt"
	"theWayToGo/pkg/charpt11/exercises/11.7/float64"
)

func main() {
	fa := float64.NewFloat64Array()
	fa.Fill(5)
	fmt.Printf("before sorted: %s\n", fa)
	fmt.Printf("isSorted: %v\n", float64.IsSorted(fa))
	float64.Sort(fa)
	fmt.Printf("after sorted: %s\n", fa)
	fmt.Printf("isSorted: %v\n", float64.IsSorted(fa))
}
