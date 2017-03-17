package main

import (
	"fmt"
	//"strconv"
	"time"
)

func main() {
	
	shoe := NewSortableShoe(6)
	for _, card := range shoe{
		fmt.Println(card)
	}

	time.Sleep(5000 * time.Millisecond)
}
