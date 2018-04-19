package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func Wordcount(s string) map[string]int {
	arr_s := strings.Fields(s)
	map_s := make(map[string]int)
	for i := 0; i < len(arr_s); i++ {
		if map_s[arr_s[i]] == 0 {
			map_s[arr_s[i]] = 1
		} else {
			map_s[arr_s[i]] = map_s[arr_s[i]] + 1
		}
	}
	return map_s
}

func main() {
	s := "your mother BOOM BOOM BOOM BOOM fuck your asshole "
	res := Wordcount(s)
	fmt.Println(res)
	wc.Test(Wordcount)
}
