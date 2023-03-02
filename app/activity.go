package app

import (
	"context"
	"log"
)

func SignupUser(ctx context.Context, init Init) (string, error) {

	log.Printf("User Details from sign-up %s.\n\n", init)
	return "success", nil
}
