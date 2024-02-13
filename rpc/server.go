package rpc

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eonianmonk/go-blog/assets/pb_generated/blogpb"
	"github.com/eonianmonk/go-blog/rpc/storage"
)

func (s *server) CreatePost(ctx context.Context, ph *blogpb.PostHeader) (*blogpb.PostId, error) {
	if ph.Body == nil {
		return nil, fmt.Errorf("body required fot post")
	}
	id, err := s.q.CreatePost(ctx, storage.CreatePostParams{
		Title:  sql.NullString{String: ph.Title},
		Author: sql.NullString{String: ph.Author},
		Body:   sql.NullString{String: *ph.Body},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %s", err.Error())
	}
	return &blogpb.PostId{Id: id}, nil
}

func (s *server) ReadPost(ctx context.Context, id *blogpb.PostId) (*blogpb.PostHeader, error) {
	p, err := s.q.GetPost(ctx, id.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get post %d: %s", id.Id, err.Error())
	}
	return &blogpb.PostHeader{
		PostId: p.ID,
		Author: p.Author.String,
		Title:  p.Title.String,
		Body:   &p.Body.String,
	}, nil
}

func (s *server) UpdatePost(ctx context.Context, ph *blogpb.PostHeader) (*blogpb.PostId, error) {
	id, err := s.q.Update(ctx, storage.UpdateParams{
		Title:  sql.NullString{String: ph.Title},
		Author: sql.NullString{String: ph.Author},
		Body:   sql.NullString{String: *ph.Body},
		ID:     ph.PostId,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update post: %s", err.Error())
	}
	return &blogpb.PostId{Id: id.ID}, nil
}

func (s *server) DeletePost(ctx context.Context, postid *blogpb.PostId) (*blogpb.PostId, error) {
	id, err := s.q.Delete(ctx, postid.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete value: %s", err.Error())
	}
	return &blogpb.PostId{Id: id}, nil
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
		panic("limited headers request unimplemented")
	}

	return &rh, nil
}
