package models

import "time"

type Region struct {
	ID         int       `gorm:"column:region_id;primaryKey;autoIncrement" json:"region_id"`
	RegionName string    `gorm:"column:region_name" json:"region_name"`
	Countries  []Country `gorm:"foreignKey:RegionID" json:"countries"`
}

type Country struct {
	ID         string     `gorm:"column:country_id;primaryKey;size:2" json:"country_id"`
	CountryName string    `gorm:"column:country_name" json:"country_name"`
	RegionID   int        `gorm:"column:region_id" json:"region_id"`
	Region     Region     `gorm:"foreignKey:RegionID" json:"region"`
	Locations  []Location `gorm:"foreignKey:CountryID" json:"locations"`
}

type Location struct {
	ID            int       `gorm:"column:location_id;primaryKey;autoIncrement" json:"location_id"`
	StreetAddress string    `gorm:"column:street_address" json:"street_address"`
	PostalCode    string    `gorm:"column:postal_code" json:"postal_code"`
	City          string    `gorm:"column:city;not null" json:"city"`
	StateProvince string    `gorm:"column:state_province" json:"state_province"`
	CountryID     string    `gorm:"column:country_id;size:2;not null" json:"country_id"`
	Country       Country   `gorm:"foreignKey:CountryID" json:"country"`
	Departments   []Department `gorm:"foreignKey:LocationID" json:"departments"`
}

type Job struct {
	ID        int     `gorm:"column:job_id;primaryKey;autoIncrement" json:"job_id"`
	JobTitle  string  `gorm:"column:job_title;not null" json:"job_title"`
	MinSalary float64 `gorm:"column:min_salary" json:"min_salary"`
	MaxSalary float64 `gorm:"column:max_salary" json:"max_salary"`
	Employees []Employee `gorm:"foreignKey:JobID" json:"employees"`
}

type Department struct {
	ID            int        `gorm:"column:department_id;primaryKey;autoIncrement" json:"department_id"`
	DepartmentName string    `gorm:"column:department_name;not null" json:"department_name"`
	LocationID    int        `gorm:"column:location_id" json:"location_id"`
	Location      Location   `gorm:"foreignKey:LocationID" json:"location"`
	Employees     []Employee `gorm:"foreignKey:DepartmentID" json:"employees"`
}

type Employee struct {
	ID          int        `gorm:"column:employee_id;primaryKey;autoIncrement" json:"employee_id"`
	FirstName   string     `gorm:"column:first_name" json:"first_name"`
	LastName    string     `gorm:"column:last_name;not null" json:"last_name"`
	Email       string     `gorm:"column:email;not null" json:"email"`
	PhoneNumber string     `gorm:"column:phone_number" json:"phone_number"`
	HireDate    time.Time  `gorm:"column:hire_date;not null" json:"hire_date"`
	JobID       int        `gorm:"column:job_id;not null" json:"job_id"`
	Job         Job        `gorm:"foreignKey:JobID" json:"job"`
	Salary      float64    `gorm:"column:salary;not null" json:"salary"`
	ManagerID   *int       `gorm:"column:manager_id" json:"manager_id"`
	Manager     *Employee  `gorm:"foreignKey:ManagerID" json:"manager"`
	DepartmentID *int      `gorm:"column:department_id" json:"department_id"`
	Department  *Department `gorm:"foreignKey:DepartmentID" json:"department"`
	Subordinates []Employee `gorm:"foreignKey:ManagerID" json:"subordinates"`
	Dependents   []Dependent `gorm:"foreignKey:EmployeeID" json:"dependents"`
}

type Dependent struct {
	ID          int     `gorm:"column:dependent_id;primaryKey;autoIncrement" json:"dependent_id"`
	FirstName   string  `gorm:"column:first_name;not null" json:"first_name"`
	LastName    string  `gorm:"column:last_name;not null" json:"last_name"`
	Relationship string `gorm:"column:relationship;not null" json:"relationship"`
	EmployeeID  int     `gorm:"column:employee_id;not null" json:"employee_id"`
	Employee    Employee `gorm:"foreignKey:EmployeeID" json:"employee"`
}

//AuthUser model for authentication
type AuthUser struct {
	EmployeeID       int        `gorm:"primaryKey;" json:"employee_id"`
	Email            string     `gorm:"type:varchar(100);unique;not null"`
	Password         string     `gorm:"type:varchar(255);not null"` // Password should be hashed
	LastLogin        *time.Time // Track the last login time
	Role             string     `gorm:"type:varchar(20);default:admin"` // Default role is admin
	ResetToken       string     `gorm:"type:varchar(255)"`              // For password reset functionality
	ResetTokenExpiry *time.Time // For tracking token expiration
	Employee         Employee   `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"employee"`
}

// Token model for managing JWT tokens
type Token struct {
	UserID    int       `gorm:"not null"`
	Token     string    `gorm:"type:text;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	User      AuthUser  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
