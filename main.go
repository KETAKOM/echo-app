package main

import (
	"github.com/KETAKOM/echo-app/route"
)

func main() {
	// Echo instance
	e := route.Init()
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
