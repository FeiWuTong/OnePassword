package service

import (
	"fmt"
	"onepassword/core"
)

type Manager struct {
	kv *core.KVStore
	fp string
}

func NewManager(filename string) *Manager {
	return &Manager{
		kv: core.NewKVStore(),
		fp: filename,
	}
}

func (m *Manager) SetSK(sk string) {
	m.kv.KeyGen(sk)
}

func (m *Manager) Load() bool {
	if err := m.kv.Read(m.fp); err != nil {
		fmt.Printf("[ERROR] Fail to load file: %s", err)
		return false
	} else {
		fmt.Println("[INFO] Successfully load file")
		return true
	}
}

func (m *Manager) Save() bool {
	if err := m.kv.Write(m.fp); err != nil {
		fmt.Printf("[ERROR] KVStore write err: %s\n", err)
		return true
	} else {
		fmt.Println("[INFO] KVStore write ok")
		return true
	}
}

func (m *Manager) GetSKHash() string {
	return m.kv.KeyHash()
}

func (m *Manager) ListName() []string {
	dataList := m.kv.List()
	nameList := make([]string, len(dataList))
	for i, entry := range dataList {
		nameList[i] = entry[0]
	}
	return nameList
}

func (m *Manager) List() {
	dataList := m.kv.List()
	fmt.Println("[INFO] List all records")
	fmt.Printf("%-20s|%-20s|%-20s|%-20s\n", " App Name", " Username", " Password", " Notes")
	fmt.Println("-------------------------------------------------------------------------------------")
	for _, entry := range dataList {
		fmt.Printf(" %-19s| %-19s| %-19s| %-19s\n", entry[0], entry[1], entry[2], entry[3])
	}
}

func (m *Manager) Find(key string) {
	if v, ok := m.kv.Get(key); !ok {
		fmt.Printf("[WARN] App (%s) not found, please re-check the name\n", key)
	} else {
		fmt.Println("[INFO] a record is found")
		fmt.Printf("%-20s|%-20s|%-20s|%-20s\n", "\tApp Name", "\tUsername", "\tPassword", "\tNotes")
		fmt.Println("-------------------------------------------------------------------------------------")
		fmt.Printf(" %-19s| %-19s| %-19s| %-19s\n", key, v.Username, v.Password, v.Notes)
	}
}

func (m *Manager) Add(key string, value []string) {
	if _, ok := m.kv.Get(key); ok {
		fmt.Printf("[ERROR] App (%s) already existed\n", key)
		return
	}
	account := core.NewAccount(value)
	if account == core.EmptyAccount {
		return
	}
	if ok := m.kv.Update(key, account); !ok {
		fmt.Printf("[ERROR] App (%s) has not been added\n", key)
		return
	} else {
		fmt.Printf("[INFO] App (%s) has been added\n", key)
	}
	m.Save()
}

func (m *Manager) Update(key string, value []string) {
	account := core.NewAccount(value)
	if account == core.EmptyAccount {
		return
	}
	if ok := m.kv.Update(key, account); !ok {
		fmt.Printf("[ERROR] App (%s) has not been updated\n", key)
		return
	} else {
		fmt.Printf("[INFO] App (%s) has been updated\n", key)
	}
	m.Save()
}

func (m *Manager) Delete(key string) {
	if ok := m.kv.Delete(key); !ok {
		fmt.Printf("[ERROR] App (%s) has not been deleted\n", key)
		return
	} else {
		fmt.Printf("[INFO] App (%s) has been deleted\n", key)
	}
	m.Save()
}
