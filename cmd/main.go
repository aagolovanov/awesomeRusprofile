package main

import (
	"fmt"
	"github.com/aagolovanov/awesomeRusprofile/pkg"
)

func main() {
	_, err := pkg.GetMainInfo("5258081758")
	if err != nil {
		fmt.Println(err)
	}
}
