package main

import "fmt"

func main() {

	// 方法一: 类似数组的声明, 但是不指定大小
	//创建一个切片
	s1 := []int{}  //等同于: s := make([]int,0)
	s1 = append(s1, 1, 2, 3) //一定要重新赋值给变量s, 不然会报错: evaluated but not used
	fmt.Println(s1) // [1 2 3]

	//方法二: make创建
	//创建一个大小为2, 容量为10的切片.
	s2 := make([]int, 2, 10)  //前面两个为以0填充, 如果容量不够也会自动扩展
	s2[0]=99 //设置第一个值为: 99
	s2 = append(s2, 1, 2, 3)
	fmt.Println(s2) // [0 0 1 2 3]

	//创建并初始化
	s3 := []int{1, 2, 3, 4, 5}
	fmt.Println(s3)
}
