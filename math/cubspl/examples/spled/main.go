package main

import (
	"fmt"

	gogdk "github.com/gitchander/go-lang/gtk/gdk"
	"github.com/gitchander/heuristic/math/cubspl"

	"github.com/conformal/gotk3/cairo"
	"github.com/conformal/gotk3/gdk"
	"github.com/conformal/gotk3/gtk"
)

func main() {

	r := Rectangle{
		Min: cubspl.Point{50, 50},
		Max: cubspl.Point{450.5, 450.5},
	}

	e, err := NewSplineEditor(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer e.Close()

	if err := runGUI(e); err != nil {
		fmt.Println(err.Error())
	}
}

func runGUI(e *SplineEditor) error {

	gtk.Init(nil)

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return err
	}

	drawingArea, err := gtk.DrawingAreaNew()
	if err != nil {
		return err
	}

	// Event handlers
	{
		window.Connect("destroy", gtk.MainQuit)

		window.Connect("configure-event", func(win *gtk.Window, event *gdk.Event) {

			var (
				width  = drawingArea.GetAllocatedWidth()
				height = drawingArea.GetAllocatedHeight()
			)

			e.Resize(width, height)
		})

		window.Connect("key-press-event", func(win *gtk.Window, event *gdk.Event) {

			keyEvent := &gdk.EventKey{event}

			if e.KeyPress(keyEvent.KeyVal()) {
				win.QueueDraw()
			}
		})

		drawingArea.Connect("button-press-event", func(da *gtk.DrawingArea, event *gdk.Event) {

			var eb gogdk.EventButton
			eb.FromNative(event.Native())

			x, y := eb.Pos()
			if e.ButtonPress(x, y) {
				da.QueueDraw()
			}
		})

		drawingArea.Connect("button-release-event", func(da *gtk.DrawingArea, event *gdk.Event) {

			e.ButtonRelease()
		})

		drawingArea.Connect("motion-notify-event", func(da *gtk.DrawingArea, event *gdk.Event) {

			var em gogdk.EventMotion
			em.FromNative(event.Native())

			if (gogdk.EventMask(em.State()) & gogdk.BUTTON_PRESS_MASK) != 0 {
				if e.ButtonMove(em.Pos()) {
					da.QueueDraw()
				}
			}
		})

		drawingArea.Connect("draw", func(da *gtk.DrawingArea, context *cairo.Context) {

			e.DrawCairoNative(context.Native())
		})
	}

	drawingArea.SetEvents(int(
		gogdk.POINTER_MOTION_MASK |
			gogdk.POINTER_MOTION_HINT_MASK |
			gogdk.BUTTON_PRESS_MASK |
			gogdk.BUTTON_RELEASE_MASK))

	window.Add(drawingArea)
	window.SetTitle(e.Name())
	window.SetSizeRequest(500, 500)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.ShowAll()

	gtk.Main()

	return nil
}
