package main

import (
	"log"
	"sort"
	"strings"
)

func main() {
	log.Printf("%+v", groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {

	groupMap := map[string][]string{}

	for _, str := range strs {
		strAry := strings.Split(str, "")
		sort.Strings(strAry)
		strKey := ""
		for _, v := range strAry {
			strKey = strKey + v
		}
		groupMap[strKey] = append(groupMap[strKey], str)
	}

	group := [][]string{}
	for _, v := range groupMap {
		group = append(group, v)
	}
	return group
}

func groupAnagrams2(strs []string) [][]string {

	groupMap := map[[26]byte][]string{}

	for _, str := range strs {
		strKey := [26]byte{}
		for _, v := range str {
			strKey[v-'a']++
		}
		groupMap[strKey] = append(groupMap[strKey], str)
	}

	group := [][]string{}
	for _, v := range groupMap {
		group = append(group, v)
	}
	return group
}
