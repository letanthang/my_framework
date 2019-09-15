package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/letanthang/my_framework/db/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOneStudent() (*types.Student, error) {
	var student types.Student
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := Client.Database(dbName).Collection("student").FindOne(ctx, struct{}{}).Decode(&student)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func GetAllStudent() (*[]types.Student, error) {
	var students []types.Student
	findOptions := options.Find()
	findOptions.SetLimit(30)
	cur, err := Client.Database(dbName).Collection("student").Find(context.TODO(), struct{}{}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var student types.Student
		err = cur.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}

func GetStudent(req types.StudentReq) (*[]types.Student, error) {
	var students []types.Student
	findOptions := options.Find()
	findOptions.SetLimit(30)
	var filter map[string]interface{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &filter)
	cur, err := Client.Database(dbName).Collection("student").Find(context.TODO(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var student types.Student
		err = cur.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}

func SearchStudent(req types.StudentSearchReq) (*[]types.Student, error) {
	var students []types.Student
	findOptions := options.Find()
	findOptions.SetLimit(30)
	filter := bson.M{}
	if req.ID != 0 {
		filter["id"] = req.ID
	}

	if req.Email != "" {
		filter["email"] = primitive.Regex{Pattern: req.Email, Options: "i"}
	}
	if req.ClassName != "" {
		filter["class_name"] = primitive.Regex{Pattern: req.ClassName, Options: "i"}
	}

	if req.Name != "" {
		filter["$or"] = bson.A{
			bson.M{"first_name": primitive.Regex{Pattern: req.Name, Options: "i"}},
			bson.M{"last_name": primitive.Regex{Pattern: req.Name, Options: "i"}},
		}
	}

	cur, err := Client.Database(dbName).Collection("student").Find(context.TODO(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var student types.Student
		err = cur.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}

func GroupStudentByLastName() (*[]map[string]interface{}, error) {
	var students []map[string]interface{}

	pipeline := bson.A{
		bson.D{{"$match", bson.D{{"class_name", "golang3008"}}}},
		bson.D{{"$group", bson.D{
			{"_id", "$last_name"},
			{"last_name", bson.D{{"$first", "$last_name"}}},
			{"class_name", bson.D{{"$first", "$class_name"}}},
			{"first_name", bson.D{{"$push", "$first_name"}}},
			{"id", bson.D{{"$push", "$id"}}},
		}}},
	}

	cur, err := Client.Database(dbName).Collection("student").Aggregate(context.TODO(), pipeline)

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var student map[string]interface{}
		err = cur.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}
