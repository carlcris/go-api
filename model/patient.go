package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Patient struct {
	PatientId string
	FirstName string
	LastName  string
	MI        string
	BirthDate string
	Email     string
	Phone     string
	SSN       string
	Address   Address `gorm:"embedded"`
}
type Address struct {
	AddressId string
	Address1  string
	City      string
	State     string
	Zip       string
}

func (p *Patient) SavePatient() (*Patient, error) {

	db, err := CreateDatabaseConnection()

	if err != nil {
		return &Patient{}, err
	}

	err = db.Create(&p).Error
	if err != nil {
		return &Patient{}, err
	}

	return p, nil
}

func GetPatientAddressById(c *gin.Context) {
	var a Address

	uid := c.Param("id")
	if db, err := CreateDatabaseConnection(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else {
		if err = db.Find(&a, "address_id = ?", uid).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": a})
	}
}

func GetPatientByID(c *gin.Context) {
	var p Patient

	uid := c.Param("id")
	if db, err := CreateDatabaseConnection(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else {
		result := db.Table("patient").
			Select("patient.patient_id, patient.first_name, patient.last_name, patient.mi, patient.birth_date, patient.ssn, patient.phone, patient.email, address.address_id, address.address1, address.city, address.state, address.zip").
			Joins("JOIN address ON patient.patient_id = address.patient_id").
			Where("patient.patient_id =?", uid).
			Find(&p)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func GetPatientList(c *gin.Context) {

	var p []Patient
	db, err := CreateDatabaseConnection()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := db.Table("patient").
		Select("patient.patient_id, patient.first_name, patient.last_name, patient.mi, patient.birth_date, patient.ssn, patient.phone, patient.email, address.address_id, address.address1, address.city, address.state, address.zip").
		Joins("JOIN address ON patient.patient_id = address.patient_id").
		Find(&p)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}
func GetPatientAddress(c *gin.Context) {

	var p Patient
	db, err := CreateDatabaseConnection()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := db.Table("patient").
		Select("patient.patient_id, patient.first_name, patient.last_name, patient.dob, address.address_id, address.address1, address.city, address.state, address.zip").
		Joins("JOIN address ON patient.patient_id = address.patient_id").
		Find(&p)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": p})

}
