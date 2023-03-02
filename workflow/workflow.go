package workflow

import (
	"log"
	"temporal_wf/activity"
	"temporal_wf/app"
	"temporal_wf/signal"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type UserDetails struct {
	name    string
	email   string
	country string
}

// @@@SNIPSTART
func SignUpWorkflow(ctx workflow.Context, input app.Init) (string, error) {

	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        0, // unlimited retries
		NonRetryableErrorTypes: []string{"InvalidUser", "InvalidSession"},
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)

	var Output string

	initErr := workflow.ExecuteActivity(ctx, activity.InitActivity, input).Get(ctx, &Output)

	if initErr != nil {
		return "", initErr
	}

	user := &UserDetails{
		name:    "sandeep",
		email:   "sandeep.kongathi@infracloud.io",
		country: "India",
	}

	signupErr := workflow.ExecuteActivity(ctx, activity.SignUp, user).Get(ctx, &Output)

	if signupErr != nil {
		return "", signupErr
	}

	workflow.Sleep(ctx, 5*time.Second)
	log.Println("Waiting for Notification")
	// Customer paid bill
	if status := signal.ReciveSignal(ctx, signal.NOTIFICATION_SIGNAL); !status {
		log.Println("Workflow couldn't be completed! ")

	}
	workflow.ExecuteActivity(ctx, activity.CompleteWF, nil).Get(ctx, nil)
	return "success", nil
}

// @@@SNIPEND
