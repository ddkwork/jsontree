package main

import (
	"jsontree"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("jsontree", func(w *unison.Window) {
		w.Content().AddChild(jsontree.New().Layout())
	})
}
