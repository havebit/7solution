package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(5)
	start := time.Now()
	for i := range 5 {
		go slowPrint(i)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}

func slowPrint(n int) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println(n)
	wg.Done()
}

func peter() func() int {
	n := 1
	return func() int {
		n++
		return n
	}
}

func couple(str string) []string {
	s := []rune(str)
	r := []string{}
	for s = append(s, '-'); len(s) > 1; s = s[2:] {
		r = append(r, string(s[:2]))
	}
	return r
}

func reverse(list [4]int) [4]int {
	var r [4]int

	for i, j := 0, 3; i < 4; i, j = i+1, j-1 {
		r[i] = list[j]
	}
	return r
}

func printPrime(n int) {
	for i := n; i > 1; i-- {
		count := 0
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 1 {
			fmt.Print(i, " ")
		}
	}
}

func monthlyInstallment(f, rate, year float64) float64 {
	return ((f * rate * year) + f) / (year * 12)
}

func add(a, b int) int {
	return a + b
}
