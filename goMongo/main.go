package main

import(
	"context"
    "fmt"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type person struct{
	Id    string `bson:"title,omitempty"`
	Name  string `bson:"name,omitempty"`
	Age   string `bson:"age,omitempty"`
	City  string `bson:"city,omitempty"`
	Time  int    `bson:"time,omitempty"`
}

func main(){
	fmt.Println("Getting mongo client")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil{
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	err=client.Connect(ctx)
	if err!=nil{
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err=client.Ping(ctx,nil)
	if err!=nil{
		fmt.Println("Error")
	}
	fmt.Println("Connected to MongoDB")

	databases,_:=client.ListDatabaseNames(ctx,bson.M{})
	fmt.Println(databases)

    db:=client.Database("person")
    col:=db.Collection("personDetails")
    fmt.Println(col)

    res,err:=col.InsertOne(ctx,bson.D{
    	{"title","100"},
    	{"name","Jackson"},
    	{"age","21"},
    	{"city","Kansas"},
    	{"time","100"},
    })
    if err!=nil{
    	log.Fatal(err)
    }
    fmt.Println(res)
}