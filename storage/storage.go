package storage

import (
	"context"

	"github.com/sheeiavellie/ozon050724/graph/model"
)

type Storage interface {
	GetPostByID(
		ctx context.Context,
		id int,
	) (*model.Post, error)

	GetPosts(ctx context.Context) ([]*model.Post, error)

	GetCommentByID(
		ctx context.Context,
		id int,
	) (*model.Comment, error)

	GetPostComments(
		ctx context.Context,
		postId int,
	) ([]*model.Comment, error)

	CreatePost(
		ctx context.Context,
		input *model.NewPost,
	) (*model.Post, error)

	CreateComment(
		ctx context.Context,
		input *model.NewComment,
	) (*model.Comment, error)

	UpdatePost(
		ctx context.Context,
		input *model.UpdatePost,
	) (*model.Post, error)
}
