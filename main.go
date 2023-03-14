package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"temporal_wf/model"
	"temporal_wf/signal"
	"temporal_wf/workflow"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"gopkg.in/yaml.v2"
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
	router.GET("/initiateDynamicWF", initiateDynamicWF(tc))
	router.GET("/notify", notify(tc))
	router.GET("/getStatus", getStatus(tc))
	router.GET("/getWorkflowStatus", getWorkflowStatus(tc))
	router.GET("/showInterceptor", showInterceptor(tc))
	router.Run("localhost:9090")

}

func getStatus(tc client.Client) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		c.IndentedJSON(http.StatusOK, gin.H{"message": "active"})
	}
	return gin.HandlerFunc(fn)

}

func getWorkflowStatus(tc client.Client) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		workflowId := c.Query("workflowId")
		log.Println(workflowId)

		runId := c.Query("runId")
		log.Println(runId)

		queryType := "__stack_trace"

		response, err := tc.QueryWorkflow(context.Background(), workflowId, runId, queryType)
		if err != nil {
			log.Fatalln("Unable to query workflow", err)
		}
		//log.Println("response ---> ", response.Get(""))

		var result interface{}
		if err := response.Get(&result); err != nil {
			log.Fatalln("Unable to decode query result", err)
		}
		log.Println("Received query result", "Result", result)

		c.IndentedJSON(http.StatusOK, gin.H{"message": "active"})
	}
	return gin.HandlerFunc(fn)

}

func initiateWF(tc client.Client) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		var init model.Init
		if err := c.BindJSON(&init); err != nil {
			return
		}
		log.Printf("Session init %s.\n\n", init)

		id := (uuid.New()).String()

		options := client.StartWorkflowOptions{
			ID:        id,
			TaskQueue: model.SignupUserQueue,
		}

		we, err := tc.ExecuteWorkflow(context.Background(), options, workflow.SignUpWorkflow, init)
		if err != nil {
			log.Fatalln("Unable to start the Workflow:", err)
		}

		//log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())
		log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

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

func showInterceptor(tc client.Client) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		workflowId := c.Query("workflowId")
		log.Println(workflowId)

		content, err := os.ReadFile("data/interceptor.json")
		if err != nil {
			log.Fatalln(err)
		}

		var data model.Interceptor
		err = json.Unmarshal(content, &data)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Data is ", data)

		for _, v := range data {
			if v.WorkflowId == workflowId {
				c.JSON(200, v)
				break
			}
		}

	}
	return gin.HandlerFunc(fn)
}

func initiateDynamicWF(tc client.Client) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		data, err := os.ReadFile("workflow.yml")
		if err != nil {
			log.Fatalln("failed to load dsl config file", err)
		}
		var dslWorkflow workflow.Workflow
		if err := yaml.Unmarshal(data, &dslWorkflow); err != nil {
			log.Fatalln("failed to unmarshal dsl config", err)
		}

		id := (uuid.New()).String()

		workflowOptions := client.StartWorkflowOptions{
			ID:        id,
			TaskQueue: model.SignupUserQueue,
		}

		we, err := tc.ExecuteWorkflow(context.Background(), workflowOptions, workflow.SimpleDSLWorkflow, dslWorkflow)
		if err != nil {
			log.Fatalln("Unable to execute workflow", err)
		}
		log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

		c.IndentedJSON(http.StatusOK, id)

	}
	return gin.HandlerFunc(fn)
}
