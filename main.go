package main

import (
	"log"
	"os"

	"github.com/FelipeFAlves/protobuf-example/contact"
	"google.golang.org/protobuf/proto"
	"fmt"
	// "google.golang.org/protobuf/types/known/timestamppb" // descomentar quando usar
)

func main() {
	//  p1 := &contact.Person{
	// 	Name:        "Felipe",
	// 	Id:          32,
	// 	Email:       "felipealvesferreira05@gmail.com",
	// 	LastUpdated: timestamppb.Now(),
	// }

	// p2 := &contact.Person{
	// 	Name:        "Felipe2",
	// 	Id:          33,
	// 	Email:       "felipealvesferreira05@gmail.com2",
	// 	LastUpdated: timestamppb.Now(),
	// }

	// addressBook := &contact.AddressBook{
	// 	People: []*contact.Person{
	// 		p1, p2,
	// 	},
	// }

	// out, err := proto.Marshal(addressBook)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = os.WriteFile("address.db", out, 0644)
	// if err != nil {
	// 	log.Fatalln(err)
	// }   

	//Essa parte anterior do código é para armazenar daquele .db

	data,err := os.ReadFile("address.db")
	if err != nil {
		log.Fatalln(err)
	}

	addressBook := &contact.AddressBook {}
	err = proto.Unmarshal(data,addressBook)

	if err != nil {
		log.Fatalln(err)
	}

	//No Errors

	for _, person := range addressBook.People {
		fmt.Println(person.Id, person.Name, person.Email)
	}
	
}
