package handlers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// Base URL of the OpenAI API.
const openaiApiBaseUrl string = "https://api.openai.com"

func ProxyHandler(ctx *gin.Context) {
	// Get the request path (excluding the hostname)
	requestPath := ctx.Request.URL.Path

	// Get the OpenAI API endpoint
	endpoint, err := url.JoinPath(openaiApiBaseUrl, requestPath)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "failed to infer the OpenAI API endpoint",
		})
		return
	}

	// Get the request method
	requestMethod := ctx.Request.Method

	// Get the body
	requestBody := ctx.Request.Body

	// Get the header
	requestHeader := ctx.Request.Header

	// Create the request
	req, err := http.NewRequest(requestMethod, endpoint, requestBody)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "failed to create the request",
		})
		return
	}

	// Set the header
	req.Header.Add("Authorization", requestHeader.Get("Authorization"))
	req.Header.Add("Content-Type", requestHeader.Get("Content-Type"))
	req.Header.Add("User-Agent", requestHeader.Get("User-Agent"))
	req.Header.Add("Accept", requestHeader.Get("Accept"))

	// Create a client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "failed to send the request",
		})
		return
	}

	// Send the response data back to user
	ctx.DataFromReader(resp.StatusCode, resp.ContentLength, "application/json", resp.Body, convertHttpHeaderToStringMap(resp.Header))

	// Close the body reader
	resp.Body.Close()
}

func convertHttpHeaderToStringMap(header http.Header) map[string]string {

	// Create an empty map
	headers := make(map[string]string)

	for key, values := range header {

		// Join the values
		// Each value is separated by a comma
		value := strings.Join(values, ", ")

		// Set the key-value pair
		headers[key] = value
	}

	return headers
}
