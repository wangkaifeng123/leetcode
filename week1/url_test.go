package week1

func replaceSpaces(S string, length int) string {
	//从后往前排
	count := 0
	Bytes := []byte(S)
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			count++
		}
	}

	LengthS := length + 2*count - 1
	resLen := LengthS
	for i := length - 1; i >= 0; i-- {
		if S[i] != ' ' {
			Bytes[LengthS] = S[i]
			LengthS--
		} else {
			Bytes[LengthS] = '0'
			Bytes[LengthS-1] = '2'
			Bytes[LengthS-2] = '%'
			LengthS = LengthS - 3
		}
	}
	return string(Bytes[:resLen+1])
}
