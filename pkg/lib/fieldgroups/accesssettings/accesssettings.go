package accesssettings

import (
	"errors"

	"github.com/creasty/defaults"
)

// AccessSettingsFieldGroup represents the AccessSettingsFieldGroup config fields
type AccessSettingsFieldGroup struct {
	AuthenticationType             string `default:"Database" validate:"" json:"AUTHENTICATION_TYPE" yaml:"AUTHENTICATION_TYPE"`
	FeatureAnonymousAccess         bool   `default:"true" validate:"" json:"FEATURE_ANONYMOUS_ACCESS" yaml:"FEATURE_ANONYMOUS_ACCESS"`
	FeatureDirectLogin             bool   `default:"true" validate:"" json:"FEATURE_DIRECT_LOGIN" yaml:"FEATURE_DIRECT_LOGIN"`
	FeatureGithubLogin             bool   `default:"false" validate:"" json:"FEATURE_GITHUB_LOGIN" yaml:"FEATURE_GITHUB_LOGIN"`
	FeatureGoogleLogin             bool   `default:"false" validate:"" json:"FEATURE_GOOGLE_LOGIN" yaml:"FEATURE_GOOGLE_LOGIN"`
	FeatureInviteOnlyUserCreation  bool   `default:"false" validate:"" json:"FEATURE_INVITE_ONLY_USER_CREATION" yaml:"FEATURE_INVITE_ONLY_USER_CREATION"`
	FeaturePartialUserAutocomplete bool   `default:"true" validate:"" json:"FEATURE_PARTIAL_USER_AUTOCOMPLETE" yaml:"FEATURE_PARTIAL_USER_AUTOCOMPLETE"`
	FeatureUsernameConfirmation    bool   `default:"true" validate:"" json:"FEATURE_USERNAME_CONFIRMATION" yaml:"FEATURE_USERNAME_CONFIRMATION"`
	FeatureUserCreation            bool   `default:"true" validate:"" json:"FEATURE_USER_CREATION" yaml:"FEATURE_USER_CREATION"`
	FeatureUserLastAccessed        bool   `default:"true" validate:"" json:"FEATURE_USER_LAST_ACCESSED" yaml:"FEATURE_USER_LAST_ACCESSED"`
	FeatureUserLogAccess           bool   `default:"false" validate:"" json:"FEATURE_USER_LOG_ACCESS" yaml:"FEATURE_USER_LOG_ACCESS"`
	FeatureUserMetadata            bool   `default:"false" validate:"" json:"FEATURE_USER_METADATA" yaml:"FEATURE_USER_METADATA"`
	FeatureUserRename              bool   `default:"false" validate:"" json:"FEATURE_USER_RENAME" yaml:"FEATURE_USER_RENAME"`
	FreshLoginTimeout              string `default:"10m" validate:"" json:"FRESH_LOGIN_TIMEOUT" yaml:"FRESH_LOGIN_TIMEOUT"`
	UserRecoveryTokenLifetime      string `default:"30m" validate:"" json:"USER_RECOVERY_TOKEN_LIFETIME" yaml:"USER_RECOVERY_TOKEN_LIFETIME"`
}

// NewAccessSettingsFieldGroup creates a new AccessSettingsFieldGroup
func NewAccessSettingsFieldGroup(fullConfig map[string]interface{}) (*AccessSettingsFieldGroup, error) {
	newAccessSettingsFieldGroup := &AccessSettingsFieldGroup{}
	defaults.Set(newAccessSettingsFieldGroup)

	if value, ok := fullConfig["AUTHENTICATION_TYPE"]; ok {
		newAccessSettingsFieldGroup.AuthenticationType, ok = value.(string)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("AUTHENTICATION_TYPE must be of type string")
		}
	}
	if value, ok := fullConfig["FEATURE_ANONYMOUS_ACCESS"]; ok {
		newAccessSettingsFieldGroup.FeatureAnonymousAccess, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_ANONYMOUS_ACCESS must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_DIRECT_LOGIN"]; ok {
		newAccessSettingsFieldGroup.FeatureDirectLogin, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_DIRECT_LOGIN must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_GITHUB_LOGIN"]; ok {
		newAccessSettingsFieldGroup.FeatureGithubLogin, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_GITHUB_LOGIN must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_GOOGLE_LOGIN"]; ok {
		newAccessSettingsFieldGroup.FeatureGoogleLogin, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_GOOGLE_LOGIN must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_INVITE_ONLY_USER_CREATION"]; ok {
		newAccessSettingsFieldGroup.FeatureInviteOnlyUserCreation, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_INVITE_ONLY_USER_CREATION must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_PARTIAL_USER_AUTOCOMPLETE"]; ok {
		newAccessSettingsFieldGroup.FeaturePartialUserAutocomplete, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_PARTIAL_USER_AUTOCOMPLETE must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_USERNAME_CONFIRMATION"]; ok {
		newAccessSettingsFieldGroup.FeatureUsernameConfirmation, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_USERNAME_CONFIRMATION must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_USER_CREATION"]; ok {
		newAccessSettingsFieldGroup.FeatureUserCreation, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_USER_CREATION must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_USER_LAST_ACCESSED"]; ok {
		newAccessSettingsFieldGroup.FeatureUserLastAccessed, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_USER_LAST_ACCESSED must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_USER_LOG_ACCESS"]; ok {
		newAccessSettingsFieldGroup.FeatureUserLogAccess, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_USER_LOG_ACCESS must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_USER_METADATA"]; ok {
		newAccessSettingsFieldGroup.FeatureUserMetadata, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_USER_METADATA must be of type bool")
		}
	}
	if value, ok := fullConfig["FEATURE_USER_RENAME"]; ok {
		newAccessSettingsFieldGroup.FeatureUserRename, ok = value.(bool)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FEATURE_USER_RENAME must be of type bool")
		}
	}
	if value, ok := fullConfig["FRESH_LOGIN_TIMEOUT"]; ok {
		newAccessSettingsFieldGroup.FreshLoginTimeout, ok = value.(string)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("FRESH_LOGIN_TIMEOUT must be of type string")
		}
	}
	if value, ok := fullConfig["USER_RECOVERY_TOKEN_LIFETIME"]; ok {
		newAccessSettingsFieldGroup.UserRecoveryTokenLifetime, ok = value.(string)
		if !ok {
			return newAccessSettingsFieldGroup, errors.New("USER_RECOVERY_TOKEN_LIFETIME must be of type string")
		}
	}

	return newAccessSettingsFieldGroup, nil
}
