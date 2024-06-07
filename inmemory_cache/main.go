package main

import (
	"fmt"
	"github.com/nnno/golang-sandbox/inmemory_cache/cache"
	"strconv"
)

type Author struct {
	Id   int
	Name string
}

func main() {
	/**
	c := NewUserCache()

	row := User{ID: 1, Name: "test"}
	c.Set(strconv.Itoa(row.ID), row)
	if v, ok := c.Get(strconv.Itoa(row.ID)); ok {
		fmt.Println(v.ID)
	}

	row = User{ID: 2, Name: "hoge"}
	c.Set(strconv.Itoa(row.ID), row)

	row = User{ID: 3, Name: "fuga"}
	c.Set(strconv.Itoa(row.ID), row)

	for _, v := range c.GetAll() {
		fmt.Println(v.ID)
	}
	*/

	c := cache.New()
	row := Author{Id: 1, Name: "test"}
	c.Set(strconv.Itoa(row.Id), row)
	if v, ok := c.Get(strconv.Itoa(row.Id)); ok {
		fmt.Println(v.(Author).Id)
	}

	row = Author{Id: 2, Name: "hoge"}
	c.Set(strconv.Itoa(row.Id), row)
	row = Author{Id: 3, Name: "fuga"}
	c.Set(strconv.Itoa(row.Id), row)

	m := c.GetAll()
	for _, v := range m {
		fmt.Println(v.Object.(Author).Id, v.Object.(Author).Name)
	}

}
