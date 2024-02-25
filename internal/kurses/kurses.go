package kurses

import (
	"log"
	"os"

	"github.com/gbin/goncurses"
)

const (
	HEIGHT = 40
	WIDTH  = 40
)

func Kurses() {
	var active int

	files, err := os.ReadDir(os.Getenv("HOME") + "/.kube")
	menu := []string{}
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Type().IsRegular() {
			menu = append(menu, file.Name())
		}
	}

	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer goncurses.End()

	goncurses.Raw(true)
	goncurses.Echo(false)
	goncurses.Cursor(0)
	stdscr.Clear()
	stdscr.Keypad(true)

	my, mx := stdscr.MaxYX()
	y, x := 2, (mx/2)-(WIDTH/2)

	win, _ := goncurses.NewWindow(HEIGHT, WIDTH, y, x)
	win.Keypad(true)

	stdscr.Print("Use TAB/BACKSPACE or UP/DOWN arrows or j/k to go up and down, Press enter to select \n")
	stdscr.Print("'q' to exit \n")
	stdscr.Refresh()

	printmenu(win, menu, active)

	for {
		ch := stdscr.GetChar()
		switch goncurses.Key(ch) {
		case 'q':
			return
		case 'k':
			if active == 0 {
				active = len(menu) - 1
			} else {
				active -= 1
			}

		case 'j':
			if active == len(menu)-1 {
				active = 0
			} else {
				active += 1
			}

		case goncurses.KEY_BACKSPACE:
			if active == 0 {
				active = len(menu) - 1
			} else {
				active -= 1
			}

		case goncurses.KEY_TAB:
			if active == len(menu)-1 {
				active = 0
			} else {
				active += 1
			}

		case goncurses.KEY_UP:
			if active == 0 {
				active = len(menu) - 1
			} else {
				active -= 1
			}

		case goncurses.KEY_DOWN:
			if active == len(menu)-1 {
				active = 0
			} else {
				active += 1
			}

		case goncurses.KEY_RETURN, goncurses.KEY_ENTER, goncurses.Key('\r'):
			stdscr.MovePrintf(my-2, 0, "#%d: %s",
				active,
				menu[active])
			if _, err := os.Lstat(os.Getenv("HOME") + "/.kube/config"); err == nil {
				os.Remove(os.Getenv("HOME") + "/.kube/config")
			}
			os.Symlink(os.Getenv("HOME")+"/.kube/"+menu[active], os.Getenv("HOME")+"/.kube/config")
			stdscr.ClearToEOL()
			stdscr.Refresh()
			// set this to return after a selection as opposed to having to press 'q'
			// return

		default:
			stdscr.MovePrintf(my-2, 0, "Character pressed = %3d/%c",
				ch, ch)
			stdscr.ClearToEOL()
			stdscr.Refresh()
		}

		printmenu(win, menu, active)
	}
}

func printmenu(w *goncurses.Window, menu []string, active int) {
	y, x := 2, 2
	w.Box(0, 0)
	for i, s := range menu {
		if i == active {
			w.AttrOn(goncurses.A_REVERSE)
			w.MovePrint(y+i, x, s)
			w.AttrOff(goncurses.A_REVERSE)
		} else {
			w.MovePrint(y+i, x, s)
		}
	}
	w.Refresh()
}
