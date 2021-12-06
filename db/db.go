//DB 用于与数据库通过json通信的包
package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//People，从数据库中取到的用户信息，当用户有更多的信息的时候，只用修改People类型，然后在DSL编写的文件中，添加正确的$值即可
type People struct {
	Name   string
	Amount string
}

//Peoples 用于存每个用户的信息
var Peoples []People

func init() {
	content, err := ioutil.ReadFile("./db/data.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &Peoples)
	if err != nil {
		log.Fatal(err)
	}
}
