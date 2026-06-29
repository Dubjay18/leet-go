package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Freq, windowFreq :=  [26]int{}, [26]int{}


	l := 0
	r := len(s1) -1


	for i,_ := range s1 {
		s1Freq[s1[i] - 'a']++
	}
		for i := l; i <= r; i++ {
			windowFreq[s2[i] - 'a']++
		} 
	 for r < len(s2) {
		matches := 0
		for v,_ := range windowFreq{
			if s1Freq[v] == windowFreq[v] {
				matches++
			}
		}
		if matches == 26 {
			return true
		} else {
					windowFreq[s2[l]-'a']-- 
			l++
			r++
			if r < len(s2) {
				windowFreq[s2[r]-'a']++ 
			}	
		}
	 }
	 return false
}