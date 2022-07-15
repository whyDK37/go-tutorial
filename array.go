package main

import "fmt"

func main() {
	//这里我们创建了一个数组 a 来存放刚好 5 个 int。元素的类型和长度都是数组类型的一部分。数组默认是零值的，对于 int 数组来说也就是 0。
	var a [5]int
	fmt.Println("emp:", a)
	//我们可以使用 array[index] = value 语法来设置数组指定位置的值，或者用 array[index] 得到值。
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	//使用内置函数 len 返回数组的长度
	fmt.Println("len:", len(a))
	//使用这个语法在一行内初始化一个数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	//数组的存储类型是单一的，但是你可以组合这些数据来构造多维的数据结构。
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	var i, j, k int
	//
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println("len(balance2) = ", len(balance2))
	/* 输出每个数组元素的值 */
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//
	fmt.Println("将索引为 1 和 3 的元素初始化")
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}
}
