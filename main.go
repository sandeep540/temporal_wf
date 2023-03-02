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

func main() {

	//context := *gin.Context,

	log.Printf("Starting Go Gin Server on Post 9090")

	tc, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	fmt.Printf("Creating Temporal client", tc)

	router := gin.Default()
	router.POST("/initiateWF", initiateWF(tc))
	router.POST("/notify", notify())
	router.GET("/getStatus", getStatus())
	router.Run("localhost:9090")

}

func getStatus() gin.HandlerFunc {

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

		c.IndentedJSON(http.StatusCreated, init)

	}
	return gin.HandlerFunc(fn)
}

func notify() gin.HandlerFunc {

	fn := func(c *gin.Context) {
		workflowId := c.Query("workflowId")
		err := signal.SendNotifySignal(workflowId, true)
		if err != nil {
			log.Fatalln("Unable to notify.", err)
		}
	}

	return gin.HandlerFunc(fn)
}
