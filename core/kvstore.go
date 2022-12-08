package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type KVStore struct {
	*Crypto

	filename     string
	cache        map[string]string
	encryptCache map[string]string
}

func NewKVStore(filename string) *KVStore {
	kv := &KVStore{
		Crypto:       new(Crypto),
		filename:     filename,
		cache:        make(map[string]string),
		encryptCache: make(map[string]string),
	}
	return kv
}

func (kv *KVStore) Get(key string) (string, bool) {
	value, ok := kv.cache[key]
	return value, ok
}

func (kv *KVStore) Update(key, value string) bool {
	if oldValue, ok := kv.cache[key]; ok {
		if oldValue == value {
			return true
		}
	}
	ev, err := kv.Encrypt(value)
	if err != nil {
		fmt.Printf("[ERROR] Encryption err: %s\n", err)
		return false
	}
	kv.cache[key] = value
	kv.encryptCache[key] = ev
	if err := kv.Write(); err != nil {
		fmt.Printf("[ERROR] KVStore write err: %s\n", err)
		return false
	}
	return true
}

func (kv *KVStore) Delete(key string) bool {
	if _, ok := kv.cache[key]; ok {
		delete(kv.cache, key)
		delete(kv.encryptCache, key)
		if err := kv.Write(); err != nil {
			fmt.Printf("[ERROR] KVStore write err: %s\n", err)
			return false
		}
	}
	return true
}

func (kv *KVStore) List() [][]string {
	var retList [][]string
	for k, v := range kv.cache {
		entry := []string{k, v}
		retList = append(retList, entry)
	}
	return retList
}

func (kv *KVStore) Read() error {
	rawKV, err := ioutil.ReadFile(kv.filename)
	if err != nil {
		return err
	}
	kv.encryptCache = make(map[string]string)
	if err = json.Unmarshal(rawKV, &kv.encryptCache); err != nil {
		return err
	}
	kv.cache = make(map[string]string)
	for k, v := range kv.encryptCache {
		if dv, err := kv.Decrypt(v); err != nil {
			return err
		} else {
			kv.cache[k] = dv
		}
	}
	return nil
}

func (kv *KVStore) Write() error {
	dat, err := json.Marshal(kv.encryptCache)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(kv.filename, dat, 0644)
	return err
}
