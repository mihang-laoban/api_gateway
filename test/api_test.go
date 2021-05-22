package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var (
	a   Str
	str = `{
		"data": "this"
	}`
)

type Str struct {
	Data string
}

func marshal() {
	b := Str{Data: "that"}
	c, _ := json.Marshal(b)
	fmt.Println(string(c))
}

func unmarshal(str string) {
	data := []byte(str)
	if err := json.Unmarshal(data, &a); err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.Data)
}

func decode(str []byte) {
	if err := json.Unmarshal(str, &a); err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

func TestReq(t *testing.T) {
	//marshal()
	//unmarshal()
	data := strings.NewReader(str)
	res, err := http.Post("http://localhost:5000/this", "application/json; charset=utf-8", data)
	if err != nil || res.StatusCode != http.StatusOK {
		//c.Status(http.StatusServiceUnavailable)
		return
	}

	defer res.Body.Close()
	bodyByte, _ := ioutil.ReadAll(res.Body)
	decode(bodyByte)
	fmt.Println(string(bodyByte))
}
