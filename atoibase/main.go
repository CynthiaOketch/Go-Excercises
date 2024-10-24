package main

import ()

func AtoiBase(s string,base string) int{
	if len(base) < 1{
		return 0
	}
	basemap := make(map[rune]int)
	for i,c := range base {
		if c == '-' || c == basemap[c] > 0 {
			return 0
		}
		basemap[c] = i + 1
	}
	baselen := len(base)
	res := 0 
	for _,c := range s {
		val,seen :=
	}

}