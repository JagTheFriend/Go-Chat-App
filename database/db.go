package database

type DB struct {
	StoredData map[string]any
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
	return db.StoredData[key]

}

func (db *DB) Set(key string, value any) {
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
