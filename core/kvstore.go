package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var EmptyAccount Account

type Account struct {
	Username string
	Password string
	Notes    string
}

func NewAccount(aInfo []string) Account {
	if len(aInfo) < 3 {
		return Account{}
	}
	return Account{
		Username: aInfo[0],
		Password: aInfo[1],
		Notes:    aInfo[2],
	}
}

type KVStore struct {
	*Crypto

	cache        map[string]Account
	encryptCache map[string]Account
}

func NewKVStore() *KVStore {
	kv := &KVStore{
		Crypto:       new(Crypto),
		cache:        make(map[string]Account),
		encryptCache: make(map[string]Account),
	}
	return kv
}

func (kv *KVStore) Get(key string) (Account, bool) {
	value, ok := kv.cache[key]
	return value, ok
}

func (kv *KVStore) Update(key string, value Account) bool {
	ep, err := kv.Encrypt(value.Password)
	if err != nil {
		fmt.Printf("[ERROR] Encryption err: %s\n", err)
		return false
	}
	kv.cache[key] = value
	ev := value
	ev.Password = ep
	kv.encryptCache[key] = ev
	return true
}

func (kv *KVStore) Delete(key string) bool {
	if _, ok := kv.cache[key]; ok {
		delete(kv.cache, key)
		delete(kv.encryptCache, key)
	}
	return true
}

func (kv *KVStore) List() [][]string {
	var retList [][]string
	for k, v := range kv.cache {
		entry := []string{k, v.Username, v.Password, v.Notes}
		retList = append(retList, entry)
	}
	return retList
}

func (kv *KVStore) Read(filename string) error {
	rawKV, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	kv.encryptCache = make(map[string]Account)
	if err = json.Unmarshal(rawKV, &kv.encryptCache); err != nil {
		return err
	}
	kv.cache = make(map[string]Account)
	for k, v := range kv.encryptCache {
		if dp, err := kv.Decrypt(v.Password); err != nil {
			return err
		} else {
			dv := v
			dv.Password = dp
			kv.cache[k] = dv
		}
	}
	return nil
}

func (kv *KVStore) Write(filename string) error {
	dat, err := json.Marshal(kv.encryptCache)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, dat, 0644)
	return err
}
