package handlers

import (
	"github.com/eonianmonk/go-blog/assets/pb_generated/blogpb"
	mw "github.com/eonianmonk/go-blog/http/middleware"
	"github.com/eonianmonk/go-blog/http/responses"
	"github.com/gofiber/fiber/v2"
)

func DeletePost(c *fiber.Ctx) error {
	log := mw.GetLog(c)
	rpc := mw.GetRPC(c)
	id, err := getPostIdParam(c)
	if err != nil {
		log.Infof("invalid id in delete post: %s", err.Error())
		return responses.SendErrorResponse(c, "invalid post id", fiber.StatusBadRequest)
	}
	_, err = rpc.DeletePost(c.Context(), &blogpb.PostId{Id: int32(id)})
	if err != nil {
		log.Infof("failed to delete post: %s", err.Error())
		return responses.SendErrorResponse(c, "failed to delete post", fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusOK)
}
