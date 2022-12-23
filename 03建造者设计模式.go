package main

import "fmt"

type Itemer interface {
	Price() float64 //价格
	Name() string   //名字
	//	Category() string //分类
}

type Food struct {
}

type Drink struct {
}

type Hanburger struct {
}

func (Hanburger) Price() float64 {
	return 12.00
}
func (Hanburger) Name() string {
	return "Hanburger"
}

type FriedChicken struct {
	Food
}

func (FriedChicken) Price() float64 {
	return 18.00
}
func (FriedChicken) Name() string {
	return "FriedChicken"
}

type Cola struct {
	Drink
}

func (Cola) Price() float64 {
	return 3.00
}
func (Cola) Name() string {
	return "Cola"
}

type Beer struct {
	Drink
}

func (Beer) Price() float64 {
	return 5.00
}
func (Beer) Name() string {
	return "beer"
}

// 条目接口集合的别名
type Item []Itemer

// 建造者
type Builder struct {
}

func (i *Item) AddItem(Item ...Itemer) {
	*i = append(*i, Item...)
}

func (i *Item) GetCost() (cost float64) {
	for _, v := range *i {
		cost += v.Price()
	}
	return
}
func (i *Item) ShowItems() (msg string) {
	for _, v := range *i {
		msg += "名称：" + v.Name() + "\n"
	}
	return msg
}
func (b Builder) CreateBuilder() *Item {
	item := new(Item)
	item.AddItem(new(FriedChicken), new(Beer))
	return item
}
func main() {
	b := Builder{}
	item := b.CreateBuilder()
	msg := item.ShowItems()
	price := item.GetCost()
	fmt.Println(msg, price)
}
