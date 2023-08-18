package main

import (
	"brms/config"
	"brms/endpoints/management"
	"io"
	"os"

	"brms/pkg/file"
	"brms/pkg/middlewares"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func setApp(file *os.File) *fiber.App { // setting up middlewares
	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${time}] [${severity}] ${path} ${method} (${ip}) ${status} ${latency} - ${message}\n",
		CustomTags: map[string]logger.LogFunc{
			"message": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				if bodyBytes := c.Response().Body(); bodyBytes != nil {
					var bodyData map[string]interface{}

					if err := json.Unmarshal(bodyBytes, &bodyData); err == nil {
						msgValue, _ := bodyData["message"].(string)
						return output.WriteString(msgValue)
					}
				}
				return 0, nil
			},
			"severity": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				status := c.Response().StatusCode()

				if status == fiber.StatusInternalServerError {
					return output.WriteString("WARNING")
				}
				return output.WriteString("INFO")
			},
		},
		Output: io.MultiWriter(os.Stdout, file),
	}))

	app.Use(middlewares.UndefinedRoutesMiddleware())

	app.Use(middlewares.ErrorMiddleware())

	return app
}

func main() {
	file, err := file.OpenLogFile()
	if err != nil{
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer file.Close()

	app := setApp(file)

	// register rule management routes
	management.Routes(app)

	if err := app.Listen(fmt.Sprintf(":%s", config.GetConfig().Port)); err != nil {
		log.New(file, "ERROR: ", log.Ldate|log.Ltime).Println("Application failed to start running: ", err)
		os.Exit(1)
	}
}
