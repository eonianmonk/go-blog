package rpc

import (
	"context"
	"fmt"

	"github.com/eonianmonk/go-blog/assets/pb_generated/blogpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreatePost(ctx context.Context, ph *blogpb.PostHeader) (*blogpb.PostId, error) {

	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}

func (s *server) ReadPost(ctx context.Context, id *blogpb.PostId) (*blogpb.PostHeader, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPost not implemented")
}

func (s *server) UpdatePost(ctx context.Context, ph *blogpb.PostHeader) (*blogpb.PostId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}

func (s *server) DeletePost(ctx context.Context, id *blogpb.PostId) (*blogpb.PostId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}

func (s *server) GetHeaders(ctx context.Context, ph *blogpb.RequestHeaders) (*blogpb.ReturnHeaders, error) {
	rh := blogpb.ReturnHeaders{
		Headers: make([]*blogpb.PostHeader, 0),
	}
	if ph.Number == -1 {
		h, err := s.q.GetAllHeaders(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get all headers")
		}
		for i := range h {
			rh.Headers = append(rh.Headers, &blogpb.PostHeader{
				Author: h[i].Author.String,
				Title:  h[i].Title.String,
				PostId: h[i].ID,
			})
		}
	} else {
		panic("unimplemented")
	}

	return &rh, nil
}
