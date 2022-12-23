package main

import (
	"errors"
	"fmt"
	"reflect"
)

// 动态代理调用接口
type Handle interface {
	Invoke(proxy *Proxy, method *Method, args []interface{}) ([]interface{}, error)
}
type Proxy struct {
	target  interface{}        //目标类
	Methods map[string]*Method //转载不同的方法
	handle  Handle             //同一处理的接口
}
type Method struct {
	value reflect.Value
}

func NewProxy(target interface{}, h Handle) *Proxy {
	ty := reflect.TypeOf(target)           //用来显示目标类动态的真实类型
	value := reflect.ValueOf(target)       //获取目标类的值
	Methods := make(map[string]*Method, 0) //初始化目标类的方法map
	for i := 0; i < value.NumMethod(); i++ {
		method := value.Method(i)
		Methods[ty.Method(i).Name] = &Method{value: method}
	}
	return &Proxy{
		target:  target,
		Methods: Methods,
		handle:  h,
	}
}

// 代理调用代理方法
func (p *Proxy) InvokeMethod(name string, args ...interface{}) ([]interface{}, error) {
	return p.handle.Invoke(p, p.Methods[name], args)
}

// 这里相当于调用原方法，在该方法外可以做方法增强，需要调用者自己实现
func (m *Method) Invoke(args ...interface{}) (res []interface{}, err error) {
	//用来捕捉异常
	defer func() {
		if p := recover(); p != nil {
			err = errors.New(fmt.Sprintf("%s", p))
		}
	}()
	//处理函数
	params := make([]reflect.Value, 0)
	if args != nil {
		//调用方法
		for i := 0; i < len(args); i++ {
			params = append(params, reflect.ValueOf(args[i]))
		}
	}
	//调用方法
	call := m.value.Call(params)
	res = make([]interface{}, 0)
	if call != nil && len(call) > 0 {
		for i := 0; i < len(call); i++ {
			res = append(res, call[i].Interface())
		}
	}
	fmt.Println("处理新的功能")
	return
}

type Test struct {
}

func (t *Test) Invoke(proxy *Proxy, method *Method, args []interface{}) ([]interface{}, error) {
	return args, nil
}

// 用来加载方法的结构体
func main() {
	t := Test{}
	t.Invoke(new(Proxy), new(Method), []interface{}{"a", "b", "c"})
	var h Handle
	h = &t
	proxy := NewProxy("xueshen", h)
	values, err := proxy.InvokeMethod("xueshen", "a", "b", "c")
	if err != nil {
		return
	}
	fmt.Println(values)
}
