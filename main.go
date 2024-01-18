// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"io/ioutil"
	"log"
	"os"

	. "github.com/gbin/goncurses"
)

 const (
	 HEIGHT = 40
	 WIDTH  = 40
 )

 func main() {
	 var active int

	 files, err := ioutil.ReadDir(os.Getenv("HOME")+"/.kube")
	 //var filearr []string
	 menu := []string{}
	 if err != nil {
			 log.Fatal(err)
	 }

	 for _, file := range files {
		 if file.Mode().IsRegular() {
			 menu = append(menu, file.Name())
		 }
	 }

	 stdscr, err := Init()
	 if err != nil {
		 log.Fatal(err)
	 }
	 defer End()

	 Raw(true)
	 Echo(false)
	 Cursor(0)
	 stdscr.Clear()
	 stdscr.Keypad(true)

	 my, mx := stdscr.MaxYX()
	 y, x := 2, (mx/2)-(WIDTH/2)

	 win, _ := NewWindow(HEIGHT, WIDTH, y, x)
	 win.Keypad(true)

	 stdscr.Print("Use TAB/BACKSPACE or UP/DOWN arrows or j/k to go up and down, Press enter to select \n")
	 stdscr.Print("'q' to exit \n")
	 stdscr.Refresh()

	 printmenu(win, menu, active)

	 for {
		 ch := stdscr.GetChar()
		 switch Key(ch) {
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
		 case KEY_BACKSPACE:
			 if active == 0 {
				 active = len(menu) - 1
			 } else {
				 active -= 1
			 }
		 case KEY_TAB:
			 if active == len(menu)-1 {
				 active = 0
			 } else {
				 active += 1
			 }
		 case KEY_UP:
			 if active == 0 {
				 active = len(menu) - 1
			 } else {
				 active -= 1
			 }
		 case KEY_DOWN:
			 if active == len(menu)-1 {
				 active = 0
			 } else {
				 active += 1
			 }
		 case KEY_RETURN, KEY_ENTER, Key('\r'):
			 stdscr.MovePrintf(my-2, 0, "#%d: %s",
				 active,
				 menu[active])
				 if _, err := os.Lstat(os.Getenv("HOME")+"/.kube/config"); err == nil {
					os.Remove(os.Getenv("HOME")+"/.kube/config")
				}
				os.Symlink(os.Getenv("HOME")+"/.kube/"+menu[active], os.Getenv("HOME")+"/.kube/config")
			 stdscr.ClearToEOL()
			 stdscr.Refresh()
		 default:
			 stdscr.MovePrintf(my-2, 0, "Character pressed = %3d/%c",
				 ch, ch)
			 stdscr.ClearToEOL()
			 stdscr.Refresh()
		 }

		 printmenu(win, menu, active)
	 }
 }

 func printmenu(w *Window, menu []string, active int) {
	 y, x := 2, 2
	 w.Box(0, 0)
	 for i, s := range menu {
		 if i == active {
			 w.AttrOn(A_REVERSE)
			 w.MovePrint(y+i, x, s)
			 w.AttrOff(A_REVERSE)
		 } else {
			 w.MovePrint(y+i, x, s)
		 }
	 }
	 w.Refresh()
 }
