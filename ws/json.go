package ws

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Message struct {
	Name string
	Body string
	Time int64
}
type jsonObject struct {
	Object ObjectType
}

type ObjectType struct {
	Buffer_size int
	Database    []DatabaseType
}

type DatabaseType struct {
	Host   string
	User   string
	Pass   string
	Type   string
	Name   string
	Tables []TablesType
}

type TablesType struct {
	Name     string
	Statment string
	Regex    string
	Types    []TypesType
}

type TypesType struct {
	Id    string
	Value string
}

func Json(message string) string {
	file, err := ioutil.ReadFile("./ws/message.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	} else {
		log.Printf("%s\n", string(file))
		var jsontype jsonObject
		json.Unmarshal(file, &jsontype)
		log.Printf("Result: %v\n", jsontype)
		// log.Printf("Name = %s\n", result.Name)
		// log.Printf("Body = %s\n", result.Body)
		// log.Printf("Time = %d\n", result.Time)
	}
	return string(file)
}
