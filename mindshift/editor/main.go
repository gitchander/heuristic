package main

import (
	"os"
	"path/filepath"

	"github.com/mattn/go-gtk/gdk"
	//"github.com/mattn/go-gtk/gdkpixbuf"
	//"github.com/mattn/go-gtk/glib"

	"github.com/mattn/go-gtk/gtk"
)

func main() {

	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("MindShift Editor")
	window.Connect("destroy", gtk.MainQuit)
	window.SetPosition(gtk.WIN_POS_CENTER)

	w := makeElements()

	window.Add(w)
	window.SetSizeRequest(640, 480)
	window.ShowAll()

	gtk.Main()
}

func makeElements() gtk.IWidget {

	vbox := gtk.NewVBox(false, 0)

	// toolbar
	vbox.PackStart(newToolbar(), false, false, 3)

	// custom
	{
		vpaned := gtk.NewHPaned()

		vpaned.Pack1(newTextView(), false, false)
		vpaned.Pack2(newDrawingArea(), false, false)

		vbox.Add(vpaned)
	}

	return vbox
}

func newToolbar() gtk.IWidget {

	toolbar := gtk.NewToolbar()
	toolbar.SetStyle(gtk.TOOLBAR_ICONS)

	btnnew := gtk.NewToolButtonFromStock(gtk.STOCK_MEDIA_PLAY)
	btnclose := gtk.NewToolButtonFromStock(gtk.STOCK_CLEAR)

	separator := gtk.NewSeparatorToolItem()
	btncustom := gtk.NewToolButton(nil, "Custom")
	btnmenu := gtk.NewMenuToolButtonFromStock("gtk.STOCK_CLOSE")
	btnmenu.SetArrowTooltipText("This is a tool tip")

	/*
		btnnew.OnClicked(onToolButtonClicked)
		btnclose.OnClicked(onToolButtonClicked)
		btncustom.OnClicked(onToolButtonClicked)
	*/

	toolmenu := gtk.NewMenu()
	menuitem := gtk.NewMenuItemWithMnemonic("8")
	menuitem.Show()
	toolmenu.Append(menuitem)
	menuitem = gtk.NewMenuItemWithMnemonic("16")
	menuitem.Show()
	toolmenu.Append(menuitem)
	menuitem = gtk.NewMenuItemWithMnemonic("32")
	menuitem.Show()
	toolmenu.Append(menuitem)
	btnmenu.SetMenu(toolmenu)

	toolbar.Insert(btnnew, -1)
	toolbar.Insert(btnclose, -1)
	toolbar.Insert(separator, -1)
	toolbar.Insert(btncustom, -1)
	toolbar.Insert(btnmenu, -1)

	return toolbar
}

func newDrawingArea() gtk.IWidget {

	drawingArea := gtk.NewDrawingArea()
	drawingArea.SetSizeRequest(100, 100)
	//drawingArea.SetSensitive(true)
	drawingArea.SetCanFocus(true)

	allocation := drawingArea.GetAllocation()

	pixmap := gdk.NewPixmap(drawingArea.GetWindow().GetDrawable(), allocation.Width, allocation.Height, 24)
	gc := gdk.NewGC(pixmap.GetDrawable())

	dir := "/home/chander/work/langs/golang/path/src/cwork/src/mind_shift/images/theme3/64"
	nameActiveBlock := filepath.Join(dir, "active_block.png")

	imageActiveBlock := gtk.NewImageFromFile(nameActiveBlock)
	pixbufActiveBlock := imageActiveBlock.GetPixbuf()

	render := func() {

		drawable := pixmap.GetDrawable()

		//gc.SetRgbBgColor(gdk.NewColorRGB(255, 0, 0))
		gc.SetRgbFgColor(gdk.NewColorRGB(0x16, 0x05, 0x05))

		drawable.DrawRectangle(gc, true, 0, 0, -1, -1)

		drawable.DrawPixbuf(gc, pixbufActiveBlock, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
	}

	drawingArea.Connect("configure-event", func() {
		if pixmap != nil {
			pixmap.Unref()
		}

		allocation := drawingArea.GetAllocation()

		pixmap = gdk.NewPixmap(drawingArea.GetWindow().GetDrawable(), allocation.Width, allocation.Height, 24)
		gc = gdk.NewGC(pixmap.GetDrawable())

		render()
	})

	drawingArea.Connect("expose-event", func() {
		if pixmap != nil {
			drawingArea.GetWindow().GetDrawable().DrawDrawable(gc, pixmap.GetDrawable(), 0, 0, 0, 0, -1, -1)
		}
	})

	return drawingArea
}

func newTextView() gtk.IWidget {

	str := `{
	"Name": "Cross 2",
	"Author": "Chander",
	"Description": "",
	"Difficulty": 1,
		"Primitives": {
			"Cells": [
			"----456----",
			"----456----",
			"----456----",
			"----456----",
			"11115231111",
			"22224-52222",
			"33333163333",
			"----456----",
			"----456----",
			"----456----",
			"----456----"
		],
		"Identifiers": [
			{ "Id": "1", "Location": { "Start": [-14, -4], "Finish": [0, 0] } },
			{ "Id": "2", "Location": { "Start": [-14, 0], "Finish": [0, 0] } },
			{ "Id": "3", "Location": { "Start": [-14, 4], "Finish": [0, 0] } },
			{ "Id": "4", "Location": { "Start": [20, 0], "Finish": [0, 0] } },
			{ "Id": "5", "Location": { "Start": [14, 0], "Finish": [0, 0] } },
			{ "Id": "6", "Location": { "Start": [8, 0], "Finish": [0, 0] } }
		]
	}
}`

	textView := gtk.NewTextView()
	textView.ModifyFontEasy("Monospace 12")
	//textView.SetWrapMode(gtk.WRAP_WORD)

	textBuffer := textView.GetBuffer()
	textBuffer.SetText(str)

	scrolledWindow := gtk.NewScrolledWindow(nil, nil)
	scrolledWindow.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	scrolledWindow.SetShadowType(gtk.SHADOW_IN)
	scrolledWindow.Add(textView)
	scrolledWindow.SetSizeRequest(100, 100)

	vbox := gtk.NewVBox(false, 0)

	vbox.Add(scrolledWindow)

	//PackStart(newToolbar(), false, false, 3)
	//
	{
		/*
			textView := gtk.NewTextView()
			textView.ModifyFontEasy("Monospace 12")

			scrolledWindow := gtk.NewScrolledWindow(nil, nil)
			scrolledWindow.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
			scrolledWindow.SetShadowType(gtk.SHADOW_IN)
			scrolledWindow.Add(textView)
			scrolledWindow.SetSizeRequest(100, 100)

			vbox.Add(scrolledWindow)
		*/
	}

	return vbox
}
