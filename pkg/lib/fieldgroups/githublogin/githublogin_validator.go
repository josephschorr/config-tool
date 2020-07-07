package githublogin

import (
	"cuelang.org/go/pkg/strings"
	"github.com/quay/config-tool/pkg/lib/shared"
)

// Validate checks the configuration settings for this field group
func (fg *GitHubLoginFieldGroup) Validate() []shared.ValidationError {

	// Make empty errors
	errors := []shared.ValidationError{}

	// If github trigger is off
	if fg.FeatureGithubLogin == false {
		return errors
	}

	// Check for config
	if fg.GithubLoginConfig == nil {
		newError := shared.ValidationError{
			Tags:    []string{"GITHUB_LOGIN_CONFIG"},
			Policy:  "A is Required",
			Message: "GITHUB_LOGIN_CONFIG is required for GitHubLogin",
		}
		errors = append(errors, newError)
		return errors
	}

	// Check for endpoint
	if fg.GithubLoginConfig.GithubEndpoint == "" {
		newError := shared.ValidationError{
			Tags:    []string{"GITHUB_LOGIN_CONFIG.GITHUB_ENDPOINT"},
			Policy:  "A is Required",
			Message: "GITHUB_LOGIN_CONFIG.GITHUB_ENDPOINT is required for GitHubLogin",
		}
		errors = append(errors, newError)
	}

	// Check for endpoint
	if !strings.HasPrefix(fg.GithubLoginConfig.GithubEndpoint, "http://") && !strings.HasPrefix(fg.GithubLoginConfig.GithubEndpoint, "https://") {
		newError := shared.ValidationError{
			Tags:    []string{"GITHUB_LOGIN_CONFIG.GITHUB_ENDPOINT"},
			Policy:  "A is URL",
			Message: "GITHUB_LOGIN_CONFIG.GITHUB_ENDPOINT must be a url",
		}
		errors = append(errors, newError)
	}

	// Check for client id
	if fg.GithubLoginConfig.ClientId == "" {
		newError := shared.ValidationError{
			Tags:    []string{"GITHUB_LOGIN_CONFIG.CLIENT_ID"},
			Policy:  "A is Required",
			Message: "GITHUB_LOGIN_CONFIG.CLIENT_ID is required for GitHubLogin",
		}
		errors = append(errors, newError)
	}

	// Check for endpoint
	if fg.GithubLoginConfig.ClientSecret == "" {
		newError := shared.ValidationError{
			Tags:    []string{"GITHUB_LOGIN_CONFIG.CLIENT_SECRET"},
			Policy:  "A is Required",
			Message: "GITHUB_LOGIN_CONFIG.CLIENT_SECRET is required for GitHubLogin",
		}
		errors = append(errors, newError)
	}

	// If restricted orgs, make sure
	if fg.GithubLoginConfig.OrgRestrict == true && len(fg.GithubLoginConfig.AllowedOrganizations) == 0 {
		newError := shared.ValidationError{
			Tags:    []string{"GITHUB_LOGIN_CONFIG.ORG_RESTRICT", "GITHUB_LOGIN_CONFIG.ALLOWED_ORGANIZATIONS"},
			Policy:  "A is Required",
			Message: "GITHUB_LOGIN_CONFIG.ALLOWED_ORGANIZATIONS must contain values if GITHUB_LOGIN_CONFIG.ORG_RESTRICT is true",
		}
		errors = append(errors, newError)
	}

	// Check OAuth endpoint
	success := shared.ValidateGitHubOAuth(fg.GithubLoginConfig.ClientId, fg.GithubLoginConfig.ClientSecret)
	if !success {
		newError := shared.ValidationError{
			Tags:    []string{"GITHUB_TRIGGER_CONFIG.CLIENT_ID", "GITHUB_TRIGGER_CONFIG.CLIENT_SECRET"},
			Policy:  "A is Required",
			Message: "Could not verify GitHub OAuth credentials",
		}
		errors = append(errors, newError)
	}

	return errors

}