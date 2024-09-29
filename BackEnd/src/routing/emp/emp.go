package emp

import (
	"github.com/gofiber/fiber/v2"
	//"github.com/phonsing-Hub/EmployeeSystem/src/models"
	"gorm.io/gorm"
	//"time"
)

func EmployeesRouting(app fiber.Router, db *gorm.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return getAllEmployeeDetails(c, db)
	})
	app.Get("/:id", func(c *fiber.Ctx) error {
		return getAllEmployeeById(c,db)
	})
	app.Post("/new", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}

func getAllEmployeeDetails(c *fiber.Ctx, db *gorm.DB) error {
	var employeeDetails []struct {
		ID         uint    `json:"id"`
		Name       string  `json:"firstname"`
		Lastname   string  `json:"lastname"`
		Email      string  `json:"email"`
		Phone      string  `json:"phone"`
		Department string  `json:"department"`
		JobTitle   string  `json:"positions"`
		Salary     float64 `json:"salary"`
	}

	query := `
	SELECT
	    e.employee_id AS id,
	    e.first_name AS name,
	    e.last_name AS lastname,
	    e.email,
	    e.phone_number AS phone,
	    d.department_name AS department,
	    j.job_title AS job_title,
	    e.salary
	FROM
	    employees e
	LEFT JOIN
	    departments d ON e.department_id = d.department_id
	LEFT JOIN
	    jobs j ON e.job_id = j.job_id;`

	if err := db.Raw(query).Scan(&employeeDetails).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch employee details"})
	}

	return c.Status(200).JSON(employeeDetails)
}

func getAllEmployeeById(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	var employeeById struct {
		ID         uint    `json:"id"`
		Name       string  `json:"firstname"`
		Lastname   string  `json:"lastname"`
		Email      string  `json:"email"`
		Phone      string  `json:"phone"`
		Department string  `json:"departmentname"`
		JobTitle   string  `json:"positions"`
		Salary     float64 `json:"salary"`
	}

	query := `
	SELECT
	    e.employee_id AS id,
	    e.first_name AS name,
	    e.last_name AS lastname,
	    e.email,
	    e.phone_number AS phone,
	    d.department_name AS department,
	    j.job_title AS job_title,
	    e.salary
	FROM
	    employees e
	LEFT JOIN
	    departments d ON e.department_id = d.department_id
	LEFT JOIN
	    jobs j ON e.job_id = j.job_id
		WHERE e.employee_id = ?;`

	if err := db.Raw(query, id).Scan(&employeeById).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch employee"})
	}

	if employeeById.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}

	return c.Status(200).JSON(employeeById)
}
