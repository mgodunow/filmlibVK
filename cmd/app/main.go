package main

import (
	"filmLib/app"

	_ "github.com/lib/pq"
)

func main() {
	a := app.NewApp()
	a.ListenAndServe()
	}