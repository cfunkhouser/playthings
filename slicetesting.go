package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genTestData(cap int) []int {
	return rand.Perm(cap)
}

func testCappedSlice(data []int, passes int) time.Duration {
	start := time.Now()
	for i := passes; i >= 0; i-- {
		s := make([]int, 0, len(data))
		for _, d := range data {
			s = append(s, d)
		}
	}
	delta := time.Since(start)
	return delta / time.Duration(passes)
}

func testUncappedSlice(data []int, passes int) time.Duration {
	start := time.Now()
	for i := passes; i >= 0; i-- {
		s := make([]int, 0)
		for _, d := range data {
			s = append(s, d)
		}
	}
	delta := time.Since(start)
	return delta / time.Duration(passes)
}

func testUncappedLengthSetIdxSlice(data []int, passes int) time.Duration {
	start := time.Now()
	for i := passes; i >= 0; i-- {
		s := make([]int, len(data))
		for x, d := range data {
			s[x] = d
		}
	}
	delta := time.Since(start)
	return delta / time.Duration(passes)
}

func testCappedLengthSetIdxSlice(data []int, passes int) time.Duration {
	start := time.Now()
	for i := passes; i >= 0; i-- {
		s := make([]int, len(data), len(data))
		for x, d := range data {
			s[x] = d
		}
	}
	delta := time.Since(start)
	return delta / time.Duration(passes)
}

func main() {
	passes := 1000
	for _, d := range []int{100, 1000, 10000, 100000, 1000000} {
		data := genTestData(d)
		fmt.Printf("Timing %d passes with %d data size\n", passes, len(data))
		fmt.Printf("\tcapped, 0 length, append:\t\t%v\n", testCappedSlice(data, passes))
		fmt.Printf("\tuncapped, 0 length, append:\t\t%v\n", testUncappedSlice(data, passes))
		fmt.Printf("\tcapped, set length, direct:\t\t%v\n", testCappedLengthSetIdxSlice(data, passes))
		fmt.Printf("\tuncapped, set length, direct:\t\t%v\n", testUncappedLengthSetIdxSlice(data, passes))
	}
}
