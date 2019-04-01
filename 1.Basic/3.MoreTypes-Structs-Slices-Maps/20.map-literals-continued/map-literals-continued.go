package main

import "fmt"

type vertex struct {
	Lat, Long float64
}

var m = map[string]vertex{
	"Bell Labs": {0, 1},
	"Google":    {2, 3},
}

func main() {
	fmt.Println(m)
}
