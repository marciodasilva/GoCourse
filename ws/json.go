package ws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ResponseA struct {
	Page   int
	Fruits []string
}

type ResponseB struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

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
func Responses() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(1.23)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("go gopher")
	fmt.Println(string(strB))

	slcD := []string{"Apple", "Pinaple", "Mango"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"Apple": 5, "Guava": 10}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &ResponseA{
		Page:   1,
		Fruits: []string{"Apple", "Peach", "Pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Printf("The string(res1B) value is = %s\n\n", string(res1B))

	res2D := &ResponseB{
		Page:   2,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Printf("The string(res2B) value is = %s\n\n", string(res2B))

	fmt.Println("-----------Unmarshal ResponseB -----------")
	str := `{"page":1, "fruits": ["apple","mango"]}`
	res := &ResponseB{}
	json.Unmarshal([]byte(str), &res)
	fmt.Printf("The res value is = %v\n\n", res)
	fmt.Printf("The res.Fruits value is = %s\n\n", res.Fruits[1])

	fmt.Println("-----------Decoding JSON data to GO value-----------")
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		log.Println(err)
	}
	fmt.Printf("The dat value is = %v\n\n", dat)
	// fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Printf("dat[num] float64 value is = %g\n\n", num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Printf("dat[num] string value is = %s\n\n", str1)
}
