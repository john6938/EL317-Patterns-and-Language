package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

func setupServer() {
	// Only start server if we are running on heroku
	if isOnHeroku() {
		log.Println("Setting up REST API handlers")
		// Create router for gin library
		router := gin.Default()

		// Sentence API
		router.GET("/api/processSentence/:sentence", processSentenceHandler)

		// Our static files
		router.Static("/js", "./static/js")
		router.Static("/css", "./static/css")

		// Link our index.html to default path
		router.StaticFile("/", "./static/index.html")

		// Log error in case of server crash
		log.Fatal(router.Run(":" + getHerokuPort()))
	}
}

func processSentenceHandler(c *gin.Context) {
	// Get sentence string from api call
	sentence := c.Param("sentence")
	// Parse sentence using WebFromatter
	parsedSentence := processSentence(sentence, WebFormatter{})
	// Use string.Builder to increase performance when adding strings
	response := strings.Builder{}
	for _, s := range parsedSentence {
		response.WriteString(s)
		response.WriteString(" ")
	}

	// Send response with our sentence and HTTP 200 status code (OK)
	c.String(http.StatusOK, response.String())
}

// Check if env variable "PORT" is defined, if yes than we are running on Heroku
func isOnHeroku() bool {
	if getHerokuPort() != "" {
		log.Println("Detected Heroku")
		return true
	}
	return false
}

// Returns port to which the server must bind
func getHerokuPort() string {
	return os.Getenv("PORT")
}
