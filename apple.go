package main

import (
	f "fmt"
)

func main() {
	var originalCount int = 10
	var eatenCount int = 4

	f.Println("I started with", originalCount)
	f.Println("Some jerk ate", eatenCount)
	f.Println("There are", originalCount - eatenCount)
}
