package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"temporal_wf/app"
	"temporal_wf/signal"
	"temporal_wf/workflow"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	//context := *gin.Context,

	log.Printf("Starting Go Gin Server on Post 9090")

	tc, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	fmt.Printf("Creating Temporal client %d", tc)

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/initiateWF", initiateWF(tc))
	router.GET("/notify", notify(tc))
	router.GET("/getStatus", getStatus(tc))
	router.Run("localhost:9090")

}

func getStatus(tc client.Client) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		c.IndentedJSON(http.StatusOK, gin.H{"message": "active"})
	}
	return gin.HandlerFunc(fn)

}

func initiateWF(tc client.Client) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		var init app.Init
		if err := c.BindJSON(&init); err != nil {
			return
		}
		log.Printf("Session init %s.\n\n", init)

		id := (uuid.New()).String()

		options := client.StartWorkflowOptions{
			ID:        id,
			TaskQueue: app.SignupUserQueue,
		}

		we, err := tc.ExecuteWorkflow(context.Background(), options, workflow.SignUpWorkflow, init)
		if err != nil {
			log.Fatalln("Unable to start the Workflow:", err)
		}

		log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

		c.IndentedJSON(http.StatusCreated, id)

	}
	return gin.HandlerFunc(fn)
}

func notify(tc client.Client) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		workflowId := c.Query("workflowId")
		log.Println(workflowId)
		err := signal.SendNotifySignal(workflowId, true)
		if err != nil {
			log.Fatalln("Unable to notify.", err)
		}

		//response, err := tc.QueryWorkflow(context.Background(), workflowId, runID, queryType)
		/*if err != nil {
			log.Fatalln("Unable to start the Workflow:", err)
		} */

		//Un-comment for pricing
		c.Writer.WriteHeader(204)
	}

	return gin.HandlerFunc(fn)
}
