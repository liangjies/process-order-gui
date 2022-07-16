package main

import (
	"process-order/pkgui"

	"github.com/ying32/govcl/vcl"
)

// go build -i -ldflags="-s -w -H windowsgui" -tags tempdll
func main() {
	vcl.Application.SetScaled(true)
	vcl.Application.SetTitle("project1")
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&pkgui.Form1)
	// vcl.Application.CreateForm(&pkgui.Form2)
	vcl.Application.Run()
}
