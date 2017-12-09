package main

import "fmt"

func main() {
	coll := []string{
		"hello",
		"test",
		"hrllowrold",
	}

	checkColl := func(v string) string {
		if len(v) < 5 {
			return v + "***"
		}
		return v
	}
	fmt.Println(Map(checkColl, coll))
}

func Map(f func(v string) string, vs []string) []string {
	if len(vs) == 0 {
		return []string{"hello"}
	}
	return append(
		[]string{f(vs[0])},
		Map(f, vs[1:])...)
}
