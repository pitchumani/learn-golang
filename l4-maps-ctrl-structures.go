package main

import "fmt"
import "strings"
import "sort"

func map1() {
	users := map[string]string{}
	users["jason"] = "jason@example.com"
	fmt.Println(users["jason"])
}

// initialzing a map
func map2() {
	users := map[string]string{
		"kurt@example.com": "Kurt",
		"jason@example.com": "Jason",
		"akram@example.com": "Akram",
	}
	fmt.Println(users)
}

func map3() {
	var users map[string]string
	fmt.Println(len(users), users)

	// initialize the map - mandatory
	users = make(map[string]string)

	users["kurt@example.com"] = "Kurt"
	users["jason@example.com"] = "Jason"
	users["akram@example.com"] = "Akram"
	fmt.Println(len(users), users)
}

func map4() {
	var users map[string]string
	// runtime error as the map is not initialized
	users["joe@example.com"] = "Joe"
	fmt.Println(users)
}

func map5() {
	/*
	// not comparable type can't be  used as key
	var a map[func()]string
	var b map[map[int]int]string
	var c map[[]int]string

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// struct type can be used as key when all of its
	// members are of comparable types
	*/
}

func map6() {
	// iterating map using range
	users := map[string]string{
		"kurt@example.com": "Kurt",
		"jason@example.com": "Jason",
		"akram@example.com": "Akram",
	}

	for k, v := range users {
		fmt.Println("Email:", k, " Name:", v)
	}

	for key := range users {
		fmt.Println("Email:", key)
	}
}

func map7() {
	users := map[string]string{
		"Kurt": "kurt@example.com",
		"Jason": "jason@example.com",
		"Akram": "akram@example.com",
	}
	fmt.Println(users)
	delete(users, "Kurt")
	fmt.Println(users)
	non_existing_key := "Non existing key"
	value, ok := users[non_existing_key]
	if !ok {
		fmt.Printf("key %s doesn't exist in users\n", non_existing_key)
	} else {
		fmt.Println("%s:%s", non_existing_key, value)
	}
}

func map8() {
	// use map for counting words in a sentence
	counts := map[string]int{}

	sentence := "This is an example sentence. The sentence can be split into words using the strings package."
	words := strings.Fields(strings.ToLower(sentence))
	for _, w := range words {
		// ignored index, word's count is incremented
		//counts[w] = counts[w] + 1
		counts[w]++
	}
	fmt.Println(counts)
}

func map9() {
	// updating map in place can't be done when the key is complex type (e.g. struct)
	type User struct {
		ID int
		Name string
	}
	data := map[int]User{}
	user1 := User{ID: 1, Name: "Jason"}
	data[1] = user1
	//data[1].Name = "Akram"  -- error
	// retrieve value from map, update it and reinsert value into the map
	user2 := data[1]
	user2.Name = "Akram"
	data[1] = user2
	fmt.Printf("%+v\n", data)
}

func map10() {
	// take the keys from the map
	months := map[int]string{
		1: "Jan",
		2: "Feb",
		3: "Mar",
		4: "Apr",
		5: "May",
		6: "Jun",
		7: "Jul",
		8: "Aug",
		9: "Sep",
		10: "Oct",
		11: "Nov",
		12: "Dec",
	}

	keys := make([]int, 0, len(months))

	for k := range months {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	// iterating maps in Go is random, can't be in any particular order
	// retrieve keys to a slice, sort the slice and use the sorted slice
	// to retrieve the value from map in sorted order
	// sort function updates the slice in place
	sort.Ints(keys)
	fmt.Println(keys)
}
		
func main() {
	map1()
	map2()
	map3()
	//map4()
	//map5()
	map6()
	map7()
	map8()
	map9()
	map10()
}
