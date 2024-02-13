package handlers

import (
	"github.com/eonianmonk/go-blog/assets/pb_generated/blogpb"
	mw "github.com/eonianmonk/go-blog/http/middleware"
	"github.com/eonianmonk/go-blog/http/responses"
	"github.com/gofiber/fiber/v2"
)

func GetHeaders(c *fiber.Ctx) error {
	log := mw.GetLog(c)
	rpc := mw.GetRPC(c)

	rpcReq := blogpb.RequestHeaders{
		Number: -1, // request all
		Offset: 0,
	}
	headers, err := rpc.GetHeaders(c.Context(), &rpcReq)
	if err != nil {
		log.Infof("failed to get post headers from rpc: %s", err.Error())
		return responses.SendErrorResponse(c, "failed to fetch headers", fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(headers)

}
