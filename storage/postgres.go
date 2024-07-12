package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sheeiavellie/ozon050724/graph/model"
)

const postgresPingTimeout = 3 * time.Second

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(
	ctx context.Context,
	connStr string,
) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening postgres: %w", err)
	}

	ctxPing, cancel := context.WithTimeout(ctx, postgresPingTimeout)
	defer cancel()

	err = db.PingContext(ctxPing)
	if err != nil {
		return nil, fmt.Errorf("error pinging postgres: %w", err)
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (ps *PostgresStorage) Close() error {
	err := ps.db.Close()
	if err != nil {
		return fmt.Errorf("error closing postgres: %w", err)
	}
	return nil
}

func (ps *PostgresStorage) Init(
	ctx context.Context,
) error {
	return nil
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
