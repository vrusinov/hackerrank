// SPDX-FileCopyrightText: 2022 Hackerrank Authors
// SPDX-FileCopyrightText: 2022 Vladimir Rusinov
//
// SPDX-License-Identifier: Apache-2.0
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type void struct{}

func twoStrings(s1 string, s2 string) string {
    // No point in checking substrings over length 1.
	// So create a set of all letters in s1, and check if any of them are in s2
	var member void
	s1_letters := make(map[rune]void)
	for _, char := range s1 {
		s1_letters[char] = member
	}
	for _, char := range s2 {
		if _, ok := s1_letters[char]; ok {
			return "YES"
		}
	}
	return "NO"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	var writer *bufio.Writer

	if (os.Getenv("OUTPUT_PATH") != "") {
    	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    	checkError(err)

    	defer stdout.Close()

    	writer = bufio.NewWriterSize(stdout, 1024 * 1024)
	} else {
		writer = bufio.NewWriterSize(os.Stdout, 1024 * 1024)
	}

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        s1 := readLine(reader)

        s2 := readLine(reader)

        result := twoStrings(s1, s2)

        fmt.Fprintf(writer, "%s\n", result)
    }

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
