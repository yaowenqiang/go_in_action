package main

func sliceDemo() {
	slice := make([]string, 5)

	//类型， 长度 ，容量
	slice2 := make([]int, 3, 5)
	// 长度不能大于容量
	// slice3 := make([]string 5,3)

	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}

	slice4 := []int{10, 20, 30}

	slice5 := []string{99: ""}

	//数组
	array := [3]int{10, 20, 30}
	//切片
	slice6 := []int{10, 20, 30}

	//nil切片
	var slice7 []int

	//创建空的整形切片
	slice8 := make([]int, 0)

	slice9 := []int{}

	slice10 := []int{10, 20, 30, 40}

	slice10[1] = 25

	slice11 := []int{10, 20, 30, 40, 50}

	slice12 := slice11[1:3]

	slice12[1] = 35
	//slice11[2] = 35

	slice13 := []int{10, 20, 30, 40, 50}

	slice14 := slice13[1:3]

	slice15 := append(slice14, 60)

	slice16 := []int{10, 20, 30, 40}

	//会创建新的数组，且容量翻倍
	//超过1000会变成1.23倍，每次增加25%
	slice17 := append(slice16, 50)

}
