package handlers

import (
	"github.com/eonianmonk/go-blog/assets/pb_generated/blogpb"
	mw "github.com/eonianmonk/go-blog/http/middleware"
	"github.com/eonianmonk/go-blog/http/responses"
	"github.com/gofiber/fiber/v2"
)

func GetPost(c *fiber.Ctx) error {
	log := mw.GetLog(c)
	rpc := mw.GetRPC(c)

	id, err := getPostIdParam(c)
	if err != nil {
		log.Infof("invalid in get post: %s", err.Error())
		return responses.SendErrorResponse(c, "invalid post id", fiber.StatusBadRequest)
	}
	post, err := rpc.ReadPost(c.Context(), &blogpb.PostId{Id: int32(id)})
	if err != nil {
		log.Infof("failed to get post %d from rpc: %s", id, err.Error())
		return responses.SendErrorResponse(c, "failed to get post", fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(post)
}
