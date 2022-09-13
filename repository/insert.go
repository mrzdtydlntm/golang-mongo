package repository

import (
	"context"
	"fmt"
	"gomongo/config"
	"gomongo/models"
	"log"
)

var ctx = context.Background()

func Insert() {
	db, err := config.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, models.Student{"Wick", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, models.Student{"Ethan", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}
