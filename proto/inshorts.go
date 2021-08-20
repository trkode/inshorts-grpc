package proto

import (
	context "context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
)

type Service struct {
	// UnimplementedArticlesServer
	// serve Serve
	Connect *sql.DB
}
// type Serve interface {
// 	Initdb()
// 	CreateArticle(user *DbUser) (*DbUser, error)
// 	// GetCustomer(id string) (*DbUser, error)
// 	// SearchCustomer(email string, firstName string) (*DbUser, error)
// 	// ListAllCustomer() *[]DbUser
// 	// UpdateCustomer(id string, user *DbUser) (*DbUser, error)
// 	// DeleteCustomer(id *string) error
// }
func (ser *Service) CreateArticle(ctx context.Context, req *CreateArticleRequest) (*Article, error) {
	db := ser.Connect
	var a Article
	a.Title = strings.TrimSpace(req.GetTitle())
	a.Subtitle = strings.TrimSpace(req.GetSubtitle())
	a.Content = strings.TrimSpace(req.GetContent())
	a.CreationTimestamp, _ = ptypes.TimestampProto(time.Now())

	if len(a.Title) == 0 || len(a.Subtitle) == 0 || len(a.Content) == 0 {
		log.Printf("Title/Subtitle/Content Cannot be empty")
		return &a, errors.New("Title/Subtitle/Content Cannot be empty")
	}
	//var id int
	if err := db.QueryRow(
		"INSERT INTO info1(title, subtitle, content, creationtimestamp) VALUES($1, $2, $3, $4) returning id",
		a.Title,
		a.Subtitle,
		a.Content,
		a.CreationTimestamp,
	).Scan(&a.ID); err != nil {
		log.Println("unable to insert article: ", err)
		return &a, errors.New("Unknown")
	}

	//a.ID = int64(id)
	return &a, nil

}
