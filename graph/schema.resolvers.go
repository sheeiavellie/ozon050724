package graph

import (
	"context"
	"fmt"

	"github.com/sheeiavellie/ozon050724/graph/model"
)

func (r *mutationResolver) CreatePost(
	ctx context.Context,
	input model.NewPost,
) (*model.Post, error) {
	post, err := r.storage.CreatePost(ctx, &input)
	if err != nil {
		return nil, fmt.Errorf("error creating post: %w", err)
	}
	return post, nil
}

func (r *mutationResolver) CreateComment(
	ctx context.Context,
	input model.NewComment,
) (*model.Comment, error) {
	comment, err := r.storage.CreateComment(ctx, &input)
	if err != nil {
		return nil, fmt.Errorf("error creating comment: %w", err)
	}
	return comment, nil
}

func (r *mutationResolver) UpdatePost(
	ctx context.Context,
	input model.UpdatePost,
) (*model.Post, error) {
	updatedPost, err := r.storage.UpdatePost(ctx, &input)
	if err != nil {
		return nil, fmt.Errorf("error updating post: %w", err)
	}
	return updatedPost, nil
}

func (r *queryResolver) Post(
	ctx context.Context,
	id int,
) (*model.Post, error) {
	post, err := r.storage.GetPostByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting post by id: %w", err)
	}
	return post, nil
}

func (r *queryResolver) Posts(
	ctx context.Context,
) ([]*model.Post, error) {
	posts, err := r.storage.GetPosts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting posts: %w", err)
	}
	return posts, nil
}

func (r *queryResolver) Comment(
	ctx context.Context,
	id int,
) (*model.Comment, error) {
	comment, err := r.storage.GetCommentByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting comment by id: %w", err)
	}
	return comment, nil
}

func (r *queryResolver) Comments(
	ctx context.Context,
	postID int,
) ([]*model.Comment, error) {
	comments, err := r.storage.GetPostComments(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("error getting post comments: %w", err)
	}
	return comments, nil
}

func (r *resolver) Mutation() MutationResolver { return &mutationResolver{r} }
func (r *resolver) Query() QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *resolver }
type queryResolver struct{ *resolver }
