package main

import (
	"log"
)

func main() {
	log.Printf("%+v", isMatch("aa", "a"))                    //false
	log.Printf("%+v", isMatch("a", ".*"))                    //true
	log.Printf("%+v", isMatch("a", ".*..a*"))                //false
	log.Printf("%+v", isMatch("aa", "a*aaa"))                //false
	log.Printf("%+v", isMatch("aaa", "a*a"))                 //true
	log.Printf("%+v", isMatch("aab", "c*a*b"))               //true
	log.Printf("%+v", isMatch("ab", ".*.."))                 //true
	log.Printf("%+v", isMatch("abbbcd", "ab*bbbcd"))         //true
	log.Printf("%+v", isMatch("mississippi", "mis*is*ip*.")) //true
	log.Printf("%+v", isMatch("abcdede", "ab.*de"))          //true
	log.Printf("%+v", isMatch("bbbba", ".*a*a"))             //true
	log.Printf("%+v", isMatch("ab", ".*....c*"))             //false
	log.Printf("%+v", isMatch("mississippi", "mis*is*p*."))  //false
	log.Printf("%+v", isMatch("b", "aaa."))                  //false

}

func isMatch(s string, p string) bool {
	var sIndex = 0
	for k := 0; k < len(p); k++ {
		str1 := p[k]
		if k+1 < len(p) && p[k+1] == '*' {
			i := 0
			for i = 0; i <= len(s)-sIndex; i++ {
				match := isMatch(s[sIndex+i:], p[k+2:])
				if match {
					return true
				} else if sIndex+i < len(s) && (s[sIndex+i] == str1 || str1 == '.') {
					continue
				} else {
					return false
				}
			}
			return false
		} else {
			if sIndex >= len(s) {
				return false
			}
			if s[sIndex] == str1 || str1 == '.' {
				sIndex = sIndex + 1
			} else {
				return false
			}
		}
	}
	if sIndex < len(s) {
		return false
	}
	return true
}
