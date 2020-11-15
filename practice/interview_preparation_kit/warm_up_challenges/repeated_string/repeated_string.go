// Package main implements Repeated String problem
//
// https://www.hackerrank.com/challenges/repeated-string/problem
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
  var stringLength = int64(len(s))
  var loops int64 = n / stringLength
  var prefixLength int64 = n % stringLength

  var occurancesInString int64 = 0
  var occurancesInPrefix int64 = 0

  for pos, char := range s {
    if char == 'a' {
      occurancesInString++
      if int64(pos) < prefixLength {
        occurancesInPrefix++
      }
    }
  }

  return occurancesInString * loops + occurancesInPrefix
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    var writer *bufio.Writer

    if (os.Getenv("OUTPUT_PATH") != "") {
      stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
      checkError(err)
      writer = bufio.NewWriterSize(stdout, 1024 * 1024)
      defer stdout.Close()
    } else {
      writer = bufio.NewWriterSize(os.Stdout, 1024 * 1024)
    }

    s := readLine(reader)

    n, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    result := repeatedString(s, n)

    fmt.Fprintf(writer, "%d\n", result)

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
