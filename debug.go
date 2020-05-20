package main

import (
	"fmt"
	"strconv"
	"time"
)

func testRoutine() {
	for i := 0; i < 10; i++ {
		func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(500 * time.Millisecond)
}

func main() {
	testRoutine()
	scoreStr := "8.9"
	score, err := strconv.ParseFloat(scoreStr, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(score)

}
