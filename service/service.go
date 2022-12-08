package service

import (
	"fmt"
	"onepassword/core"
)

type Manager struct {
	kv *core.KVStore
}

func NewManager(filename string) *Manager {
	return &Manager{
		kv: core.NewKVStore(filename),
	}
}

func (m *Manager) SetSK(sk string) {
	m.kv.KeyGen(sk)
}

func (m *Manager) Load() bool {
	if err := m.kv.Read(); err != nil {
		fmt.Printf("[ERROR] Fail to load file: %s", err)
		return false
	} else {
		fmt.Println("[INFO] Successfully load file")
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
	fmt.Printf("%-20s|%-20s\n", "\tApp Name", "\tPassword")
	fmt.Println("-----------------------------------------------")
	for i := 0; i < len(dataList); i++ {
		fmt.Printf("%-20s|%-20s\n", "\t"+dataList[i][0], "\t"+dataList[i][1])
	}
}

func (m *Manager) Find(key string) {
	if v, ok := m.kv.Get(key); !ok {
		fmt.Printf("[WARN] App (%s) not found, please re-check the name\n", key)
	} else {
		fmt.Println("[INFO] a record is found")
		fmt.Printf("%-20s|%-20s\n", "\tApp Name", "\tPassword")
		fmt.Println("-----------------------------------------------")
		fmt.Printf("%-20s|%-20s\n", "\t"+key, "\t"+v)
	}
}

func (m *Manager) Add(key, value string) {
	if ok := m.kv.Update(key, value); !ok {
		fmt.Printf("[ERROR] App (%s) has not been added\n", key)
	} else {
		fmt.Printf("[INFO] App (%s) has been added\n", key)
	}
}

func (m *Manager) Update(key, value string) {
	if ok := m.kv.Update(key, value); !ok {
		fmt.Printf("[ERROR] App (%s) has not been updated\n", key)
	} else {
		fmt.Printf("[INFO] App (%s) has been updated\n", key)
	}
}

func (m *Manager) Delete(key string) {
	if ok := m.kv.Delete(key); !ok {
		fmt.Printf("[ERROR] App (%s) has not been deleted\n", key)
	} else {
		fmt.Printf("[INFO] App (%s) has been deleted\n", key)
	}
}
