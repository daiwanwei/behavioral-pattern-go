package main

import (
	"fmt"
	"pattern-golang/behavioral/visitor/museums"
)

type American struct {
	Citizenship string
}

func (a *American) VisitForWest(exhibit *museums.WestExhibit) {
	fmt.Printf("amercian visit %s exhibit\n", exhibit.GetName())
}

func (a *American) VisitForEast(exhibit *museums.EastExhibit) {
	fmt.Printf("amercian visit %s exhibit\n", exhibit.GetName())
}

type Asian struct {
	Citizenship string
}

func (a *Asian) VisitForWest(exhibit *museums.WestExhibit) {
	fmt.Printf("asian visit %s exhibit\n", exhibit.GetName())
}

func (a *Asian) VisitForEast(exhibit *museums.EastExhibit) {
	fmt.Printf("asian visit %s exhibit\n", exhibit.GetName())
}
