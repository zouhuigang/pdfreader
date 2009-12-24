package main

import (
	"http"
	"io"
	"pdfread"
	"svg"
	"strm"
	"os"
	"fmt"
)

var pd *pdfread.PdfReaderT

// hello world, the web server
func HelloServer(c *http.Conn, req *http.Request) {
	c.SetHeader("Content-Type", "image/svg+xml; charset=utf-8")
	page := strm.Int(req.URL.RawQuery, 1) - 1
	io.WriteString(c, string(svg.Page(pd, page)))
}

func complain(err string) {
	fmt.Printf("%susage: pdserve foo.pdf\n", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		complain("")
	}
	pd = pdfread.Load(os.Args[1])
	if pd == nil {
		complain("Could not load pdf file!\n\n")
	}
	http.Handle("/hello", http.HandlerFunc(HelloServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: ", err.String())
	}
}