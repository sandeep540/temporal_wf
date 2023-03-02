package app

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// @@@SNIPSTART
func SignUpWorkflow(ctx workflow.Context, input Init) (string, error) {

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

	withdrawErr := workflow.ExecuteActivity(ctx, SignupUser, input).Get(ctx, &Output)

	if withdrawErr != nil {
		return "", withdrawErr
	}

	return "success", nil
}

// @@@SNIPEND
