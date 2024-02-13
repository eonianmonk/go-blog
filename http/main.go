package http

import (
	"github.com/eonianmonk/go-blog/assets/pb_generated/blogpb"
	"github.com/eonianmonk/go-blog/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(addr string, rpcEp string, log *logrus.Logger) {
	rpc, err := rpcClient(rpcEp)
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		middleware.SetRPC(c, rpc)
		middleware.SetLog(c, log)
		return c.Next()
	})
	setupRoutes(app)
	if err := app.Listen(addr); err != nil {
		panic(err)
	}
}

func rpcClient(ep string) (blogpb.BlogClient, error) {
	conn, err := grpc.Dial(ep,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	c := blogpb.NewBlogClient(conn)
	return c, nil
}
