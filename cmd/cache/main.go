package main

import (
	"cache/internal/list"
	"fmt"
	"os"
)

func main() {
	//c := cache.NewCombinedCache(3, 1*time.Second) // Почему нельзя передавать "time.Second"?
	//c.Add("key1", "val1")
	//c.Add("key2", "val2")
	//
	//res, ok := c.Get("key1")
	//fmt.Println("key1:", "Exists:", ok, "Value:", res)
	//
	//res, ok = c.Get("key2")
	//fmt.Println("key2:", "Exists:", ok, "Value:", res)
	//
	//time.Sleep(2 * time.Second)
	//
	//c.Add("key3", "val3")
	//
	//res, ok = c.Get("key1")
	//fmt.Println("key1:", "Exists:", ok, "Value:", res)
	//
	//res, ok = c.Get("key2")
	//fmt.Println("key2:", "Exists:", ok, "Value:", res)
	//
	//res, ok = c.Get("key3")
	//fmt.Println("key3:", "Exists:", ok, "Value:", res)

	l := list.NodeList{}

	l.Append("node1")
	l.Append("node2")
	l.Append("node3")

	l.Show(os.Stdout)

	l.Revert()

	l.Prepend("node4")
	l.Append("node5")

	l.Show(os.Stdout)

	fmt.Printf("%#v\n", l.Head)
	fmt.Printf("%#v\n", l.Tail)

	item, ok := l.Search("node5")
	if ok {
		fmt.Println(item)
	}
}
