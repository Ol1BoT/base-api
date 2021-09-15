package db

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	FirstName string
	MiddleName string
	LastName string
	PrimaryEmail     string
	SecondaryEmail string
	DOB       *time.Time
	NetworkName string
	Address []Address
	Student Student
	StudentYear []StudentYear
	MedicalDetails MedicalDetails
	Immunisations []Immunisations
	Relationships []Person `gorm:"many2many:relationships;"`
	
	
}

type Student struct {
	PersonID uint
	EntryDate *time.Time
	House string
	TutorGroup string
	LeavingDate *time.Time
	StudentContact []Person
}

type Relationships struct {

	PersonID uint `gorm:"primaryKey"`
	RelationID uint `gorm:"primaryKey"`
	RelationType string
	LetRelationViewPersonMarks bool

}

type MedicalDetails struct {
	PersonID uint
	

}

type Immunisations struct {
	PersonID uint
	Name string
	For string
	Brand string
	AdministeredDate *time.Time
	ExpiryDate *time.Time

}

type MedicalConditions struct {
	PersonID uint

}


type StudentYear struct {
	PersonID uint
	Year uint16
	YearLevel int8
	
}

type PersonHistory struct {
	Person
}

type Address struct {
	PersonID uint
	AddressID uint `gorm:"primaryKey"`
	PostalAddress1 string
	PostalAddress2 string
	PostalAddress3 string
	PostalSuburb string
	PostalPostcode string
	HomeAddress1 string
	HomeAddress2 string
	HomeAddress3 string
	HomeSuburb string
	HomePostCode string

}

type Class struct {
	gorm.Model
	Year uint16
	Term uint8
	ClassCode string
	Description string
	ImportIntoLMS bool
	// MarkBook MarkBook
	AssessmentBook MarkBook

}

type MarkBook struct {
	gorm.Model
	MarkBookType string
	Marks []Mark
}

//Lets be able to import from Excel
type Mark struct {
	MarkBookID uint
	PersonID uint
	MarkName string
	MarkDescription string
	ScoreType string //Is it a Number or Letter 
	GradeScore string
	
}
