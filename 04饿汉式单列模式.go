package main

import "fmt"

type Singleton1 struct {
}

var singleton1 *Singleton1

// 初始化
func init() {
	singleton1 = &Singleton1{}
}

// 获取实例
func GetSingletom1() *Singleton1 {
	return singleton1
}
func main() {
	fmt.Println(singleton1)
}
