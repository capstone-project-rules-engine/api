package management

import (
	"brms/endpoints/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App){
	app.Post("/insertRule", InsertRuleSet)
	app.Get("/fetchRules", ListAllRuleSet)
}

func InsertRuleSet(c *fiber.Ctx) error{
	// check if method is not post
	if c.Method() != fiber.MethodPost{
		return fiber.NewError(fiber.StatusMethodNotAllowed, "invalid http method")
	}

	// parse to struct
	var ruleSet models.RuleSet
	if err := c.BodyParser(&ruleSet); err != nil{
		return fiber.NewError(fiber.StatusUnprocessableEntity, "The request entity contains invalid or missing data")
	}

	// validate required fields

	// insert ruleSet
	mongoID, err := InsertOneRule(ruleSet)
	if err != nil{
		if err.Error() == "rule set already exists"{
			return fiber.NewError(fiber.StatusConflict, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	c.Set("Location: ", fmt.Sprintf("%s/%s", c.BaseURL(), mongoID))

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "new rule set inserted",
	})
}

func ListAllRuleSet(c *fiber.Ctx) error{
	if c.Method() != fiber.MethodGet{
		return fiber.NewError(fiber.StatusMethodNotAllowed, "invalid http method")
	}

	// fetch all rule set
	list, err := FetchAllRules()
	if err != nil{
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// check if no rule set from mongo
	if len(list) == 0{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "rule set list empty",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "listing all rule sets",
		"details": list,
	})
}