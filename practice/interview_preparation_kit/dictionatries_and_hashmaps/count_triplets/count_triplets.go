// SPDX-FileCopyrightText: 2023 Hackerrank Authors
// SPDX-FileCopyrightText: 2023 Vladimir Rusinov
// SPDX-License-Identifier: Apache-2.0

// https://www.hackerrank.com/challenges/count-triplets-1/problem
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
	"math/big"
)

func countTripletsNotWorking(arr []int64, r int64) int64 {  //nolint:deadcode,unused
	// Idea: put array values into map of value => number of times value appears.
	// Than for each value in map check if we have value*r and value*r*r in the map.
	//
	// However some tests fail and I'm not sure why.
	m := make(map[int64]int64)
	for _, v := range arr {
		// Go maps return "zero value" if there is no key
		m[v] += 1
	}
	var result int64
	if r == 1 {
		// Special case - need to do binomial coefficient for all values over 3
		for _, c := range m {
			z := big.Int{}
			result += z.Binomial(c, 3).Int64()
		}
	} else {
		for v, c := range m {
			// Again, maps "zero value" will return 0 if we don't have either of a triplets in map
			result += c * m[v*r] * m[v*r*r]
		}
	}
	return result
}

func countTriplets(arr []int64, r int64) int64 {
	// Another approach: have two hashmaps of "before" and "after" counters.
	// Loop over array and when we see a number check if there's n/r in before and n*r in after
	before := make(map[int64]int64)
	after := make(map[int64]int64)

	// Fill the after array:
	for _, v := range arr {
		after[v] += 1
	}

	var result int64

	for _, v := range arr {
		// Drop current element from after
		if after[v] != 0 {
			after[v] -= 1
		}
		// If it's divisible by r, check if we can make a triplet:
		if v % r == 0 {
			result += before[v/r] * 1 * after[v*r]
		}
		before[v] += 1
	}
	return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(nr[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    r, err := strconv.ParseInt(nr[1], 10, 64)
    checkError(err)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int64

    for i := 0; i < int(n); i++ {
        arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arr = append(arr, arrItem)
    }

    ans := countTriplets(arr, r)

    fmt.Fprintf(writer, "%d\n", ans)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
