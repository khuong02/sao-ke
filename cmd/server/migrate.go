package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Open the CSV file
	file, err := os.Open("chuyen_khoan.csv") // replace with your actual file path
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Initialize CSV reader with flexible options
	reader := csv.NewReader(file)
	reader.LazyQuotes = true // Allows for fields with unescaped quotes
	reader.TrimLeadingSpace = true

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("err:", err)
	}

	// Setup MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://host.docker.internal:27017") // replace with your MongoDB URI
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("sao_ke").Collection("transactions") // replace with your database and collection name

	// Process CSV rows (skipping header)
	for i, row := range records {
		if i == 0 {
			continue // skip header row
		}

		dateParts := strings.Split(row[0], "_")
		if len(dateParts) < 1 {
			log.Fatal("Date parsing error: incorrect format in date_time column")
		}
		dateStr := dateParts[0]

		// Parse the date only
		date, err := time.Parse("02/01/2006", dateStr)
		if err != nil {
			log.Fatal("Date parsing error:", err)
		}

		// Convert trans_no, credit, and debit to appropriate types
		transNo, _ := strconv.Atoi(row[1])
		credit, _ := strconv.ParseFloat(row[2], 64)
		debit, _ := strconv.ParseFloat(row[3], 64)
		detail := row[4]

		// Create a document for MongoDB
		doc := bson.D{
			{Key: "date_time", Value: date},
			{Key: "trans_no", Value: transNo},
			{Key: "credit", Value: credit},
			{Key: "debit", Value: debit},
			{Key: "detail", Value: detail},
		}

		// Insert document into MongoDB
		_, err = collection.InsertOne(context.TODO(), doc)
		if err != nil {
			log.Fatal("MongoDB insertion error:", err)
		}
		fmt.Println("Inserted record:", doc)
	}
}
