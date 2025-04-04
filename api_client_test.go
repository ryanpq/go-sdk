/**
 * Go SDK for OpenFGA
 *
 * API version: 1.x
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://openfga.dev/community
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"github.com/openfga/go-sdk/telemetry"
	"net/http"
	"testing"
	"time"
)

func TestApiClientCreatedWithDefaultTelemetry(t *testing.T) {
	cfg := Configuration{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		ApiUrl:     "http://localhost:8080/",
	}
	_ = NewAPIClient(&cfg)

	telemetry1 := telemetry.Get(telemetry.TelemetryFactoryParameters{Configuration: cfg.Telemetry})
	telemetry2 := telemetry.Get(telemetry.TelemetryFactoryParameters{Configuration: cfg.Telemetry})

	if telemetry1 != telemetry2 {
		t.Fatalf("Telemetry instance should be the same")
	}
}
