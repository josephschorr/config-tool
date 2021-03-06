package elasticsearch

import (
	"testing"
)

// TestValidateSchema tests the ValidateSchema function
func TestValidateElasticSearch(t *testing.T) {

	// Valid config
	logsModelConfig := map[interface{}]interface{}{
		"elasticsearch_config": map[interface{}]interface{}{
			"access_key": "elastic",
			"secret_key": "nf567zCQgIVMrDvDGIpIHScG",
			"host":       "bfd70499058e4485854f8bacf06af627.us-central1.gcp.cloud.es.io",
			"port":       9243,
		},
	}

	// Define test data
	var tests = []struct {
		name   string
		config map[string]interface{}
		want   string
	}{

		{name: "Valid", config: map[string]interface{}{"LOGS_MODEL": "database"}, want: "valid"},
		{name: "MissingConfig", config: map[string]interface{}{"LOGS_MODEL": "elasticsearch"}, want: "invalid"},
		{name: "ValidConfig", config: map[string]interface{}{"LOGS_MODEL": "elasticsearch", "LOGS_MODEL_CONFIG": logsModelConfig}, want: "valid"},
	}

	// Iterate through tests
	for _, tt := range tests {

		// Run specific test
		t.Run(tt.name, func(t *testing.T) {

			// Get validation result
			fg, err := NewElasticSearchFieldGroup(tt.config)
			if err != nil && tt.want != "typeError" {
				t.Errorf("Expected %s. Received %s", tt.want, err.Error())
			}

			validationErrors := fg.Validate()

			// Get result type
			received := ""
			if len(validationErrors) == 0 {
				received = "valid"
			} else {
				received = "invalid"
			}

			// Compare with expected
			if tt.want != received {
				t.Errorf("Expected %s. Received %s", tt.want, received)
			}

		})
	}

}
