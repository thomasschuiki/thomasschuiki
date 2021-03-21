package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {
	out, err := os.Create("dist/index.html")
	if err != nil {
		return err
	}
	defer out.Close()

	data := map[string]interface{}{}

	buf, err := ioutil.ReadFile("vitae.yml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseGlob("tmpl/*")
	if err != nil {
		return err
	}
	err = tmpl.Execute(out, data)
	if err != nil {
		return err
	}
	return nil
}
