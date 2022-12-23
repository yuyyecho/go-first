package main

import "fmt"

type NonRechargeableBattery interface {
	Use()
}
type RechargeableBattery interface {
	Use()
	Charge()
}
type NonRechargeableA struct {
}

func (n *NonRechargeableA) Use() {
	fmt.Println("使用一次性电池")
}

// 为对象添加适配器
type AdapterNonOrYes struct {
	NonRechargeableBattery
}

func (a *AdapterNonOrYes) Charge() {
	fmt.Println("一次性电池无法充电")
}

type RechargeableBatteryAbstract struct {
}

func (r *RechargeableBatteryAbstract) Use() {
	fmt.Println("正在使用蓄电电池")
}
func (r *RechargeableBatteryAbstract) Charge() {
	fmt.Println("正在为储电电池充电")
}

type RechargeableB struct {
	RechargeableBatteryAbstract
}

func (r *RechargeableB) Use() {
	fmt.Println("使用一次性电池B")
}
func (r *RechargeableB) Charge() {
	fmt.Println("正在为储电电池B充电")
}
func main() {
	var battery RechargeableBattery
	battery = &AdapterNonOrYes{&NonRechargeableA{}}
	battery.Use()
	battery.Charge()

	battery = &RechargeableB{}
	battery.Use()
	battery.Charge()
}
