package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var singleton *Singleton
var once sync.Once

func GetSingleton() *Singleton {
	if singleton != nil {
		once.Do(func() {
			singleton = &Singleton{}
		})
	}
	return singleton
}
func main() {
	singleton := GetSingleton()
	fmt.Println(singleton)
}
