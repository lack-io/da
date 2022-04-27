package idea

func countGoodTriplets(arr []int, a int, b int, c int) int {
	length := len(arr)

	count := 0
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < length; k++ {
				if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					count += 1
				}
			}
		}
	}

	return count
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
