package main

import (
	"fmt"
	"os"

	lib "./lib"
)

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		n, g := lib.ParseFile(filename)
		group := g.GetBestPath(n)
		g.SendAnts(group, n)
	} else if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguments.")
	}
}
