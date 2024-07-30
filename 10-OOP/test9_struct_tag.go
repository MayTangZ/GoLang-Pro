package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// type resume struct {
// 	Name string  'info:"name" doc:"姓名"'
// 	Sex  string  'info:"Sex"'
// }

type resume struct {
	Name string `info:"name" doc:"姓名"`
	Sex  string `info:"Sex"`
}

type Moive struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
	Price int      `json:"rmb"`
	Actor []string `json:"actor"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagstring := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", tagstring, "doc:", tagDoc)

	}
}

func main() {
	var re resume

	findTag(&re)
	//fmt.Printf("%v\n", re)
	movie := Moive{"喜剧之王", 2000, 10, []string{"xingye", "zhangbozhi"}}

	//struct-->json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
		return

	}

	fmt.Printf("jsonStr=%s\n", jsonStr)

	//jsonStr ->struct
	//jsonStr:={"title":"喜剧之王","release":2000,"price":10,"actor":["xingye","zhangbozhi"]}
	myMovie := Moive{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json Unmarshal")
		return

	}

	fmt.Printf("%v\n", myMovie)

}
