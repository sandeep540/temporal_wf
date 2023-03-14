package activity

import (
	"context"
	"log"
)

type SampleActivities struct {
}

func (a *SampleActivities) StartActivity(ctx context.Context) (string, error) {

	log.Printf("StartActivity")
	return "Init Activity Completed", nil
}

func (a *SampleActivities) CreateTenet(ctx context.Context) (string, error) {
	log.Printf("CreateTenet")

	return "success", nil
}

func (a *SampleActivities) FinishActivity(ctx context.Context) (string, error) {
	log.Printf("FinishActivity")

	return "Show Pricing Plans for User", nil
}
