package storage

import (
	"context"

	"github.com/sheeiavellie/ozon050724/graph/model"
)

type PostgresStorage struct {
}

func NewPostgresStorage() (*PostgresStorage, error) {
}

func ClosePostgresStorageConnection() error {
}

func (ps *PostgresStorage) GetPostByID(
	ctx context.Context,
	id int,
) (*model.Post, error) {
}

func (ps *PostgresStorage) GetPosts(
	ctx context.Context,
) ([]*model.Post, error) {
}

func (ps *PostgresStorage) GetCommentByID(
	ctx context.Context,
	id int,
) (*model.Comment, error) {
}

func (ps *PostgresStorage) GetPostComments(
	ctx context.Context,
	postId int,
) ([]*model.Comment, error) {
}

func (ps *PostgresStorage) CreatePost(
	ctx context.Context,
	input *model.NewPost,
) (*model.Post, error) {
}

func (ps *PostgresStorage) CreateComment(
	ctx context.Context,
	input *model.NewComment,
) (*model.Comment, error) {
}

func (ps *PostgresStorage) UpdatePost(
	ctx context.Context,
	input *model.UpdatePost,
) (*model.Post, error) {
}
