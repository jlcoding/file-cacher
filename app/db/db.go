package db

import (
	"file-cacher/config"
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/gogf/gf/os/glog"
)

//type Db struct {
var Client *badger.DB

//}

func init() {
	open()
}

func open() {
	opts := badger.DefaultOptions(config.Config.BadgerDB.Dir)
	var err error
	Client, err = badger.Open(opts)
	if err != nil {
		glog.Error("open db error", err)
	} else {
		glog.Info("open db success!")
	}
}

func close() {
	err := Client.Close()
	if nil == err {
		glog.Error("close db error", err)
	}
}

func Update(key string, value string) {
	transaction := Client.NewTransaction(true)
	defer transaction.Discard()
	err := transaction.Set([]byte(key), []byte(value))
	if err == nil {
		_ = transaction.Commit()
	} else {
		glog.Error("update data set error", err)
	}
}

func Get(key string) string {
	transaction := Client.NewTransaction(false)
	defer transaction.Discard()
	if item, err := transaction.Get([]byte(key)); err == nil {
		value, err := item.ValueCopy(nil)
		if err == nil {
			fmt.Println(value)
			return string(value)
		}
	}
	return ""
}

func Delete(key string) {
	transaction := Client.NewTransaction(true)
	defer transaction.Discard()
	err := transaction.Delete([]byte(key))
	if err != nil {
		glog.Error("delete data fail", err)
	}
}

func List() {
	transaction := Client.NewTransaction(false)
	defer transaction.Discard()
	iter := badger.DefaultIteratorOptions
	it := transaction.NewIterator(iter)
	for it.Rewind(); it.Valid(); it.Next() {
		item := it.Item()
		fmt.Printf("key: %s\n", item.Key())
		value, _ := item.ValueCopy(nil)
		fmt.Printf("value: %s\n", value)
	}
}
