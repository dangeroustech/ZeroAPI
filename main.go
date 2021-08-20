package main

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

type nodeStatus struct {
	Address string `json:"address"`
	Version string `json:"version"`
	Online  bool   `json:"online"`
}

func run_zt_command(command string) string {
	out, err := exec.Command("/usr/sbin/zerotier-cli", "-j", command).Output()
	log.Default().Println(string(out))
	log.Default().Println(err)
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func main() {
	app := fiber.New()

	// GET /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("These Are Not The Endpoints You're Looking For...")
	})

	// GET /node
	app.Get("/node", func(c *fiber.Ctx) error {
		var response nodeStatus
		json.Unmarshal([]byte(run_zt_command("info")), &response)
		return c.JSON(response)
	})

	log.Fatal(app.Listen(":3000"))
}
