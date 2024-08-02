package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	r := gin.Default()
	r.GET("/", helloWorldHandler)
	r.GET("/pdf", pdfHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func helloWorldHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "i am alive",
	})
}

func pdfHandler(c *gin.Context) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, World!")
	pdf.Output(c.Writer)
}
