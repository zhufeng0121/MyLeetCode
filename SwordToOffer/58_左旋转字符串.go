package SwordToOffer

func reverseLeftWords(s string, n int) string {
	arr := []byte(s)

	reverseStr(arr[:n])
	reverseStr(arr[n:])
	reverseStr(arr)

	return string(arr)
}

func reverseStr(arr []byte) []byte {
	for i := 0; i < len(arr)>>1; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}