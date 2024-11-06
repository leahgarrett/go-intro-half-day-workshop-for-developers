package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Id        float64 `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Address   string  `json:"address"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Skill     string  `json:"skill"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("mock_data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Opened mock_data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)
	fmt.Println("Unmarshaled users", len(users.Users))
	fmt.Println()

	fmt.Println("Users with software skill:")
	for i := 0; i < len(users.Users); i++ {
		if strings.Contains(strings.ToLower(users.Users[i].Skill), "software") {
			fmt.Println("User id#", users.Users[i].Id, "Name:", users.Users[i].FirstName, users.Users[i].LastName)
		}
	}

	fmt.Println("\nUsers from Australia:")
	for i := 0; i < len(users.Users); i++ {
		if strings.ToLower(users.Users[i].Country) == "australia" {
			fmt.Println("User id#", users.Users[i].Id, "Name:", users.Users[i].FirstName, users.Users[i].LastName)
		}
	}
}
