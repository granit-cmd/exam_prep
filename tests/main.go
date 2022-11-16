package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("file")
	if err != nil {
		log.Fatal(err)
	}

	ints, _ := ReadInts(strings.NewReader(string(content)))

	average1 := average(ints)
	median := median(ints)
	sd := math.Sqrt(float64(average1))

	fmt.Println("Average: ", average)
	fmt.Println("Median: ", median)
	fmt.Println("Standard Deviation: ", int(sd))
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func median(data []int) int {
	dataCopy := make([]int, len(data))
	copy(dataCopy, data)
	sort.Ints(dataCopy)

	var median int
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2])
	} else {
		median = dataCopy[l/2]
	}
	return median
}

func Variance(n []int) int {
	var a []int
	var b []int
	var c int
	for i := 0; i < len(n); i++ {
		a = append(a, float64(n[i])-(average(n)))
		b = append(b, float64(a[i])*float64(a[i]))
		c = c + b[i]
	}

	k := len(n)
	d := c / float64(k)
	return d
}

func average(ints []int) int {
	res := 0
	count := 0
	for i := 0; i <= len(ints)-1; i++ {

		res += ints[i]
		count++
	}
	return res / count
}
