package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/phonsing-Hub/EmployeeSystem/src/models"
	"github.com/phonsing-Hub/EmployeeSystem/src/utils"
	"gorm.io/gorm"
	"time"
)

func SetAuthentication(app fiber.Router, db *gorm.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return auth(c, db)
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		return register(c, db)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return login(c, db)
	})
}

type Employees struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"pass"`
	Role     string `json:"role"`
}

func auth(c *fiber.Ctx, db *gorm.DB) error {
	var user struct {
		ID         uint    `json:"id"`
		Name       string  `json:"firstname"`
		Lastname   string  `json:"lastname"`
		Email      string  `json:"email"`
		Phone      string  `json:"phone"`
		Department string  `json:"departmentname"`
		JobTitle   string  `json:"positions"`
		Salary     float64 `json:"salary"`
	}
	token := c.Cookies("emp_auth")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	deta, err := utils.VerifyToken(c.Cookies("emp_auth"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	id, ok := deta["id"]
	if !ok {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve id")
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

	if err := db.Raw(query, id).Scan(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch employee"})
	}

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}

	return c.Status(200).JSON(user)
}

func register(c *fiber.Ctx, db *gorm.DB) error {
	e := new(Employees)
	if err := c.BodyParser(e); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	hash, _ := utils.HashPassword(e.Password)
	data := models.AuthUser{
		EmployeeID: int(e.ID),
		Email:      e.Email,
		Password:   hash,
		Role:       e.Role,
	}
	result := db.Create(&data)
	if result.Error != nil {
		return c.Status(fiber.StatusConflict).SendString("create user unsuccessful: " + result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).SendString("create user successful!")
}

func login(c *fiber.Ctx, db *gorm.DB) error {
	e := new(Employees)
	if err := c.BodyParser(e); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var user models.AuthUser
	//   SELECT id, email, password, role FROM mydb.auth_users WHERE email = ?
	if err := db.Select("employee_id, email, password, role").Where("email = ?", e.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("user not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("error retrieving user")
	}

	if !utils.CheckPasswordHash(e.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid password")
	}

	//set Tonken
	token, err := utils.CreateToken(uint(user.EmployeeID), user.Email, user.Role)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString("create user unsuccessful: " + err.Error())
	}
	now := time.Now()
	user.LastLogin = &now
	//UPDATE auth_users SET last_login = <time.Now> WHERE id = <user.ID>;
	if err := db.Model(&user).Update("last_login", user.LastLogin).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("error updating last login time: " + err.Error())
	}

	var existingToken models.Token
	err = db.Where("user_id = ?", user.EmployeeID).First(&existingToken).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).SendString("error checking for existing token: " + err.Error())
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//INSERT INTO mydb.tokens (user_id, token, expires_at) VALUES(?, ?, ?);
		tokenModel := models.Token{
			UserID:    user.EmployeeID,
			Token:     token,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		}
		if err := db.Create(&tokenModel).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("error creating token: " + err.Error())
		}
	} else {
		// UPDATE tokens SET token = <newTokenValue>, expires_at = <newExpiryTime>, updated_at = <currentTime> WHERE id = existingTokenID;
		if err := db.Model(&existingToken).
			Where("user_id = ?", user.EmployeeID).
			Updates(models.Token{Token: token, ExpiresAt: time.Now().Add(24 * time.Hour)}).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("error updating token: " + err.Error())
		}
	}
	// set database : CREATE EVENT delete_expired_tokens ON SCHEDULE EVERY 1 HOUR DO DELETE FROM tokens WHERE expires_at < NOW();

	c.Cookie(&fiber.Cookie{
		Name:     "emp_auth",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(fiber.StatusOK).SendString("login successful!")
}

// func logout(c *fiber.Ctx, db *gorm.DB) error {
// 	return c.SendString("logout")
// }
