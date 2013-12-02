// Copyright (c) 2009 Helmar Wodtke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// The MIT License is an OSI approved license and can
// be found at
//   http://www.opensource.org/licenses/mit-license.php

// HTTP-server example.
package main

import (
	"fmt"
	"github.com/nathankerr/pdfreader"
	"github.com/nathankerr/pdfreader/strm"
	"github.com/nathankerr/pdfreader/svg"
	"io"
	"net/http"
	"os"
)

var pd *pdfreader.PdfReaderT

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "image/svg+xml; charset=utf-8")
	page := strm.Int(req.URL.RawQuery, 1) - 1
	io.WriteString(w, string(svg.Page(pd, page)))
}

func complain(err string) {
	fmt.Printf("%susage: pdserve foo.pdf\n", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		complain("")
	}
	pd = pdfreader.Load(os.Args[1])
	if pd == nil {
		complain("Could not load pdf file!\n\n")
	}
	http.Handle("/hello", http.HandlerFunc(HelloServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
