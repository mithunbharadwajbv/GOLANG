package entity

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetSamplePerson(id string) Person {
	return Person{ID: id, Name: "Mithun"}
}
