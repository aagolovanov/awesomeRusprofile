package main

import (
	"fmt"
	"github.com/aagolovanov/awesomeRusprofile/pkg"
	"log"
)

func main() {
	comp, err := pkg.GetMainInfo("5258081758")
	if err != nil {
		log.Panicln(err)
	}

	kpp, err := pkg.GetCompanyKPP(comp)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf(
		"ИНН %s\n"+
			"КПП: %s\n"+
			"Название: %s\n"+
			"ФИО Рук.:%s\n", comp.INN, kpp, comp.Name, comp.FIO)
}
