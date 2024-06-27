package main

import (
	"jsontree"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("jsontree", func(w *unison.Window) {
		w.Content().AddChild(jsontree.New().Layout())
	})
}
