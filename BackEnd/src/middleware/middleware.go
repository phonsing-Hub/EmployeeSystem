package middleware
import(
	"github.com/gofiber/fiber/v2"
	"github.com/phonsing-Hub/EmployeeSystem/src/utils"

)

func Token(c *fiber.Ctx) error {
	_, err := utils.VerifyToken(c.Cookies("emp_auth"));
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized");
	}
	return c.Next();
}