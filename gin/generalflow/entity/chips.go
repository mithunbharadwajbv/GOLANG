package entity

type Company struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Chips struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	ParentCompany Company `json:"parentCompany"`
}

func GetSampleChips(id string) Chips {
	return Chips{
		ID:   id,
		Name: "lays",
		ParentCompany: Company{
			Name:     "Cola",
			Location: "Bangalore",
		},
	}
}
