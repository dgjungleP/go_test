package main

func findTheDifference(s string, t string) byte {
	var diff byte
	for i := range s {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(s)]
}

func main() {

}
