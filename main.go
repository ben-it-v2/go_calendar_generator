package main

import(
	"os"
	"calendar/calendar_gui"
)

func main() {
	calendar_gui.CreateMainWindow(len(os.Args), os.Args)
}
