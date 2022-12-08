package interact

import (
	"fmt"
	"strings"
)

const DefaultFile string = "records.json"

func EntryMemu() (string, string) {
	var sk, filename string
	fmt.Printf("[HINT] Enter OnePassword to generate sk: ")
	fmt.Scanln(&sk)
	fmt.Printf("[HINT] Enter filename for records (default '%s', empty line to skip): ", DefaultFile)
	fmt.Scanf("%s", &filename)
	if len(filename) == 0 {
		filename = DefaultFile
	}
	return sk, filename
}

func MainMemu(skhash string, appname []string) {
	fmt.Printf("[INFO] Your sk hash is starting with: <0x%s>, make sure it is correct\n", skhash[:5])
	fmt.Println("[INFO] Current recorded apps:")
	if len(appname) == 0 {
		fmt.Println("\tNone")
	} else {
		nameList := strings.Join(appname, ",")
		fmt.Println("\t[" + nameList + "]")
	}
}

func ChoiceMenu() string {
	var choice string
	fmt.Println("[HINT] You wanna:")
	fmt.Println("\tl) list. f) find. a) add. u) update. d) delete. q) quit.")
	fmt.Scanln(&choice)
	return choice
}

func FindHint() string {
	var name string
	fmt.Printf("[HINT] Enter app's name to find: ")
	fmt.Scanln(&name)
	return name
}

func AddHint() (string, string) {
	var name, pwd, yn string
	fmt.Printf("[HINT] Enter app's name to add: ")
	fmt.Scanln(&name)
	fmt.Printf("[HINT] Enter app's password: ")
	fmt.Scanln(&pwd)
	fmt.Printf("[HINT] Make sure to add? Enter y or n: ")
	fmt.Scanln(&yn)
	if len(yn) > 0 && yn[0] == 'y' {
		return name, pwd
	}
	return "", ""
}

func UpdateHint() (string, string) {
	var name, pwd, yn string
	fmt.Printf("[HINT] Enter app's name to update: ")
	fmt.Scanln(&name)
	fmt.Printf("[HINT] Enter app's password: ")
	fmt.Scanln(&pwd)
	fmt.Printf("[HINT] Make sure to update? Enter y or n: ")
	fmt.Scanln(&yn)
	if len(yn) > 0 && yn[0] == 'y' {
		return name, pwd
	}
	return "", ""
}

func DeleteHint() string {
	var name, yn string
	fmt.Printf("[HINT] Enter app's name to delete: ")
	fmt.Scanln(&name)
	fmt.Printf("[HINT] Make sure to delete? Enter y or n: ")
	fmt.Scanln(&yn)
	if len(yn) > 0 && yn[0] == 'y' {
		return name
	}
	return ""
}

// TODO: MODIFY ENTRY PASSWORD OF ONE-PASSWORD
