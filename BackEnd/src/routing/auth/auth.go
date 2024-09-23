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
	app.Post("/register", func(c *fiber.Ctx) error {
		return register(c, db)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return login(c, db)
	})
	// app.Post("/logout", func(c *fiber.Ctx) error {
	// 	return logout(c, db)
	// })
}

type Employees struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
	Role     string `json:"role"`
}

func register(c *fiber.Ctx, db *gorm.DB) error {
	e := new(Employees)
	if err := c.BodyParser(e); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	hash, _ := utils.HashPassword(e.Password)
	data := models.AuthUser{
		Email:    e.Email,
		Password: hash,
		Role:     e.Role,
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
	if err := db.Select("id, email, password, role").Where("email = ?", e.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("user not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("error retrieving user")
	}

	if !utils.CheckPasswordHash(e.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid password")
	}

	//set Tonken
	token, err := utils.CreateToken(user.ID, user.Email, user.Role)
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
	err = db.Where("user_id = ?", user.ID).First(&existingToken).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).SendString("error checking for existing token: " + err.Error())
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//INSERT INTO mydb.tokens (user_id, token, expires_at) VALUES(?, ?, ?);
		tokenModel := models.Token{
			UserID:    user.ID,
			Token:     token,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		}
		if err := db.Create(&tokenModel).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("error creating token: " + err.Error())
		}
	} else {
		// UPDATE tokens SET token = <newTokenValue>, expires_at = <newExpiryTime>, updated_at = <currentTime> WHERE id = existingTokenID;
		existingToken.Token = token
		existingToken.ExpiresAt = time.Now().Add(24 * time.Hour)
		if err := db.Save(&existingToken).Error; err != nil {
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
