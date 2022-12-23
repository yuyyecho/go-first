package main

import "fmt"

type User struct {
}

func (u *User) Login(username string, password string) error {
	fmt.Println("登录成功")
	return nil
}

// 定义代理
type UserProxy struct {
	user *User
}

func (u *UserProxy) Login(username string, password string) error {
	err := u.user.Login(username, password)
	if err != nil {
		return err
	}
	//代理后 处理其他业务
	fmt.Println("统计点击次数")
	return nil
}

func main() {
	var userProxy UserProxy = UserProxy{user: new(User)}
	userProxy.Login("xueshen", "123456")
}
