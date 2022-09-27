package main

import (
	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	Id       primitive.ObjectID `bson:"_id, omitempty"`
	AuthorId string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func documentToBlog(blog BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       blog.Id.Hex(),
		AuthorId: blog.AuthorId,
		Title:    blog.Title,
		Content:  blog.Content,
	}
}
