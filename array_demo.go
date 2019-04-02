package main

func main() {
	var array [5]int
	array2 := [5]int{1, 1, 2, 3, 4}
	array3 := [...]int{1, 2, 3, 43, 43, 4, 3, 3, 3}
	array4 := [5]int{1: 10, 2: 20}
	array4[2] = 5

	array5 := [5]*int{0: new(int), 1: new(int)}

	*array5[0] = 10
	*array5[1] = 20
	var array6 [5]string
	array7 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	array6 = array7

	var array8 [3]*string
	array9 := [3]*string{new(string), new(string), new(string)}
	*array9[0] = "Red"
	*array9[1] = "Blue"
	*array9[2] = "Green"

	array8 = array9

	var array10 [4][2]int

	array11 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

	array12 := [4][2]int{1: {0: 20}, 3: {1: 41}}

	var array13 [4][2]int

	array13[0][0] = 10
	array13[0][1] = 20
	array13[1][0] = 30
	array13[1][1] = 40

	var array14 [2][2]int
	var array15 [2][2]int

	array14[0][0] = 10
	array14[0][1] = 20
	array14[1][0] = 30
	array14[1][1] = 40

	array15 = array14

	var array16 [2]int = array15[0]

	var value17 int = array15[0][0]

	var array18 [1e6]int
	foo(array18)

	var array19 [1e6]int
	bar(&array18)

}
func bar(array *[1e6]int) {
}

func foo(array [1e6]int) {
}
