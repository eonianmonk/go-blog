package rpc

import (
	"database/sql"
	"net"

	"github.com/eonianmonk/go-blog/assets/pb_generated/blogpb"
	"github.com/eonianmonk/go-blog/rpc/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	blogpb.UnimplementedBlogServer
	q   *storage.Queries
	log *logrus.Logger
}

var (
	dbtype = "postgres"
)

// ep - listening endpoint
// connStr - database connection string
func NewServer(ep string, connStr string, log *logrus.Logger) {
	l, err := net.Listen("tcp", ep)
	if err != nil {
		panic(err)
	}
	rpcs := grpc.NewServer()
	db := loadDb(connStr)
	q := storage.New(db)
	s := server{
		log: log,
		q:   q,
	}
	blogpb.RegisterBlogServer(rpcs, s)
	if err := rpcs.Serve(l); err != nil {
		panic(err)
	}
}

func loadDb(connStr string) *sql.DB {
	db, err := sql.Open(dbtype, connStr)
	if err != nil {
		panic(err)
	}
	return db
}
