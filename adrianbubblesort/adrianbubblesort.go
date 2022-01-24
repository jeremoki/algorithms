package main

func bubblesort(arr []int) []int {



	if len(arr) == 1 {
		return arr
	}

	var i = 0

	var lastIndex = len(arr) - 1

	for i < lastIndex {

		if arr[lastIndex] < arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i+1]
			println(lastIndex)
		}
		i++
	}
	bubblesort(arr[0:lastIndex])
	return arr

}
func main() {
	var n = []int{8, 9, 1, 5, 7, 6}
	println(bubblesort(n))

}