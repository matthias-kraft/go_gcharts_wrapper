package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func writeSubPlots(writer io.Writer, charts []Chart) error {
	t := template.Must(template.ParseFiles("gchart_placeholder_template.html", "gchart_template.html"))
	t.ExecuteTemplate(writer, "gchart_template", charts)
	t.ExecuteTemplate(writer, "layout","")
	return nil
}

func main() {
	// File needs to exist, use the one in the repo to test.
	file, err := os.Open("/tmp/testdata.csv")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ReadRecords(file)
	if err != nil {
		log.Fatal(err)
	}

	var dataArray [][]float64
	dataArray = append(dataArray, data["x"])
	dataArray = append(dataArray, data["y"])
	dataArray = Transpose(dataArray)


	options := template.JS(`{
	title: 'Parabola',
	curveType: 'function',
	legend: { position: 'bottom' },
	hAxis: {viewWindow: {max: 12, min: 0}},
	intervals: {'style':'area'},
	explorer: {},
	};`)


	c := Chart{"Chart1", "Line", Data{
		[]Header{{Label: "Func"}, {Label: "x^2"}},		dataArray}, options}
	c2 := c
	c2.Name = "Chart2"

	server := func(w http.ResponseWriter, r*http.Request) {
		writeSubPlots(w, []Chart{c,c2})
	}

	http.HandleFunc("/", server)
	http.ListenAndServe(":8080", nil)
}


