package main

import (
	"fmt"
	"strings"
)

// 接口
type MessageBuilder interface {
	Build(message ...string) string
}

// 解除信息的装饰
type BaseMessageBuilder struct {
}

func (b *BaseMessageBuilder) Build(message ...string) string {
	return strings.Join(message, ",")
}

type QuoteMessageBuilder struct {
	Builder MessageBuilder
}

func (q *QuoteMessageBuilder) Build(message ...string) string {
	return "\"" + q.Builder.Build(message...) + "\""
}

type BraceMessageBuilder struct {
	builder MessageBuilder
}

func (b *BraceMessageBuilder) Build(message ...string) string {
	return "{" + b.builder.Build(message...) + "}"
}
func main() {
	var mb MessageBuilder
	mb = &BaseMessageBuilder{}
	str := mb.Build("hello word")
	fmt.Println(str)

	mb = &QuoteMessageBuilder{mb}
	str = mb.Build("hello world")
	fmt.Println(str)

	mb = &BraceMessageBuilder{mb}
	str = mb.Build("hello world")
	fmt.Println(str)
}
