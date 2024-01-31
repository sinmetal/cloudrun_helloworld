package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	metadatabox "github.com/sinmetalcraft/gcpbox/metadata/cloudrun"
	"google.golang.org/api/idtoken"
)

type IAPHandler struct {
}

func (h *IAPHandler) Handle(w http.ResponseWriter, r *http.Request) {
	const backendServiceID = "8877965452879686025"
	iapJWT := r.Header.Get("x-goog-iap-jwt-assertion")
	projectNumber, err := metadatabox.NumericProjectID()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(iapJWT)
	fmt.Println(projectNumber)
	if err := validateJWTFromComputeEngine(w, iapJWT, projectNumber, backendServiceID); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
}

// validateJWTFromComputeEngine validates a JWT found in the
// "x-goog-iap-jwt-assertion" header.
func validateJWTFromComputeEngine(w io.Writer, iapJWT, projectNumber, backendServiceID string) error {
	// iapJWT := "YmFzZQ==.ZW5jb2RlZA==.and0" // req.Header.Get("X-Goog-IAP-JWT-Assertion")
	// projectNumber := "123456789"
	// backendServiceID := "backend-service-id"
	ctx := context.Background()
	aud := fmt.Sprintf("/projects/%s/global/backendServices/%s", projectNumber, backendServiceID)

	payload, err := idtoken.Validate(ctx, iapJWT, aud)
	if err != nil {
		return fmt.Errorf("idtoken.Validate: %w", err)
	}

	// payload contains the JWT claims for further inspection or validation
	fmt.Fprintf(w, "payload: %v", payload)

	return nil
}
