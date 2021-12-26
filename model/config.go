package model

type Tabler interface {
	TableName() string
}

func (Patient) TableName() string {
	return "patient"
}

func (Address) TableName() string {
	return "address"
}
