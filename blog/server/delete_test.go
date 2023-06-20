package main

import (
	"context"
	"errors"
	"testing"

	pb "github.com/MelquesPaiva/grpc-course/blog/proto"
	"github.com/MelquesPaiva/grpc-course/blog/server/mocks"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDeleteBlog(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()

	test := mocks.NewCrudInterface(t)

	test.On("FindOneAndDelete", ctx, bson.M{"_id": "123"}).
		Return(mock.AnythingOfType("*mongo.SingleResult"))

	test.On("Decode", &bson.M{}).
		Return(nil)

	test.On("DeleteBlog", ctx, &pb.BlogId{Id: "123"}).
		Return(&empty.Empty{}, errors.New("testing error"))

	// test.On("FindOneAndDelete", ctx, bson.M{"_id": "123"}).
	// 	Return(mock.AnythingOfType("*mongo.SingleResult"))

	// test.On("Decode", &bson.M{}).
	// 	Return(nil)

	// s := Server{}
	_, err := test.DeleteBlog(ctx, &pb.BlogId{Id: "123"})

	assert.NotNil(t, err)

	test.AssertExpectations(t)

	// assert := assert.New(t)

	// mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// defer mt.Close()

	// mt.Run("success", func(mt *mtest.T) {
	// 	collection = mt.Coll
	// })
}
