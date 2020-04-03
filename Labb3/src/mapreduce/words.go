package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"


// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {
		wg := new(sync.WaitGroup)
		wg.Add(10)
		a := strings.Fields(text)
		var devide int = len(a)/10
		c := 0
		k := make(chan map[string]int)
		for j := 0; j < 10; j++{
			b:= a[c:devide]
			go func (){
				freq := make(map[string]int)
				for i := 0; i < len(b); i++{
				t := strings.Trim(b[i], ".")
				f := strings.Trim(t, ",")
				freq[strings.ToLower(f)]++
				}
			k <- freq
			wg.Done()
		}()

		c= devide
		devide = devide+ devide
		if devide > len(a){
			devide = len(a)
	}
	}
	go func() {
		wg.Wait()
		close(k)
	}()
	fre := make(map[string]int)
	for slice := range k {
		for word, count := range slice {
			fre[word] += count
		}
	}

	return fre
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data

	data, _ := ioutil.ReadFile(DataFile)
	//fmt.Printf("%#v \n", WordCount(string(data)))
	numRuns := 1
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
