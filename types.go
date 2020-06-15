package main

import "html/template"

type Header struct {
	Label, Role string
}

type Data struct {
	Header []Header
	Rows [][]float64
}

type Chart struct {
	Name template.JS
	Type template.JS
	Data Data
	Options template.JS
}