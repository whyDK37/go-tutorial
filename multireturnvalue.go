package main

import "fmt"

//(int, int) 在这个函数中标志着这个函数返回 2 个 int。
func vals() (int, int) {
	return 3, 7
}

func main() {
	//这里我们通过多赋值 操作来使用这两个不同的返回值。
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	//如果你仅仅想返回值的一部分的话，你可以使用空白定义符 _。
	_, c := vals()
	fmt.Println(c)
}
