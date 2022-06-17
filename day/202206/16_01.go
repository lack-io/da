package _02206

func duplicateZeros(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		if arr[i] != 0 {
			continue
		}
		if i < length-1 {
			for j := length - 1; j > i+1; j-- {
				arr[j] = arr[j-1]
			}
			arr[i+1] = 0
			i += 1
		}
	}
}

func duplicateZeros1(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		if arr[i] != 0 {
			continue
		}
		if i < length-1 {
			for j := length - 1; j > i+1; j-- {
				arr[j] = arr[j-1]
			}
			arr[i+1] = 0
			i += 1
		}
	}
}
