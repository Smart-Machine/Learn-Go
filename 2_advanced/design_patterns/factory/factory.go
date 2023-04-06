package main

import (
	"fmt"
	"reflect"
)

type (
	mongoDB struct {
		database map[string]string
	}

	sqlite struct {
		database map[string]string
	}

	Database interface {
		GetData(string) string
		PutData(string, string)
	}
)

func (mdb mongoDB) GetData(query string) string {
	if _, ok := mdb.database[query]; !ok {
		return ""
	}

	fmt.Println("MongoDB")
	return mdb.database[query]
}

func (mdb mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (sql sqlite) GetData(query string) string {
	if _, ok := sql.database[query]; !ok {
		return ""
	}

	fmt.Println("Sqlite")
	return sql.database[query]
}

func (sql sqlite) PutData(query string, data string) {
	sql.database[query] = data
}

func DatabaseFactory(env string) Database {
	switch env {
	case "production":
		return mongoDB{
			database: make(map[string]string),
		}
	case "development":
		return sqlite{
			database: make(map[string]string),
		}
	default:
		return nil
	}
}

func main() {
	env1 := "production"
	env2 := "development"

	db1 := DatabaseFactory(env1)
	db2 := DatabaseFactory(env2)

	db1.PutData("test", "this is mongodb data")
	fmt.Println(db1.GetData("test"))

	db2.PutData("test", "this is sqlite data")
	fmt.Println(db2.GetData("test"))

	fmt.Println(reflect.TypeOf(db1).Name())
	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(db2).Name())
	fmt.Println(reflect.TypeOf(&db2).Elem())
}
