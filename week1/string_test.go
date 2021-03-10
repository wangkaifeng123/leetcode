package week1

func restoreString(s string, indices []int) string {
	sByte := []byte(s)
	sHash := make([]byte, len(s))
	for k, v := range indices {
		sHash[v] = s[k]
	}
	for i := 0; i < len(s); i++ {
		sByte[i] = sHash[i]
	}
	return string(sByte)

}
