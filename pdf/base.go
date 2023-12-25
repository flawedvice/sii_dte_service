package pdf

import (
	"fmt"

	"github.com/go-pdf/fpdf"
)

func CreatePDF() *fpdf.Fpdf {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Println(err.Error())
	}
	return pdf
}
