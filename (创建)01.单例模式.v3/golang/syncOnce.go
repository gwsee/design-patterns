package main

import (
	"fmt"
	"sync"
)

var once2 sync.Once

type single2 struct {
}

var singleInstance2 *single2

func getInstance2() *single2 {
	if singleInstance2 == nil {
		once2.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance2 = &single2{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance2
}
