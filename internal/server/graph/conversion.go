package graph

import (
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/model"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
)

func pbUserToGraphQlUser(u *pb.User) *model.User {
	return &model.User{
		UUID:      u.Uuid,
		Login:     u.Login,
		Email:     u.Email,
		Role:      int(u.Role),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func arrPbUserToArrGraphQlUser(arr []*pb.User) []*model.User {
	newArr := []*model.User{}
	for _, u := range arr {
		newArr = append(newArr, pbUserToGraphQlUser(u))
	}

	return newArr
}
