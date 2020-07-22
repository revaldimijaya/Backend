package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Go_Backend/graph/generated"
	"Go_Backend/graph/model"
	"context"
	"errors"
	"fmt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	user := model.User{
		ID:         input.ID,
		Name:       input.Name,
		Membership: input.Membership,
		Photo:      input.Photo,
		Subscriber: input.Subscriber,
	}

	_, err := r.DB.Model(&user).Insert()

	if err != nil {
		return nil, errors.New("Insert new user failed")
	}

	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.NewUser) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("user not found!")
	}

	user.Name = input.Name
	user.Membership = input.Membership
	user.Photo = input.Photo

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("Update computer failed")
	}

	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	var user model.User

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("User not found!")
	}

	_, deleteErr := r.DB.Model(&user).Where("id = ?", id).Delete()

	if deleteErr != nil {
		return false, errors.New("Delete user failed")
	}

	return true, nil
}

func (r *mutationResolver) CreateVideo(ctx context.Context, input *model.NewVideo) (*model.Video, error) {
	video := model.Video{
		UserID:      input.UserID,
		URL:         input.URL,
		Watch:       input.Watch,
		Like:        input.Like,
		Dislike:     input.Dislike,
		Restriction: input.Restriction,
		Location:    input.Location,
		Name:        input.Name,
		Thumbnail:   input.Thumbnail,
		Premium:     input.Premium,
		Category:    input.Category,
		Description: input.Description,
		Visibility:  input.Visibility,
		Day:         input.Day,
		Month:       input.Month,
		Year:        input.Year,
	}

	_, err := r.DB.Model(&video).Insert()

	if err != nil {
		return nil, errors.New("Insert new video failed")
	}

	return &video, nil
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, id string, input *model.NewVideo) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("video not found!")
	}

	video.UserID = input.UserID
	video.URL = input.URL
	video.Watch = input.Watch
	video.Like = input.Like
	video.Dislike = input.Dislike
	video.Restriction = input.Restriction
	video.Location = input.Location

	_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("Update video failed")
	}

	return &video, nil
}

func (r *mutationResolver) Watch(ctx context.Context, id int) (bool, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("video not found!")
	}

	video.Watch += 1
	r.DB.Model(&video).Where("id = ?", id).Update()

	return true, nil
}

func (r *mutationResolver) VideoLike(ctx context.Context, id int, userid string, typeArg string) (bool, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("video not found!")
	}

	var like model.LikeVideo

	err_like := r.DB.Model(&like).Where("video_id = ? AND user_id = ?", id, userid).First()

	if err_like != nil {

		insert := model.LikeVideo{
			UserID:  userid,
			VideoID: id,
			Type:    typeArg,
		}
		_, err_insert := r.DB.Model(&insert).Insert()

		if err_insert != nil {
			return false, errors.New("insert video like failed")
		}

		return true, nil
	}

	diff_like := r.DB.Model(&like).Where("video_id = ? AND user_id = ? AND type = ?", id, userid, "like").First()

	if diff_like == nil && typeArg == "like" {
		r.DB.Model(&like).Where("video_id = ? AND user_id = ?", id, userid).Delete()

	} else if diff_like == nil && typeArg == "dislike" {
		like.Type = "dislike"
		r.DB.Model(&like).Where("video_id = ? AND user_id = ?", id, userid).Update()

	} else if diff_like != nil && typeArg == "dislike" {
		r.DB.Model(&like).Where("video_id = ? AND user_id = ?", id, userid).Delete()

	} else if diff_like != nil && typeArg == "like" {
		like.Type = "like"
		r.DB.Model(&like).Where("video_id = ? AND user_id = ?", id, userid).Update()
	}

	return true, nil
}

func (r *mutationResolver) CommentLike(ctx context.Context, id int, userid string, typeArg string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReplyLike(ctx context.Context, id int, userid string, typeArg string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, id string) (bool, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("video not found!")
	}

	_, deleteErr := r.DB.Model(&video).Where("id = ?", id).Delete()

	if deleteErr != nil {
		return false, errors.New("Delete video failed")
	}

	return true, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {
	comment := model.Comment{
		UserID:  input.UserID,
		VideoID: input.VideoID,
		Comment: input.Comment,
		Like:    input.Like,
		Dislike: input.Dislike,
		Day:     input.Day,
		Month:   input.Month,
		Year:    input.Year,
	}

	_, err := r.DB.Model(&comment).Insert()

	if err != nil {
		return nil, errors.New("Insert new comment failed")
	}

	return &comment, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, userid string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var user []*model.User

	err := r.DB.Model(&user).Order("id").Select()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to query users")
	}

	return user, nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	var video []*model.Video

	err := r.DB.Model(&video).Order("id").Select()

	if err != nil {
		return nil, errors.New("Failed to query users")
	}

	return video, nil
}

func (r *queryResolver) Comment(ctx context.Context, videoid int) ([]*model.Comment, error) {
	var comment []*model.Comment

	err := r.DB.Model(&comment).Where("video_id = ?", videoid).Select()

	if err != nil {
		return nil, errors.New("comment not found!")
	}

	return comment, nil
}

func (r *queryResolver) GetUserID(ctx context.Context, userid string) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Where("id = ?", userid).First()

	if err != nil {
		return nil, errors.New("user not found!")
	}

	return &user, nil
}

func (r *queryResolver) GetVideoByUser(ctx context.Context, userid string) ([]*model.Video, error) {
	var video []*model.Video

	err := r.DB.Model(&video).Where("user_id = ?", userid).Select()

	if err != nil {
		return nil, errors.New("video not found!")
	}

	return video, nil
}

func (r *queryResolver) GetVideoID(ctx context.Context, videoid int) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", videoid).First()

	if err != nil {
		return nil, errors.New("video not found!")
	}

	return &video, nil
}

func (r *queryResolver) GetNextVideo(ctx context.Context, videoid int) ([]*model.Video, error) {
	var video []*model.Video

	err := r.DB.Model(&video).Where("id != ?", videoid).Select()

	if err != nil {
		return nil, errors.New("video not found!")
	}

	return video, nil
}

func (r *queryResolver) GetVideoLike(ctx context.Context, videoid int, typeArg string) ([]*model.LikeVideo, error) {
	var like []*model.LikeVideo

	r.DB.Model(&like).Where("video_id = ? AND type = ?", videoid, typeArg).Select()

	return like, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) VideoDislike(ctx context.Context, id int, userid string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) CommentDislike(ctx context.Context, id int, userid string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) ReplyDislike(ctx context.Context, id int, userid string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) Like(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) Dislike(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
