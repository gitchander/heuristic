package main

import (
	"log"

	"github.com/gitchander/heuristic/math/cubspl"

	"github.com/envoker/gotk3/cairo"
	"github.com/envoker/gotk3/gdk"
	"github.com/envoker/gotk3/gtk"
)

func main() {

	r := Rectangle{
		Min: cubspl.Point{50, 50},
		Max: cubspl.Point{450.5, 450.5},
	}

	se, err := NewSplineEditor(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer se.Close()

	gtk.Init(nil)

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal(err.Error())
	}

	drawingArea, err := gtk.DrawingAreaNew()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Event handlers
	{
		window.Connect("destroy", gtk.MainQuit)

		window.Connect("configure-event", func(win *gtk.Window, event *gdk.Event) {

			var (
				width  = drawingArea.GetAllocatedWidth()
				height = drawingArea.GetAllocatedHeight()
			)

			se.Resize(width, height)
		})

		window.Connect("key-press-event", func(win *gtk.Window, event *gdk.Event) {

			keyEvent := &gdk.EventKey{event}

			if se.KeyPress(keyEvent.KeyVal()) {
				win.QueueDraw()
			}
		})

		drawingArea.Connect("button-press-event", func(da *gtk.DrawingArea, event *gdk.Event) {

			eb := gdk.EventButton{event}

			x, y := eb.Pos()
			if se.ButtonPress(x, y) {
				da.QueueDraw()
			}
		})

		drawingArea.Connect("button-release-event", func(da *gtk.DrawingArea, event *gdk.Event) {

			se.ButtonRelease()
		})

		drawingArea.Connect("motion-notify-event", func(da *gtk.DrawingArea, event *gdk.Event) {

			em := gdk.EventMotion{event}

			if (gdk.EventMask(em.State()) & gdk.BUTTON_PRESS_MASK) != 0 {
				if se.ButtonMove(em.Pos()) {
					da.QueueDraw()
				}
			}
		})

		drawingArea.Connect("draw", func(da *gtk.DrawingArea, context *cairo.Context) {

			se.DrawCairoNative(context.Native())
		})
	}

	drawingArea.SetEvents(int(
		gdk.POINTER_MOTION_MASK |
			gdk.POINTER_MOTION_HINT_MASK |
			gdk.BUTTON_PRESS_MASK |
			gdk.BUTTON_RELEASE_MASK))

	window.Add(drawingArea)
	window.SetTitle(se.Name())
	window.SetSizeRequest(500, 500)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.ShowAll()

	gtk.Main()
}
