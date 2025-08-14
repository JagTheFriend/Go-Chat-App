package database

import "sync"

type DB struct {
	StoredData map[string]any
	rw         sync.RWMutex
}

type ChannelType struct {
	Type string
	Data map[string]any
}

var db *DB
var DbChan = make(chan (ChannelType))

func NewDB() *DB {
	if db == nil {
		db = &DB{
			StoredData: make(map[string]any),
		}
	}
	return db
}

func (db *DB) Get(key string) any {
	db.rw.RLock()
	defer db.rw.RUnlock()
	return db.StoredData[key]

}

func (db *DB) Set(key string, value any) {
	db.rw.Lock()
	defer db.rw.Unlock()

	DbChan <- ChannelType{
		Type: "CREATE",
		Data: map[string]any{
			"key":   key,
			"value": value,
		},
	}
	db.StoredData[key] = value
}

func (db *DB) Delete(key string) {
	db.rw.Lock()
	defer db.rw.Unlock()

	value := db.StoredData[key]
	DbChan <- ChannelType{
		Type: "DELETE",
		Data: map[string]any{
			"key":   key,
			"value": value,
		},
	}
	delete(db.StoredData, key)
}
