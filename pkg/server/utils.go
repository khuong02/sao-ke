package server

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

func BuildFilter(search string) bson.D {
	var filter bson.D

	// If the search term is not empty, construct the $or filter for multiple fields
	if search != "" {
		// Try to convert the search term to numeric values in case it matches `trans_no`, `credit`, or `debit`
		credit, creditErr := strconv.ParseFloat(search, 64)
		debit, debitErr := strconv.ParseFloat(search, 64)
		date, dateErr := time.Parse("2006-01-02", search) // Format: YYYY-MM-DD

		// Create a slice of conditions for the $or filter
		orConditions := bson.A{}

		if creditErr == nil {
			orConditions = append(orConditions, bson.M{"credit": credit})
		}
		if debitErr == nil {
			orConditions = append(orConditions, bson.M{"debit": debit})
		}

		if dateErr == nil {
			orConditions = append(orConditions, bson.M{"date_time": date})
		}

		// Add regex filter for the `detail` field for partial text matching
		orConditions = append(orConditions, bson.M{"detail": bson.M{"$regex": primitive.Regex{Pattern: search, Options: "i"}}})

		// Apply the $or filter with all the conditions
		filter = append(filter, bson.E{Key: "$or", Value: orConditions})
	}

	return filter
}
