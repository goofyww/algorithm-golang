package arr

func FindSubArrayList(arr []int, k int) (parent [][]int) {
	n := len(arr)
	pi := 0
	parent = make([][]int, n, n)
	for i := 0; i < n; i++ {
		ii := 0
		sub := make([]int, n)
		sum := 0
		for j := i; j < n; j++ {
			sub[ii] = arr[j]
			ii++
			sum += arr[j]
			if sum > k {
				break
			}
			if sum == k && arr[j+1] != 0 {
				parent[pi] = sub
				pi++
				break
			}
		}
	}
	return parent
}
