package main

import "fmt"

type vertex struct {
	Lat, Long float64
}

var m = map[string]vertex{
	"Bell Labs": vertex{
		0, 1,
	},
	"Google": vertex{
		2, 3,
	},
}

func main() {
	fmt.Println(m)
}
