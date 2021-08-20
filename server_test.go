package main

import (
	"context"
	"log"
	"reflect"
	"strconv"
	"testing"
	"time"

	pb "github.com/trkode/inshortsapi-grpc/proto"
	"google.golang.org/grpc"
)

func getconn(t *testing.T) (conn *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	t.Cleanup(func() {
		conn.Close()
	})
	return conn
}
func TestOperations(t *testing.T) {
	//serverconn(t)
	conn := getconn(t)
	c := pb.NewArticlesClient(conn)

	t.Run("Create", func(t *testing.T) {
		var all []*pb.Article
		for i := 0; i < 2; i++ {
			req := &pb.CreateArticleRequest{
				Title:    "New title",
				Subtitle: "New Subtitle",
				Content:  "New and interesting content on...",
			}
			res, err := c.CreateArticle(context.Background(), req)
			if err != nil {
				t.Log(err)
			}
			tdiff := time.Since((res.CreationTimestamp).AsTime())
			if tdiff > time.Second {
				t.Error("invalid creation timestamp diff", tdiff)
			}
			var exp pb.Article
			exp.Title = req.Title
			exp.Subtitle = req.Subtitle
			exp.Content = req.Content
			exp.ID = res.ID
			exp.CreationTimestamp = res.CreationTimestamp
			if reflect.DeepEqual(res, exp) {
				t.Errorf("mismatch: got %v want %v", res, exp)
			}
			all = append(all, res)
		}
		t.Run("GetList", func(t *testing.T) {
			req := &pb.ListArticlesRequest{Limit: "", Offset: "0"}
			res, err := c.ListArticles(context.Background(), req)
			if err != nil {
				t.Fatal(err)
			}
			if reflect.DeepEqual(res, all) {
				t.Errorf("mismatch: got %v want %v", res, all)
			}
		})
		t.Run("Search", func(t *testing.T) {
			req := &pb.SearchArticleRequest{Q: all[0].Title}
			res, err := c.SearchArticle(context.Background(), req)
			if err != nil {
				t.Fatal(err)
			}
			if reflect.DeepEqual(res, all[0]) {
				t.Errorf("mismatch: got %v want %v", res, all[0])
			}
		})
		t.Run("Delete", func(t *testing.T) {
			req := &pb.DeleteArticleRequest{Id: strconv.FormatInt(all[0].ID, 10)}
			_, err := c.DeleteArticle(context.Background(), req)
			if err != nil {
				t.Fatal(err)
			}
		})
		t.Run("GetID", func(t *testing.T) {
			req := &pb.GetArticleRequest{Id: strconv.FormatInt(all[0].ID, 10)}
			res, err := c.GetArticle(context.Background(), req)
			if err == nil {
				t.Log("Delete article function fail")
			}
			if res != nil {
				t.Fatal(err)
			}
		})
	})
}
