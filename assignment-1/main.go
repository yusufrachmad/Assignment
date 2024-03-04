package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Id         string
	Name       string
	Address    string
	Occupation string
	Reason     string
}

var friends []Friend = []Friend{
	{
		Id:         "C1",
		Name:       "John",
		Address:    "Jalan Bangka",
		Reason:     "Saya suka golang",
		Occupation: "SE",
	},
	{
		Id:         "C2",
		Name:       "Andrew",
		Address:    "Jalan Mampang",
		Reason:     "Saya pengen belajar golang",
		Occupation: "SE",
	},
	{
		Id:         "C3",
		Name:       "Matthew",
		Address:    "Jalan Pegangsaan Timur",
		Reason:     "Saya ingin kerja di google",
		Occupation: "QA",
	},
	{
		Id:         "C4",
		Name:       "Joko",
		Address:    "Jalan Sudirman",
		Reason:     "Saya ingin Indonesia maju dengan Golang",
		Occupation: "DS",
	},
	{
		Id:         "C5",
		Name:       "Gibran",
		Address:    "Jalan Gatot Subroto",
		Reason:     "Saya ingin menjadi presiden dengan Golang",
		Occupation: "DevOps",
	},
}

func searchFriendById(id string) {
	for _, v := range friends {
		if v.Id == id {
			fmt.Printf("id: %s\n", v.Id)
			fmt.Printf("name: %s\n", v.Name)
			fmt.Printf("address: %s\n", v.Address)
			fmt.Printf("occupation: %s\n", v.Occupation)
			fmt.Printf("reason: %s\n", v.Reason)
			return
		}
	}

	fmt.Printf("friend with id %s can not be found\n", id)
}

func main() {
	var arguments []string = os.Args

	if len(arguments) < 2 {
		fmt.Println("Tolong tambahkan argumen id")
	} else {
		searchFriendById(arguments[1])
	}
}
