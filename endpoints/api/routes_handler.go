package management

import (
	"brms/endpoints/models"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/insertRuleTemplate", insertRuleTemplate)
	app.Patch("/insertRuletoRuleSet", insertRulestoRuleSet)
	app.Put("/updateRuleSet", updateRuleSet) // tambahin 2 routes ini lagi
	// app.Post("/execInput")
	app.Get("/fetchRules", ListAllRuleSet)
}

func insertRuleTemplate(c *fiber.Ctx) error {
	// check if method is not post
	if c.Method() != fiber.MethodPost {
		return fiber.NewError(fiber.StatusMethodNotAllowed, "invalid http method")
	}

	// parse to struct
	var ruleSet models.RuleSet
	if err := c.BodyParser(&ruleSet); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "The request entity contains invalid or missing data")
	}

	// validate required fields
	validator := validator.New()
	if err := validator.Struct(ruleSet); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "empty fields")
	}

	// insert ruleSet
	mongoID, err := InsertOneRule(ruleSet)
	if err != nil {
		if err.Error() == "rule set already exists" { // conflict check 
			return fiber.NewError(fiber.StatusConflict, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	c.Set("Location", fmt.Sprintf("%s/%s", c.BaseURL(), mongoID)) // set header location to satisfy 201 code

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "new rule set inserted",
	})
}

func insertRulestoRuleSet(c *fiber.Ctx) error {
	// check method
	if c.Method() != fiber.MethodPatch{
		return fiber.NewError(fiber.StatusMethodNotAllowed, "invalid http method")
	}

	// get query params
	ruleSetName := c.Query("ruleSetName")
	if ruleSetName == ""{
		return fiber.NewError(fiber.StatusBadRequest, "query parameter required")
	}

	// parse from request body to struct
	var newRules []models.Rule
	if err := c.BodyParser(&newRules); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "The request entity contains invalid or missing data")
	}

	// insert new rules
	if err := InsertRulestoRuleSet(ruleSetName, newRules); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("%d new rules has been inserted to %s", len(newRules), ruleSetName),
	})
}

func ListAllRuleSet(c *fiber.Ctx) error {
	// check method
	if c.Method() != fiber.MethodGet {
		return fiber.NewError(fiber.StatusMethodNotAllowed, "invalid http method")
	}

	// fetch all rule set
	list, err := FetchAllRules()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// check if no rule set from mongo
	if len(list) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "rule set list empty",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "listing all rule sets",
		"details": list,
	})
}

func updateRuleSet(c *fiber.Ctx) error {
	// Check if method is not PUT
	if c.Method() != fiber.MethodPut {
		return fiber.NewError(fiber.StatusMethodNotAllowed, "invalid http method")
	}

	// Get rule set name from query parameter
	ruleSetName := c.Query("ruleSetName")
	if ruleSetName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "query parameter 'ruleSetName' is required")
	}

	// Parse rule set data from request body
	var updatedRuleSet models.RuleSet
	if err := c.BodyParser(&updatedRuleSet); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "The request entity contains invalid or missing data")
	}

	// Validate the updated rule set
	validator := validator.New()
	if err := validator.Struct(updatedRuleSet); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid fields in updated rule set")
	}

	// Update the rule set in the database
	if err := UpdateRuleSet(ruleSetName, updatedRuleSet); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Rule set '%s' has been updated", ruleSetName),
	})
}
