package rtmsaw

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func Insertmonitor(proker string, status string, about string, karyawan string) (InsertedID interface{}) {
	var rtmdb Monitor
	rtmdb.Proker = proker
	rtmdb.Status = status
	rtmdb.About = about
	rtmdb.Karyawan = karyawan
	return InsertOneDoc("dbmonitor", "rtmdb", rtmdb)
}

func GetDatamonitor(targetselesai string) (data []Monitor) {
	user := MongoConnect("Monitor").Collection("rtmdb")
	filter := bson.M{"targetselesai": targetselesai}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataMonitor :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
