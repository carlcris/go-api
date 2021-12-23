package model

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Patient struct {
	PatientID string //gorm convert field name to snake_case. So PatientID become patient_id
	FirstName string //first_name
	LastName  string //last_name
	DOB       string
}
type Address struct {
	Address1 string
	City     string
	State    string
	Zip      string
}

func (p *Patient) savePatient() (*Patient, error) {
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

func GetUserByID(c *gin.Context) {
	var p Patient

	uid := c.Param("id")
	log.Println(uid)
	if db, err := CreateDatabaseConnection(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else {
		if err := db.Find(&p, "patient_id = ?", uid).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		log.Println(p)

		c.JSON(http.StatusOK, gin.H{"message": "success", "data": p})
	}
}

func GetPatientList(c *gin.Context) {

	var patients []Patient
	db, err := CreateDatabaseConnection()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err = db.Find(&patients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": patients})
}
