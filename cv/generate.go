package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	html2pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
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
	err := generateHTML()
	if err != nil {
		return err
	}
	err = generatePDF()
	if err != nil {
		return err
	}
	return nil
}

func generateHTML() error {
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

func generatePDF() error {
	// Create new PDF generator
	pdfg, err := html2pdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(600)
	pdfg.Orientation.Set(html2pdf.OrientationPortrait)
	pdfg.PageSize.Set(html2pdf.PageSizeA4)
	pdfg.MarginTop.Set(10)
	pdfg.MarginRight.Set(10)
	pdfg.MarginBottom.Set(10)
	pdfg.MarginLeft.Set(10)
	// Create a new input page from an URL
	page := html2pdf.NewPage("dist/index.html")
	page.Zoom.Set(1)

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)

	page.EnableLocalFileAccess.Set(true)
	page.DisableSmartShrinking.Set(true)
	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./dist/DI Thomas Schuiki - CV.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done
	return nil
}
