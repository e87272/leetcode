package main

import "log"

func main() {
	str := longestCommonPrefix([]string{"flower", "flow", "flight"})
	log.Printf("str : %+v", str)
}
func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		for _, v := range strs {
			if i > len(v)-1 || v[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
