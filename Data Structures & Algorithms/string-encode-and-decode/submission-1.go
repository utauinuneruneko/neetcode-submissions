type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encodeNum func(num int) string
	encodeNum = func(num int) string {
		str := fmt.Sprintf("%v", num)
		for len(str) < 3 {
			str = "0"+ str
		}
		return str
	}
	res := encodeNum(len(strs))
	for _, str := range strs {
		res += encodeNum(len(str))
	}
	return res+strings.Join(strs, "")
}

func (s *Solution) Decode(encoded string) []string {
	var decodeNum func(s string) int
	decodeNum = func(s string) int{
		for len(s) > 1 && s[0] == '0' {
			s = s[1:]
		}
		num, _ := strconv.Atoi(s)
		return num
	}
	n := decodeNum(encoded[:3])
	encoded = encoded[3:]
	lengthArr := make([]int, n)
	for i := 0; i < n; i++ {
		lengthArr[i] = decodeNum(encoded[:3])
		encoded = encoded[3:]
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = encoded[:lengthArr[i]]
		encoded = encoded[lengthArr[i]:]
	}
	return res
}
