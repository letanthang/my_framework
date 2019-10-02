package db

import (
	"context"
	"encoding/json"

	"github.com/labstack/gommon/log"
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

	res, err := Client.Database("go3008").Collection("student").InsertOne(context.TODO(), student)
	id := res.InsertedID
	return id, err
}

func DeleteStudent(id int) (*mongo.DeleteResult, error) {
	filter := bson.D{{"id", id}}
	res, err := Client.Database("go3008").Collection("student").DeleteOne(context.TODO(), filter)
	return res, err
}

func UpdateStudent(req types.StudentUpdateReq) (*mongo.UpdateResult, error) {
	filter := bson.M{"id": req.ID}

	var data bson.M
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	update := bson.M{"$set": data}
	res, err := Client.Database(dbName).Collection("student").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	go logAction("update", "student", data)
	return res, err
}

func logAction(action, collection string, data bson.M) (*mongo.InsertOneResult, error) {
	doc := bson.M{"action": action, "collection": collection, "data": data}
	res, err := Client.Database(dbName).Collection("log").InsertOne(context.TODO(), doc)
	return res, err
}
