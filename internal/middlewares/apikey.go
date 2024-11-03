package middlewares

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

const apiKeyHeaderKey = "X-API-Key"
const apiKeysFilePath = "secrets/api-keys.txt"

func ApiKeyMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// Get the API key from the request header
		apiKey := ctx.Request.Header.Get(apiKeyHeaderKey)

		// Check if the API key is provided
		if apiKey == "" {
			ctx.JSON(401, gin.H{
				"message": fmt.Sprintf("API key must be provided in the request header with the key %s", apiKeyHeaderKey),
			})

			ctx.Abort()

			return
		}

		// Load the API keys from the file stored on server
		apiKeys, err := loadApiKeys()

		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "failed to read API keys",
			})

			ctx.Abort()

			return
		}

		// Check if the API key is valid
		_, ok := apiKeys[apiKey]
		if !ok {
			ctx.JSON(401, gin.H{
				"message": "invalid API key",
			})

			ctx.Abort()

			return
		}

		ctx.Next()
	}
}

func loadApiKeys() (map[string]struct{}, error) {

	// Open the file
	apiKeysFile, err := os.Open(apiKeysFilePath)

	if err != nil {
		return nil, err
	}

	// Create a scanner
	scanner := bufio.NewScanner(apiKeysFile)

	// Create an empty map to store the API keys
	apiKeys := make(map[string]struct{})

	// Set the split function for the scanner
	scanner.Split(bufio.ScanLines)

	// Read line by line
	for scanner.Scan() {
		// Get the API key
		apiKey := scanner.Text()

		// Collect the API key
		apiKeys[apiKey] = struct{}{}
	}

	return apiKeys, nil
}
