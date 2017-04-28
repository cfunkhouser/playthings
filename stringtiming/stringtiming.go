package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func testDirectConcat(passes int, s string) time.Duration {
	start := time.Now()
	for i := passes; i >= 0; i-- {
		_ = s + "_something"
	}
	return time.Since(start) / time.Duration(passes)
}

func testBuffer(passes int, s string) time.Duration {
	start := time.Now()
	for i := passes; i >= 0; i-- {
		var b bytes.Buffer
		b.WriteString(s)
		b.WriteString("_something")
		_ = b.String()
	}
	return time.Since(start) / time.Duration(passes)
}

func testSprintf(passes int, s string) time.Duration {
	start := time.Now()
	for i := passes; i >= 0; i-- {
		_ = fmt.Sprintf("%s_something", s)
	}
	return time.Since(start) / time.Duration(passes)
}

func testArrayJoin(passes int, s string) time.Duration {
	start := time.Now()
	for i := passes; i >= 0; i-- {
		a := []string{s, "_something"}
		_ = strings.Join(a, "")
	}
	return time.Since(start) / time.Duration(passes)
}

func main() {
	testString := "test_string_or"
	for _, d := range []int{100, 1000, 10000, 100000, 1000000} {
		fmt.Printf("Timing %d passes with test string %q\n", d, testString)
		fmt.Printf("\t+= method took\t\t\t\t%v\n", testDirectConcat(d, "test_string"))
		fmt.Printf("\tbytes.Buffer method took\t\t%v\n", testBuffer(d, "test_string"))
		fmt.Printf("\tfmt.Sprintf method\t\t\t%v\n", testSprintf(d, "test_string"))
		fmt.Printf("\tarray join method\t\t\t%v\n", testArrayJoin(d, "test_string"))
	}
}
