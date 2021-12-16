package helper

// 字符串首字母转成大写
func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArray := []rune(str)
	if strArray[0] >= 97 && strArray[0] <= 122 {
		strArray[0] -= 32
	}
	return string(strArray)
}

// 字符串首字母转成小写
func StrFirstToLower(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArr := []rune(str)
	if strArr[0] < 97 || strArr[0] > 122 {
		strArr[0] += 32
	}
	return string(strArr)
}
