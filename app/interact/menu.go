package interact

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const DefaultFile string = "records.json"

func EntryMemu() (string, string) {
	var sk, filename string
	fmt.Printf("[HINT] Enter OnePassword to generate sk: ")
	fmt.Scanln(&sk)
	fmt.Printf("[HINT] Enter filename for records (default '%s', empty line to skip): ", DefaultFile)
	fmt.Scanf("%s\n", &filename)
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

func AddHint() (string, []string) {
	var app, yn string
	info := make([]string, 3)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("[HINT] Enter app's name to add: ")
	app, _ = reader.ReadString('\n')
	app = strings.TrimSpace(app)
	if len(app) == 0 {
		fmt.Println("[ERROR] app's name cannot be empty")
		return "", []string{}
	}
	fmt.Printf("[HINT] Enter app's username (empty is ok): ")
	fmt.Scanln(&info[0])
	fmt.Printf("[HINT] Enter app's password: ")
	fmt.Scanln(&info[1])
	if len(info[1]) == 0 {
		fmt.Println("[ERROR] password cannot be empty")
		return "", []string{}
	}
	fmt.Printf("[HINT] Enter app's notes (empty is ok): ")
	info[2], _ = reader.ReadString('\n')
	info[2] = strings.TrimSpace(info[2])
	fmt.Printf("[HINT] Make sure to add? Enter y or n: ")
	fmt.Scanln(&yn)
	if len(yn) > 0 && yn[0] == 'y' {
		return app, info
	}
	return "", []string{}
}

func UpdateHint() (string, []string) {
	var app, yn string
	info := make([]string, 3)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("[HINT] Enter app's name to update: ")
	app, _ = reader.ReadString('\n')
	app = strings.TrimSpace(app)
	if len(app) == 0 {
		fmt.Println("[ERROR] app's name cannot be empty")
		return "", []string{}
	}
	fmt.Printf("[HINT] Enter app's username (empty is ok): ")
	fmt.Scanln(&info[0])
	fmt.Printf("[HINT] Enter app's password: ")
	fmt.Scanln(&info[1])
	if len(info[1]) == 0 {
		fmt.Println("[ERROR] password cannot be empty")
		return "", []string{}
	}
	fmt.Printf("[HINT] Enter app's notes (empty is ok): ")
	info[2], _ = reader.ReadString('\n')
	info[2] = strings.TrimSpace(info[2])
	fmt.Printf("[HINT] Make sure to update? Enter y or n: ")
	fmt.Scanln(&yn)
	if len(yn) > 0 && yn[0] == 'y' {
		return app, info
	}
	return "", []string{}
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
