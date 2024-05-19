package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type RandomNumberResponse struct {
	Numbers []int `json:"numbers"`
}

func generateRandomNumbers(max int) (int, error) {
	min := 0
	count := 1
	url := fmt.Sprintf("http://www.randomnumberapi.com/api/v1.0/random?min=%d&max=%d&count=%d", min, max, count)

	resp, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	var randomNumberResp RandomNumberResponse
	if err := json.Unmarshal(body, &randomNumberResp); err != nil {
		return -1, err
	}

	return randomNumberResp.Numbers[0], nil
}

//func generateRandomNumbers(maxRange int) (int, error) {
//	return rand.Intn(maxRange), nil
//}

func fillArrayWithRandomIntegers(maxRange int, arr []int, startIdx, endIdx int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := startIdx; i < endIdx; i++ {
		number, _ := generateRandomNumbers(maxRange) // Generate random integers between 0 and 99
		arr[i] = number
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		return
	}
	maxRange, _ := strconv.Atoi(os.Args[1])
	arraySize, _ := strconv.Atoi(os.Args[2])
	//arraySize := 100
	numGoroutines := 8 // Number of goroutines allowed

	arr := make([]int, arraySize)

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	chunkSize := arraySize / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		startIdx := i * chunkSize
		endIdx := startIdx + chunkSize
		if i == numGoroutines-1 {
			// Adjust end index for the last chunk to ensure it covers the entire array
			endIdx = arraySize
		}
		go fillArrayWithRandomIntegers(maxRange, arr, startIdx, endIdx, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish

	end := time.Now()

	// Print the filled array
	//fmt.Println("first array ", arr, end.Sub(start))
	fmt.Println("time to fill in first array ", end.Sub(start))

	arr2 := make([]int, arraySize)

	start = time.Now()
	for i := 0; i < arraySize; i++ {
		number, _ := generateRandomNumbers(maxRange)
		arr2[i] = number
	}
	end = time.Now()

	//fmt.Println("second array ", arr2)
	fmt.Println("time to fill second array ", end.Sub(start))

}
