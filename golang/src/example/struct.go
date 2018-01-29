package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
}

func (*User) String() string {
	return "ToString"
}

func (u User) Name() string {
	if u.name == "" {
		return "---"
	}
	return u.name
}

func main() {

	var user1 User
	var user2 = User{"Zero"}
	var user3 = &User{"May"}

	fmt.Println("Type:", reflect.TypeOf(user1))
	fmt.Println("Type:", reflect.TypeOf(user2))
	fmt.Println("Type:", reflect.TypeOf(user3))

	fmt.Println(reflect.TypeOf(&user1))
	fmt.Println(reflect.TypeOf(&user2))
	fmt.Println(reflect.TypeOf(&user3))

	fmt.Println(user1.Name())
	fmt.Println(user2.Name())
	fmt.Println(user3.Name())

	fmt.Println(user1.String())
	fmt.Println(user2.String())
	fmt.Println(user3.String())

	fmt.Println(User{"Zero"}.Name())

	//User{"Zero"} 不是一个指针类型, 所以不能调用指针方法 String(), 但是可以调用Name()方法
	//fmt.Println(User{"Zero"}.String()) //error: cannot call pointer method on User literal


	//user1 不是指针类型, 会忽略掉所有指针方法(不能寻址的对象不能调用指针方法), 所以相当于没有实现fmt.Stringer接口
	var _ fmt.Stringer = user1 //error: cannot use user1 (type User) as type fmt.Stringer in assignment: User does not implement fmt.Stringer (String method has pointer receiver)
	//
	var _ fmt.Stringer = &user1
	var _ fmt.Stringer = user3


	/*
	type Stringer interface {
		String() string
	}
	*/

}

