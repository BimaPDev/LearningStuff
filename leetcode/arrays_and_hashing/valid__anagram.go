package main

func isAnagram(s string, t string) bool {
	rs, rt := []rune(s), []rune(t) // convert string to rune so it work with all of ASCII
	if len(rs) != len(rt) {        // Check if the rune equal to each other
		return false
	}
	countS, countT := make(map[rune]int), make(map[rune]int) //create a hashmap of the runes
	for _, r := range rs {
		countS[r]++
	} // count the S
	for _, r := range rt {
		countT[r]++
	}
	// count the T

	// this will count the value inside the hashmap of T and R
	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	// check if the count equal to each other. R = T
	return true
	// return true
}
