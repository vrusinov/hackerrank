package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Finds the minimum number of swaps required to sort an array.
// n^2 solution would have been to do a selection sort and count the number of
// actual swaps.
// Howewer ints are consecutive and unique, so it feels like we can use that.
func minimumSwaps(arr []int32) int32 {
	// First find the smallest element - O(N)
	min := arr[0]
	for _, v := range arr {
		if (v < min) {
			min = v
		}
	}
	// Now, let's "normalize" the array.
	for i, _ := range arr {
		arr[i] = arr[i] - min
	}
	// Build a map of value to index
	vToI := make(map[int32]int32, len(arr))
	for i, v := range arr {
		vToI[v] = int32(i)
	}
	// Now we have an array with values from 0 to N.
	var swaps int32 = 0
	for i, v := range arr {
		// fmt.Printf("i=%v, v=%v, arr=%v\n", i, v, arr)
		if (int32(i) == v) {
			// Element is in its place, no swap needed.
			continue
		}
		// else swap is needed.
		arr[vToI[int32(i)]] = v
		arr[i] = int32(i)
		vToI[v] = vToI[int32(i)]
		vToI[int32(i)] = int32(i)
		swaps += 1
		// fmt.Printf("after swap: arr=%v\n", arr)
	}
	return swaps
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	var writer *bufio.Writer

	if (os.Getenv("OUTPUT_PATH") != "") {
    	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    	checkError(err)

    	defer stdout.Close()

    	writer = bufio.NewWriterSize(stdout, 1024 * 1024)
	} else {
		writer = bufio.NewWriterSize(os.Stdout, 1024 * 1024)
	}

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    res := minimumSwaps(arr)

    fmt.Fprintf(writer, "%d\n", res)

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
