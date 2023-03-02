package signal

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

const (
	NOTIFICATION_SIGNAL = "notification_signal"
)

func SendNotifySignal(workflowID string, status bool) (err error) {
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
		return
	}

	err = temporalClient.SignalWorkflow(context.Background(), workflowID, "", NOTIFICATION_SIGNAL, status)
	if err != nil {
		log.Fatalln("Error signaling client", err)
		return
	}

	return nil
}

func ReciveSignal(ctx workflow.Context, signalName string) (status bool) {
	workflow.GetSignalChannel(ctx, signalName).Receive(ctx, &status)
	return
}
