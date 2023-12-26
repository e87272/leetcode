package main

import (
	"log"
	"regexp"
	"strings"
)

func main() {
	log.Printf("%+v", isMatch("cb", "?a"))
}
func isMatch(s string, p string) bool {

	matchbool, _ := regexp.MatchString("^"+strings.ReplaceAll(strings.ReplaceAll(p, "?", "[a-z]"), "*", "[a-z]*")+"$", s)
	return matchbool
}
