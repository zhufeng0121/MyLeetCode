package interview

func merge(A []int, m int, B []int, n int) {

	i, j := m-1, n-1

	for i >= 0 && j >= 0 {
		if A[i] > B[j] {
			A[i+j+1] = A[i]
			i--
		} else {
			A[i+j+1] = B[j]
			j--
		}
	}
	if i == -1 {
		for j >= 0 {
			A[j] = B[j]
			j--
		}

	}

}
