package middleware

import (
	"github.com/eonianmonk/go-blog/assets/pb_generated/blogpb"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

const (
	blogRPCKey = "blog"
	logKey     = "log"
)

func SetRPC(c *fiber.Ctx, rpc blogpb.BlogClient) {
	c.Locals(blogRPCKey, rpc)
}

func GetRPC(c *fiber.Ctx) blogpb.BlogClient {
	return c.Locals(blogRPCKey).(blogpb.BlogClient)
}

func SetLog(c *fiber.Ctx, log *logrus.Logger) {
	c.Locals(logKey, log)
}
func GetLog(c *fiber.Ctx) *logrus.Logger {
	return c.Locals(logKey).(*logrus.Logger)
}
