package db

import (
	"context"
	"time"

	"github.com/letanthang/mongo/sequence"
	"github.com/letanthang/my_framework/db/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertStudent(req types.StudentAddReq) (interface{}, error) {
	newID, _ := sequence.GetNextID(Client.Database(dbName).Collection("counter"),
		"student_id_seq")
	student := types.Student{
		ID:        newID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		ClassName: req.ClassName,
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := Client.Database("go3008").Collection("student").InsertOne(ctx, student)
	id := res.InsertedID
	return id, err
}

func DeleteStudent(id int) (*mongo.DeleteResult, error) {
	res, err := Client.Database("go3008").Collection("student").DeleteOne(context.TODO(), bson.D{{"id", id}})
	return res, err
}
