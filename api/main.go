package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vishalol/kafka-pubsub/kafka"
	"github.com/vishalol/kafka-pubsub/protomessage"
	"google.golang.org/protobuf/proto"
)

// gin http server instance which listens on port 8080
func main() {
	kafka.InitProducer()
	defer kafka.CloseProducer()

	router := gin.Default()

	// health check endpoint
	router.GET("/health", HealthCheck)

	// api group
	router.POST("/api/message", PublishMessage)

	router.Run(":8080")
}

func PublishMessage(c *gin.Context) {
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert message to protobuf
	protoMessage := protomessage.Message{
		Symbol:    message.Symbol,
		Val:       uint64(message.Val),
		Timestamp: uint64(time.Now().UnixNano()),
	}
	payload, err := proto.Marshal(&protoMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal protobuf message"})
		return
	}

	// Push payload to Kafka topic
	if err := kafka.ProduceMessage(payload, message.Symbol); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to produce message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent to Kafka topic"})
}

// HealthCheck is a simple health check endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "Health check passed")
}

type Message struct {
	Symbol string `json:"symbol"`
	Val    int    `json:"val"`
}
