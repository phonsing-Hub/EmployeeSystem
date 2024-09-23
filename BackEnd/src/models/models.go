package models

import (
	"gorm.io/gorm"
	"time"
)

// AuthUser model for authentication
type AuthUser struct {
	gorm.Model
	Email            string     `gorm:"type:varchar(100);unique;not null"`
	Password         string     `gorm:"type:varchar(255);not null"` // Password should be hashed
	LastLogin        *time.Time // Track the last login time
	Role             string     `gorm:"type:varchar(20);default:admin"` // Default role is admin
	ResetToken       string     `gorm:"type:varchar(255)"`              // For password reset functionality
	ResetTokenExpiry *time.Time // For tracking token expiration
}



// Token model for managing JWT tokens
type Token struct {
	gorm.Model
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"type:text;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	User      AuthUser  `gorm:"foreignKey:UserID"`
}

// Department model
type Department struct {
	gorm.Model
	DepartmentName   string                     `gorm:"type:varchar(30);not null"`
	CurrentManagerID int                        `gorm:"index"`
	Employees        []Employee                 `gorm:"foreignKey:DepartmentID"`
	ManagerHistories []DepartmentManagerHistory `gorm:"foreignKey:DepartmentID"`
}

// Employee model
type Employee struct {
	gorm.Model
	FirstName        string                     `gorm:"type:varchar(30);not null"`
	LastName         string                     `gorm:"type:varchar(30);not null"`
	Email            string                     `gorm:"type:varchar(30);not null"`
	PhoneNumber      string                     `gorm:"type:varchar(12)"`
	HireDate         time.Time                  `gorm:"not null"`
	DepartmentID     uint                       `gorm:"index"`
	Department       Department                 `gorm:"foreignKey:DepartmentID"`
	Salaries         []Salary                   `gorm:"foreignKey:EmployeeID"`
	JobHistories     []JobHistory               `gorm:"foreignKey:EmployeeID"`
	ManagerHistories []DepartmentManagerHistory `gorm:"foreignKey:EmployeeID"`
}

// Salary model
type Salary struct {
	gorm.Model
	EmployeeID uint      `gorm:"index"`
	Employee   Employee  `gorm:"foreignKey:EmployeeID"`
	Salary     float64   `gorm:"type:decimal(10,2);not null"`
	StartDate  time.Time `gorm:"not null"`
	EndDate    *time.Time
}

// Position model
type Position struct {
	gorm.Model
	JobTitle     string       `gorm:"type:varchar(30);not null"`
	MinSalary    float64      `gorm:"type:decimal(10,2)"`
	MaxSalary    float64      `gorm:"type:decimal(10,2)"`
	JobHistories []JobHistory `gorm:"foreignKey:AllJobsID"`
}

// JobHistory model
type JobHistory struct {
	gorm.Model
	EmployeeID uint      `gorm:"index"`
	Employee   Employee  `gorm:"foreignKey:EmployeeID"`
	AllJobsID  uint      `gorm:"index"`
	Position   Position  `gorm:"foreignKey:AllJobsID"`
	StartDate  time.Time `gorm:"not null"`
	EndDate    *time.Time
}

// DepartmentManagerHistory model
type DepartmentManagerHistory struct {
	gorm.Model
	EmployeeID   uint       `gorm:"index"`
	Employee     Employee   `gorm:"foreignKey:EmployeeID"`
	DepartmentID uint       `gorm:"index"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
	StartDate    time.Time  `gorm:"not null"`
	EndDate      *time.Time
}
