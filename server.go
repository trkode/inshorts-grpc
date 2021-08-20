package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/lib/pq"
	pb "github.com/trkode/inshortsapi-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	port = ":50056"
)

type Server struct {
	Connect *sql.DB
	pb.UnimplementedArticlesServer
}


func (ser *Server) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.Article, error) {
	var a pb.Article
	a.Title = strings.TrimSpace(req.GetTitle())
	a.Subtitle = strings.TrimSpace(req.GetSubtitle())
	a.Content = strings.TrimSpace(req.GetContent())
	a.CreationTimestamp = timestamppb.New(time.Now())

	if len(a.Title) == 0 || len(a.Subtitle) == 0 || len(a.Content) == 0 {
		log.Printf("Title/Subtitle/Content Cannot be empty")
		return nil, errors.New("Title/Subtitle/Content Cannot be empty")
	}
	//var id int
	if err := ser.Connect.QueryRow(
		"INSERT INTO articles(title, subtitle, content, creationtimestamp) VALUES($1, $2, $3, $4) returning id",
		a.Title,
		a.Subtitle,
		a.Content,
		time.Now(),
	).Scan(&a.ID); err != nil {
		log.Println("unable to insert article: ", err)
		return nil, errors.New("Unknown")
	}
	fmt.Println("Entry Inserted Successfully!")
	return &a, nil

}

const selBase = "SELECT id, title, subtitle, content, creationtimestamp from articles"

func (ser *Server) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (w *pb.ListArticlesResponse, err error) {
	// query := r.URL.Query()
	// // pagination list using limit and offset query parameters
	// limit := query.Get("limit")
	// offset := query.Get("offset")
	limit := req.Limit
	offset := req.Offset
	args := []interface{}{}
	sql := selBase + " ORDER BY creationtimestamp DESC "
	if limit != "" {
		sql += " LIMIT $1"
		args = append(args, limit)
	}

	if offset != "" {
		sql += " OFFSET $" + strconv.Itoa(len(args)+1)
		args = append(args, offset)
	}

	w, err = ser.listArticles(w, sql, args...)
	if err != nil {
		return w, err
	}
	fmt.Println("Sent Articles list Successfully!")
	return w, nil
}

func (ser *Server) listArticles(w *pb.ListArticlesResponse, q string, args ...interface{}) (out *pb.ListArticlesResponse, err error) {
	rows, err := ser.Connect.Query(q, args...)
	if err != nil {
		log.Println("unable to query articles: ", err)
		return nil, err
	}
	defer rows.Close()
	var all []*pb.Article
	for rows.Next() {
		var article pb.Article
		if err := ScanArticle(rows, &article); err != nil {
			log.Println("unable to scan articles: ", err)
			return nil, err
		}
		all = append(all, &article)

	}
	out = &pb.ListArticlesResponse{Articleslist: all}
	return out, nil
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

func ScanArticle(row Scanner, a *pb.Article) error {
	var temp time.Time
	if err := row.Scan(&a.ID, &a.Title, &a.Subtitle, &a.Content, &temp); err != nil {
		return err
	}
	a.CreationTimestamp = timestamppb.New(temp)
	return nil
}

func (ser *Server) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (w *pb.Article, err error) {
	i := req.Id
	id, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		log.Printf("malformed id, should be int64:%v", err)
		return
	}
	var a pb.Article
	const sql = selBase + ` WHERE id = $1`
	if err := ScanArticle(ser.Connect.QueryRow(sql, id), &a); err != nil {
		log.Println("unable to get article from db: ", err)
		return &a, err
	}
	fmt.Println("Requested Article sent Successfully!")
	return &a, nil
}
func (ser *Server) SearchArticle(ctx context.Context, req *pb.SearchArticleRequest) (w *pb.ListArticlesResponse, err error) {
	q := strings.TrimSpace(req.Q)
	w, err = ser.listArticles(w, selBase+` WHERE title like $1 OR subtitle like $1 OR content like $1`, "%"+q+"%")
	if err != nil {
		log.Printf("Unable to search articles:%v", err)
		return w, err
	}
	fmt.Println("Search query article sent Successfully!")
	return w, nil
}
func (ser *Server) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*empty.Empty, error) {
	i := req.Id
	id, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		log.Printf("malformed id, should be int64:%v", err)
		return &empty.Empty{}, err
	}

	if _, err := ser.Connect.Exec("DELETE from articles WHERE id = $1", id); err != nil {
		log.Printf("unable to delete article:%v", err)
		return &empty.Empty{}, err
	}
	fmt.Printf("Deleted article with id %v", id)
	return &empty.Empty{}, nil
}

//OpenDB connects to the psql database
func OpenDB(user, password, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to db:%w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Unable to ping db:%w", err)
	}
	return db, nil
}
func Initdb(s *sql.DB) {
	const createTable = "create table if not exists articles(" +
		"id serial NOT NULL PRIMARY KEY," +
		"title text NOT NULL," +
		"subtitle text NOT NULL," +
		"content text NOT NULL," +
		"creationtimestamp timestamp NOT NULL" +
		");"

	_, err := s.Exec(createTable)
	if err != nil {
		panic(err)
	}
}

var (
	U = flag.String("dbuser", "postgres", "Enter Username")
	P = flag.String("dbpassword", "1357", "Enter Password")
	N = flag.String("dbname", "postgres", "Enter Database Name")
)

func Execute() (error, Server) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("Failed to listen at port %v:%v", port, err)
	}
	grpcServer := grpc.NewServer()

	flag.Parse()
	db, err := OpenDB(*U, *P, *N)
	if err != nil {
		log.Printf("Error in connecting to psql database:%v", err)
	}
	defer db.Close()
	Initdb(db)
	//Connect = db
	s := Server{Connect: db}
	pb.RegisterArticlesServer(grpcServer, &s)
	return grpcServer.Serve(lis), s
}
func main() {
	fmt.Println("Welcome to the server")
	if err, _ := Execute(); err != nil {
		log.Printf("Failed to Serve:%v", err)
	}
}
