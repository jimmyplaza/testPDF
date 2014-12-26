package main

import (
	"code.google.com/p/gofpdf"
	"fmt"
	"path/filepath"
)

func tutorial01() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello World!")
	fileStr := filepath.Join(cnGofpdfDir, "pdf/tutorial01.pdf")
	err := pdf.OutputFileAndClose(fileStr)
	if err == nil {
		fmt.Println("Successfully generated pdf/tutorial01.pdf")
	} else {
		fmt.Println(err)
	}
}
