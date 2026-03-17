package main

import (
	"fmt"
	"time"
)

func DisplaySummary(out <-chan Result) {
	successful := 0
	failed := 0
	var totalTime time.Duration
	total := 0
	result := []string{}
	for res := range out {
		if err := res.Err; err == nil {
			result = append(result, res.Data)
			successful++
		} else {
			failed++
		}
		totalTime += res.Duration
		total++
	}
	fmt.Println("Aggregation Summary:")
	fmt.Printf("\tSuccessful: %d/%d\n", successful, total)
	fmt.Printf("\tFailed: %d/%d\n", failed, total)
	fmt.Printf("\tTotal time: %v\n", totalTime)
	fmt.Printf("\tData:%v\n", result)
}
