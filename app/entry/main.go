package main

import (
	"fmt"
	"onepassword/app/interact"
	"onepassword/service"
	"os"
)

func checkFile(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		if os.IsExist(err) {
			if fi.IsDir() {
				return false
			}
			return true
		}
		return false
	}
	return true
}

func workLoop() {
	skgen, filename := interact.EntryMemu()
	m := service.NewManager(filename)
	m.SetSK(skgen)
	if len(filename) == 0 {
		filename = interact.DefaultFile
	}
	if checkFile(filename) {
		fmt.Printf("[INFO] Loading file (%s)...\n", filename)
		if !m.Load() {
			return
		}
	} else {
		fmt.Printf("[INFO] File (%s) doesn't exist, starts new recording\n", filename)
	}

	interact.MainMemu(m.GetSKHash(), m.ListName())

	for {
		fmt.Println("")
		choice := interact.ChoiceMenu()
		switch choice {
		case "l":
			m.List()
		case "f":
			k := interact.FindHint()
			m.Find(k)
		case "a":
			k, v := interact.AddHint()
			m.Add(k, v)
		case "u":
			k, v := interact.UpdateHint()
			m.Update(k, v)
		case "d":
			k := interact.DeleteHint()
			m.Delete(k)
		case "q":
			return
		default:
			fmt.Println("[WARN] You should enter a valid choice")
		}
	}
}

func main() {
	workLoop()
}
