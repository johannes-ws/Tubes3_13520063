package main

/* mengembalikan border function */
func borderFunction(pattern string) []int {

	var fail [1000000]int
	fail[0] = 0
	var m int = len(pattern)
	var i int = 1 // iterator suffiks
	var j int = 0 // iterator prefiks

	for i < m {
		if pattern[j] == pattern[i] {
			fail[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			fail[i] = 0
			i++
		}
	}
	return fail[0:len(pattern)]
}

/*	menerima text dan pattern,
	mengembalikan indeks pattern pertama yang ditemukan di text,
	mengembalikan -1 jika ditak ada pattern dalam text
*/
func kmpMatch(text string, pattern string) int {
	var n int = len(text)
	var m int = len(pattern)

	var fail []int = borderFunction(pattern)

	var i int = 0 // iterator untuk text
	var j int = 0 // iterator untuk pattern

	for i < n {
		if pattern[j] == text[i] {
			if j == m-1 {
				return i - m + 1
			}
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			i++
		}
	}
	return -1
}
