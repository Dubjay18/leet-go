package main



func backspaceCompare(s string, t string) bool {
	if processString(s) ==  processString(t) {
		return true
	}
	return  false
}


func processString(str string) string {
	newStr := ""
	for _,v := range str {
			if v == '#' {
				if len(newStr) > 0 {

					newStr = newStr[:len(newStr)-1]
				}
				continue
			}
				newStr += string(v)
	}
	return newStr
}