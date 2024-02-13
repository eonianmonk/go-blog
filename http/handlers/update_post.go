package handlers

import (
	"encoding/json"

	"github.com/eonianmonk/go-blog/assets/pb_generated/blogpb"
	mw "github.com/eonianmonk/go-blog/http/middleware"
	"github.com/eonianmonk/go-blog/http/responses"
	"github.com/gofiber/fiber/v2"
)

func UpdatePost(c *fiber.Ctx) error {
	log := mw.GetLog(c)
	rpc := mw.GetRPC(c)
	var ph *blogpb.PostHeader
	err := json.Unmarshal(c.Body(), &ph)
	if err != nil {
		log.Infof("failed to update post: %s", err.Error())
		return responses.SendErrorResponse(c, err.Error(), fiber.StatusBadRequest)
	}
	_, err = rpc.UpdatePost(c.Context(), ph)
	if err != nil {
		log.Infof("rpc failed to update post: %s", err.Error())
		return responses.SendErrorResponse(c, "internal error", fiber.StatusInternalServerError)
	}
	ph.Body = nil // to avoid extra network usage

	return c.Status(fiber.StatusOK).JSON(ph)
}
