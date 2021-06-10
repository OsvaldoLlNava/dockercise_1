package Dockercise1APIModel

type Person struct {
	Id           uint   `xml:"id" json:"id" bson:"id"`
	First_Name   string `xml:"first_name" json:"first_name" bson:"first_name"`
	Last_Name    string `xml:"last_name" json:"last_name" bson:"last_name"`
	Company      string `xml:"company" json:"company" bson:"company"`
	Email        string `xml:"email" json:"email" bson:"email"`
	Ip_Address   string `xml:"ip_address" json:"ip_address" bson:"ip_address"`
	Phone_Number string `xml:"phone_number" json:"phone_number" bson:"phone_number"`
}

type People struct {
	Personas []Person `xml:"person" json:"person" bson:"person"`
}
