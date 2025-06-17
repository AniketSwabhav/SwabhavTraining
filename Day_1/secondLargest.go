package main

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var firstLargest int = -1
	var secondLargest int = -1

	for i := 0; i < len(arr); i++ {
		if arr[i] > firstLargest {
			secondLargest = firstLargest
			firstLargest = arr[i]
		} else if arr[i] > secondLargest && arr[i] != firstLargest {
			secondLargest = arr[i]
		}
	}

	println("The second largest number is:", secondLargest)
}
