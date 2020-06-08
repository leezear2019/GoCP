package main

import (
	"GoCP/CpUtil"
	"fmt"
)

type AA struct {
	va int
	vb []int
}

func main() {
	//fmt.Println("xixi")
	//arr1 := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(arr1)
	//fmt.Println(len(arr1))
	//for i, v := range arr1 {
	//	println(i, v)
	//}

	a := CpUtil.Int2d{X: 1, Y: 2}
	fmt.Println(a)
	fmt.Println(Change(a))
	fmt.Println(a)

	b := AA{1, []int{1, 2, 3, 4}}
	fmt.Println(b)
	Change2(b)
	fmt.Println(b)

	Change3(b)
	fmt.Println(b)

	Change4(&b)
	fmt.Println(b)

	//var a =
	//
	//x, y := CpUtil.GetInt2(100)
	//fmt.Println(x, y)

	//arr := []int{1, 2, 3}
	//{(1,1), (1,1), (1,1)}
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

	//sa := "1..4"
	//sb := strings.Split(sa, "..")
	//fmt.Println(sb)
	//var lb, ub int
	////lb = int(sb[0])
	//fmt.Sscanf("1..4", "%d..%d", &lb, &ub)
	//fmt.Println(lb, ub)
	//
	//sc := "1 2 3"
	//sd := strings.Split(sc, " ")
	//se := []int{}
	//for _, v := range sd {
	//	fmt.Println(v)
	//	sf, _ := strconv.Atoi(v)
	//	se = append(se, sf)
	//}
	////fmt.Printf("%T,%s", sd, sd)
	//
	//sg := [][]int{}
	//sh := make([][]int, 2, 2)
	//sg = append(sg, []int{1, 2, 3})
	//fmt.Println("sg: ", sg)
	//fmt.Println(sh)
}

func Change(a CpUtil.Int2d) CpUtil.Int2d {
	a.X += a.X
	return a
}

func Change2(a AA) {
	a.vb[0]++
}

func Change3(a AA) {
	a.va++
}

func Change4(a *AA) {
	a.va++
}
