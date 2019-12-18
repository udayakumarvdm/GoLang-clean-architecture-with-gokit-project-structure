package model

type Response struct {
	PersonID  uint   `gorm:"primary_key;column:personid"`
	LastName  string `gorm:"type:varchar(200);column:lastname;"`
	FirstName string `gorm:"type:varchar(200);column:firstname;"`
	Address   string `gorm:"type:varchar(200);column:address;"`
	City      string `gorm:"type:varchar(200);column:city;"`
}

func (Response) TableName() string {
	return "persons"
}
