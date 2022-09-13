package repository

import (
	"fmt"
	"gomongo/config"
	"gomongo/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func Find() {
	db, err := config.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]models.Student, 0)
	for csr.Next(ctx) {
		var row models.Student
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Name  :", result[0].Name)
		fmt.Println("Grade :", result[0].Grade)
	}
}
