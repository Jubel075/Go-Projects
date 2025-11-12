package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Go To-Do App")

	var tasks []string
	list := widget.NewList(
		func() int { return len(tasks) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tasks[i])
		},
	)

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter new task...")

	addButton := widget.NewButton("Add", func() {
		text := entry.Text
		if text != "" {
			tasks = append(tasks, text)
			entry.SetText("")
			list.Refresh()
		}
	})

	removeButton := widget.NewButton("Remove Selected", func() {
		selected := list.Selected
		if selected >= 0 && selected < len(tasks) {
			tasks = append(tasks[:selected], tasks[selected+1:]...)
			list.Refresh()
		}
	})

	content := container.NewVBox(
		entry,
		addButton,
		removeButton,
		list,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 400))
	myWindow.ShowAndRun()
}
