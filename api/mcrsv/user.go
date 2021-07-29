package mcrsv

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

type Server struct {
	DB     *gorm.DB
}

func (s *Server) CreateUser(ctx context.Context, user *ProtoUser) (*ProtoResponse, error) {

	u := User{}
	u.Idx = int(user.Idx)
	u.UserId = user.UserId 
	u.Username = user.Username
	u.Password = user.Password
	err := u.CreateUser(s.DB)

	if err != nil {
		return &ProtoResponse{
			Res: 0, // <<false
		}, err
	}
	return &ProtoResponse{
		Res: 1,
	}, nil
} 