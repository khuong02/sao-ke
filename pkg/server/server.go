package server

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"interview-rest/pkg/model"
	"net/http"
	"strconv"
)

type Server struct {
	router     *gin.Engine
	listenAddr string
	client     *mongo.Client
}

func NewServer(port int, client *mongo.Client) *Server {
	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(config))
	server := &Server{
		router:     router,
		listenAddr: fmt.Sprintf(":%d", port),
		client:     client,
	}

	router.Use(gin.Logger())
	router.Use(server.recoverPanic())

	router.GET("/get-transactions", server.getTransactions)

	return server
}

func (s *Server) Start() error {
	go s.router.Run(s.listenAddr)
	return nil
}

// getTransactions handles GET requests for filtering transactions
func (s *Server) getTransactions(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	// Pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))           // Default to page 1
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10")) // Default to 10 records per page

	// Calculate skip and limit
	skip := (page - 1) * pageSize
	limit := pageSize

	// Build filter using MongoDB query
	filter := BuildFilter(search)

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	// Query MongoDB with the filter
	cursor, err := s.client.
		Database("sao_ke").
		Collection("transactions").
		Find(context.TODO(), filter, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch records"})
		return
	}
	defer cursor.Close(context.TODO())

	// Parse MongoDB results into a slice of Transaction structs
	var transactions []model.Transaction
	for cursor.Next(context.TODO()) {
		var transaction model.Transaction
		err := cursor.Decode(&transaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode record"})
			return
		}
		transactions = append(transactions, transaction)
	}

	totalCount, err := s.client.Database("sao_ke").Collection("transactions").CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count records"})
		return
	}

	// Handle case where no documents are found
	if len(transactions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No records found"})
		return
	}

	resp := model.TransactionResponse{
		Transactions: transactions,
		TotalElement: len(transactions),
		Total:        int(totalCount),
		Page:         page,
		TotalPage:    int(totalCount) / pageSize,
	}

	// Return the filtered records as JSON
	c.JSON(http.StatusOK, resp)
}
