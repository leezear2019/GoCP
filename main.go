package main

import (
	"GoCP/CpUtil"
	"fmt"
	"strings"
)

func main() {
	//fmt.Println("xixi")
	//arr1 := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(arr1)
	//fmt.Println(len(arr1))
	//for i, v := range arr1 {
	//	println(i, v)
	//}

	//var a = CpUtil.Int2d{1, 2}
	//println(a)
	////println(change(a))
	//println(a)

	//var a =
	//
	//x, y := CpUtil.GetInt2(100)
	//fmt.Println(x, y)
	//
	//arr := []int{1, 2, 3}
	////{(1,1), (1,1), (1,1)}
	//fmt.Println(arr)
	//
	//a := CpUtil.Int2d{X: 1, Y: 1}
	//fmt.Println(a)
	//
	//aa2 := []CpUtil.Int2d{{1, 1}}
	//aa2 = append(aa2, CpUtil.Int2d{X: 1, Y: 3})
	//fmt.Println("aa2:", aa2)
	//
	//arr2 := [][2]int{{2, 2}}
	//fmt.Println(arr2)
	//aa := [2]int{1, 2}
	//arr2 = append(arr2, aa)
	//fmt.Println(arr2)
	//arr2 = append(arr2, [2]int{1, 3})
	//fmt.Println(arr2)
	//
	//aaa := CpUtil.Int2d{X: 1, Y: 2}
	//fmt.Println(aaa)
	//Change(aaa)
	//fmt.Println(aaa)

	sa := "1..4"
	sb := strings.Split(sa, "..")
	fmt.Println(sb)
	var lb, ub int
	//lb = int(sb[0])
	fmt.Sscanf("1..4", "%d..%d", &lb, &ub)
	fmt.Println(lb, ub)
}

func Change(a CpUtil.Int2d) {
	a.X += a.X
}
