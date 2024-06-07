package main

import (
	"fmt"
	"strconv"
)

func main() {
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
}
