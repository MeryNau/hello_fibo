package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func fibonacci(anzahl int) []int {
	f := make([]int, anzahl+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= anzahl; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f
}
func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")

}
func main() {
	a := app.New()
	w := a.NewWindow("mery")
	label := widget.NewLabel("Hello, gib eine Zahl ein")
	entry := widget.NewEntry()
	labelFibo := widget.NewLabel(" ")

	button := widget.NewButton("Fibonacci Array ausgeben", func() {
		data := entry.Text
		dataint, _ := strconv.Atoi(data)
		fmt.Println(fibonacci(dataint))
		fibo := fibonacci(dataint)
		labelFibo.SetText(arrayToString(fibo, ","))

	})

	w.SetContent(container.NewGridWithRows(4,
		label,
		entry,
		button,
		labelFibo,
	))

	w.ShowAndRun()
}
