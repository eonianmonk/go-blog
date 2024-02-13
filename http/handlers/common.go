package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var (
	PostIdParameter = "id"
)

func getPostIdParam(c *fiber.Ctx) (int, error) {
	idStr := c.Params(PostIdParameter)
	id, err := strconv.Atoi(idStr)
	return id, fmt.Errorf("failed to parse \"%s\" as int: %s", idStr, err.Error())
}
