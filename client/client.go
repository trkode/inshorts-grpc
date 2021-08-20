package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/trkode/inshortsapi-grpc/proto"
	"google.golang.org/grpc"
)

var options = []string{
	"Create Article",
	"Get Article List",
	"Get Single Article",
	"Search an Article",
	"Delete an Article",
	"Quit",
}

func showOptions() (optionSelected int) {
	for k, v := range options {
		fmt.Printf("%v ==> %s\n", k+1, v)
	}
	fmt.Println("Please Enter Option number:  ")
	optionSelected = 0
	fmt.Scanf("%d", &optionSelected)
	return
}
func main() {
	fmt.Println("Hello from Client")
	conn, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		log.Printf("Could not Connect %v", err)
	}
	defer conn.Close()
	c := pb.NewArticlesClient(conn)
	switch showOptions() {
	case 1:
		var val []string = []string{"Title", "Subtitle", "Content"}
		r := make([]string, 3)
		for v := range val {
			fmt.Printf("Enter %s: \n", val[v])
			fmt.Scan(&r[v])
		}
		req := &pb.CreateArticleRequest{
			Title:    r[0],
			Subtitle: r[1],
			Content:  r[2],
		}
		res, err := c.CreateArticle(context.Background(), req)
		Check(err)
		fmt.Println(res)

	case 2:
		req := &pb.ListArticlesRequest{Limit: "", Offset: "0"}
		res, err := c.ListArticles(context.Background(), req)
		Check(err)
		fmt.Println(res)
	case 3:
		var id string
		fmt.Print("Enter id: \n")
		fmt.Scan(&id)

		req := &pb.GetArticleRequest{Id: id}
		res, err := c.GetArticle(context.Background(), req)
		Check(err)
		fmt.Println(res)

	case 4:
		var q string
		fmt.Print("Enter query: \n")
		fmt.Scan(&q)
		req := &pb.SearchArticleRequest{Q: q}
		res, err := c.SearchArticle(context.Background(), req)
		Check(err)
		fmt.Println(res)
	case 5:
		var id string
		fmt.Print("Enter id: \n")
		fmt.Scan(&id)
		req := &pb.DeleteArticleRequest{Id: id}
		res, err := c.DeleteArticle(context.Background(), req)
		Check(err)
		fmt.Println(res)
	case 6:
		return
	default:
		fmt.Println("PLEASE ENTER VALID OPTION")
	}

}

//Check to check on err
func Check(err error) {
	if err != nil {
		log.Printf("%v", err)
		return
	}
}
