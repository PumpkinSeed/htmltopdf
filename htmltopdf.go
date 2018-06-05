package htmltopdf

import (
	"bytes"
	"io/ioutil"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

const path = "/usr/local/bin/wkhtmltopdf"

func Example(file string, pdfFile string) error {
	html, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	wkhtmltopdf.SetPath(path)

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(html))
	page.NoBackground.Set(true)
	page.DisableExternalLinks.Set(false)
	pdfg.AddPage(page)
	pdfg.Dpi.Set(350)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginTop.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile(pdfFile)
	if err != nil {
		return err
	}

	return nil
}
