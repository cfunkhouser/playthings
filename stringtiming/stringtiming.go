package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const passes = 1000000 // 1 million

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func randStringArr(lengthRange, size int) []string {
	r := make([]string, size)
	for i := 0; i < size; i++ {
		r[i] = randString(rand.Intn(lengthRange))
	}
	return r
}

func testDirectConcat(passes int, sa []string) time.Duration {
	start := time.Now()
	for i := passes; i >= 0; i-- {
		var r string
		for _, s := range sa {
			r += s
		}
	}
	return time.Since(start) / time.Duration(passes)
}

func testBuffer(passes int, sa []string) time.Duration {
	start := time.Now()
	for i := passes; i >= 0; i-- {
		var b bytes.Buffer
		for _, s := range sa {
			b.WriteString(s)
		}
		_ = b.String()
	}
	return time.Since(start) / time.Duration(passes)
}

func testArrayJoin(passes int, sa []string) time.Duration {
	start := time.Now()
	for i := passes; i >= 0; i-- {
		_ = strings.Join(sa, "")
	}
	return time.Since(start) / time.Duration(passes)
}

func main() {
	// Anything larger than 1000 takes a long time to run, for little benefit.
	for _, d := range []int{2, 10, 100, 1000} {
		data := randStringArr(25, d)
		fmt.Printf("Timing %d passes with %d strings\n", passes, d)
		fmt.Printf("  += method averaged                %v\n", testDirectConcat(d, data))
		fmt.Printf("  bytes.Buffer method averaged      %v\n", testBuffer(d, data))
		fmt.Printf("  strings.Join method averaged      %v\n", testArrayJoin(d, data))
	}
}
