package main

import (
	"log"
	"temporal_wf/activity"
	app "temporal_wf/model"
	"temporal_wf/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	log.Println("Worker Starting...")
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, app.SignupUserQueue, worker.Options{})
	w.RegisterWorkflow(workflow.SignUpWorkflow)
	w.RegisterActivity(activity.InitActivity)
	w.RegisterActivity(activity.SignUp)
	w.RegisterActivity(activity.ShowPricingPlans)
	w.RegisterActivity(activity.CompleteWF)

	w.RegisterWorkflow(workflow.SimpleDSLWorkflow)
	w.RegisterActivity(&activity.SampleActivities{})

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}

	//w1 := worker.New(c, "DSL", worker.Options{})

	/* if err := w1.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	} */
}
