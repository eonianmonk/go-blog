package handlers

import (
	"encoding/json"

	"github.com/eonianmonk/go-blog/assets/pb_generated/blogpb"
	mw "github.com/eonianmonk/go-blog/http/middleware"
	"github.com/eonianmonk/go-blog/http/responses"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	log := mw.GetLog(c)
	rpc := mw.GetRPC(c)
	var ph *blogpb.PostHeader
	//err := json.NewDecoder(c.Request().BodyStream())
	err := json.Unmarshal(c.Body(), &ph)
	if err != nil {
		log.Infof("failed to create post: %s", err.Error())
		return responses.SendErrorResponse(c, err.Error(), fiber.StatusBadRequest)
	}
	id, err := rpc.CreatePost(c.Context(), ph)
	if err != nil {
		log.Infof("rpc failed to create post: %s", err.Error())
		return responses.SendErrorResponse(c, "internal error", fiber.StatusInternalServerError)
	}
	ph.Body = nil // to avoid extra network usage
	ph.PostId = id.Id
	return c.Status(fiber.StatusOK).JSON(ph)
}
