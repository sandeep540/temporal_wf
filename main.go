package main

import (
	"context"
	"log"
	"net/http"
	"temporal_wf/app"

	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
)

func main() {

	log.Printf("Starting Go Gin Server on Post 9090")

	router := gin.Default()
	router.POST("/initiateWF", initiateWF)
	//router.POST("/signUp", signUp)
	//router.POST("/createUser", createUser)
	//router.GET("/getStatus", getStatus)
	//router.POST("/completeWF", completeWF)
	router.Run("localhost:9090")

}

func initiateWF(c *gin.Context) {

	var init app.Init

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&init); err != nil {
		return
	}
	log.Printf("Session init %s.\n\n", init)

	tc, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	//defer tc.Close()

	options := client.StartWorkflowOptions{
		ID:        "signup-701",
		TaskQueue: app.SignupUserQueue,
	}

	we, err := tc.ExecuteWorkflow(context.Background(), options, app.SignUpWorkflow, init)
	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	c.IndentedJSON(http.StatusCreated, init)
}
