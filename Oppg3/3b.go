package main

import (
	"fmt"

	"../fileutils"
)

func main() {
	slice1 := fileutils.FileToByteslice("./files/lang01.wl")
	fmt.Println(slice1)
	slice2 := fileutils.FileToByteslice("./files/lang02.wl")
	fmt.Println(slice2)
	slice3 := fileutils.FileToByteslice("./files/lang03.wl")
	fmt.Println(slice3)

}
