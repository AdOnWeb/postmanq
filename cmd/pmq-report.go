package main

import (
	"flag"
	"github.com/AdOnWeb/postmanq"
)

func main() {
	var file string
	flag.StringVar(&file, "f", postmanq.EXAMPLE_CONFIG_YAML, "configuration yaml file")
	flag.Parse()

	app := postmanq.NewReportApplication()
	if app.IsValidConfigFilename(file) {
		app.SetConfigFilename(file)
		app.Run()
	} else {
		flag.PrintDefaults()
	}
}
