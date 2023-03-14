package activity

import (
	"context"
	"log"
	"temporal_wf/model"
)

func InitActivity(ctx context.Context, init model.Init) (string, error) {

	log.Printf("init Details from init %s.\n\n", init)
	return "Init Activity Completed", nil
}

func SignUp(ctx context.Context) (string, error) {

	log.Printf("User Details from sign-up")
	return "SignUp Activity Initiated", nil
}

func CreateUser(ctx context.Context, init model.Init) (string, error) {
	log.Printf("Confirmation of Create user %s.\n\n", init)

	return "success", nil
}

func ShowPricingPlans(ctx context.Context) (string, error) {
	log.Printf("Show Pricing Plans for User")

	return "Show Pricing Plans for User", nil
}

func CompleteWF(ctx context.Context) (string, error) {
	log.Printf("Closing the Workflow ")

	return "Closing the Workflow", nil
}
