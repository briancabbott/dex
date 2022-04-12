// Code generated by entc, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/briancabbott/dex/storage/ent/db/authrequest"
	"github.com/briancabbott/entgo/dialect/sql"
)

// AuthRequest is the model entity for the AuthRequest schema.
type AuthRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID string `json:"client_id,omitempty"`
	// Scopes holds the value of the "scopes" field.
	Scopes []string `json:"scopes,omitempty"`
	// ResponseTypes holds the value of the "response_types" field.
	ResponseTypes []string `json:"response_types,omitempty"`
	// RedirectURI holds the value of the "redirect_uri" field.
	RedirectURI string `json:"redirect_uri,omitempty"`
	// Nonce holds the value of the "nonce" field.
	Nonce string `json:"nonce,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// ForceApprovalPrompt holds the value of the "force_approval_prompt" field.
	ForceApprovalPrompt bool `json:"force_approval_prompt,omitempty"`
	// LoggedIn holds the value of the "logged_in" field.
	LoggedIn bool `json:"logged_in,omitempty"`
	// ClaimsUserID holds the value of the "claims_user_id" field.
	ClaimsUserID string `json:"claims_user_id,omitempty"`
	// ClaimsUsername holds the value of the "claims_username" field.
	ClaimsUsername string `json:"claims_username,omitempty"`
	// ClaimsEmail holds the value of the "claims_email" field.
	ClaimsEmail string `json:"claims_email,omitempty"`
	// ClaimsEmailVerified holds the value of the "claims_email_verified" field.
	ClaimsEmailVerified bool `json:"claims_email_verified,omitempty"`
	// ClaimsGroups holds the value of the "claims_groups" field.
	ClaimsGroups []string `json:"claims_groups,omitempty"`
	// ClaimsPreferredUsername holds the value of the "claims_preferred_username" field.
	ClaimsPreferredUsername string `json:"claims_preferred_username,omitempty"`
	// ConnectorID holds the value of the "connector_id" field.
	ConnectorID string `json:"connector_id,omitempty"`
	// ConnectorData holds the value of the "connector_data" field.
	ConnectorData *[]byte `json:"connector_data,omitempty"`
	// Expiry holds the value of the "expiry" field.
	Expiry time.Time `json:"expiry,omitempty"`
	// CodeChallenge holds the value of the "code_challenge" field.
	CodeChallenge string `json:"code_challenge,omitempty"`
	// CodeChallengeMethod holds the value of the "code_challenge_method" field.
	CodeChallengeMethod string `json:"code_challenge_method,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthRequest) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case authrequest.FieldScopes, authrequest.FieldResponseTypes, authrequest.FieldClaimsGroups, authrequest.FieldConnectorData:
			values[i] = new([]byte)
		case authrequest.FieldForceApprovalPrompt, authrequest.FieldLoggedIn, authrequest.FieldClaimsEmailVerified:
			values[i] = new(sql.NullBool)
		case authrequest.FieldID, authrequest.FieldClientID, authrequest.FieldRedirectURI, authrequest.FieldNonce, authrequest.FieldState, authrequest.FieldClaimsUserID, authrequest.FieldClaimsUsername, authrequest.FieldClaimsEmail, authrequest.FieldClaimsPreferredUsername, authrequest.FieldConnectorID, authrequest.FieldCodeChallenge, authrequest.FieldCodeChallengeMethod:
			values[i] = new(sql.NullString)
		case authrequest.FieldExpiry:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AuthRequest", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthRequest fields.
func (ar *AuthRequest) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authrequest.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ar.ID = value.String
			}
		case authrequest.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				ar.ClientID = value.String
			}
		case authrequest.FieldScopes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field scopes", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ar.Scopes); err != nil {
					return fmt.Errorf("unmarshal field scopes: %w", err)
				}
			}
		case authrequest.FieldResponseTypes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field response_types", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ar.ResponseTypes); err != nil {
					return fmt.Errorf("unmarshal field response_types: %w", err)
				}
			}
		case authrequest.FieldRedirectURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect_uri", values[i])
			} else if value.Valid {
				ar.RedirectURI = value.String
			}
		case authrequest.FieldNonce:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nonce", values[i])
			} else if value.Valid {
				ar.Nonce = value.String
			}
		case authrequest.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				ar.State = value.String
			}
		case authrequest.FieldForceApprovalPrompt:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field force_approval_prompt", values[i])
			} else if value.Valid {
				ar.ForceApprovalPrompt = value.Bool
			}
		case authrequest.FieldLoggedIn:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field logged_in", values[i])
			} else if value.Valid {
				ar.LoggedIn = value.Bool
			}
		case authrequest.FieldClaimsUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field claims_user_id", values[i])
			} else if value.Valid {
				ar.ClaimsUserID = value.String
			}
		case authrequest.FieldClaimsUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field claims_username", values[i])
			} else if value.Valid {
				ar.ClaimsUsername = value.String
			}
		case authrequest.FieldClaimsEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field claims_email", values[i])
			} else if value.Valid {
				ar.ClaimsEmail = value.String
			}
		case authrequest.FieldClaimsEmailVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field claims_email_verified", values[i])
			} else if value.Valid {
				ar.ClaimsEmailVerified = value.Bool
			}
		case authrequest.FieldClaimsGroups:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field claims_groups", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ar.ClaimsGroups); err != nil {
					return fmt.Errorf("unmarshal field claims_groups: %w", err)
				}
			}
		case authrequest.FieldClaimsPreferredUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field claims_preferred_username", values[i])
			} else if value.Valid {
				ar.ClaimsPreferredUsername = value.String
			}
		case authrequest.FieldConnectorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field connector_id", values[i])
			} else if value.Valid {
				ar.ConnectorID = value.String
			}
		case authrequest.FieldConnectorData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field connector_data", values[i])
			} else if value != nil {
				ar.ConnectorData = value
			}
		case authrequest.FieldExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiry", values[i])
			} else if value.Valid {
				ar.Expiry = value.Time
			}
		case authrequest.FieldCodeChallenge:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code_challenge", values[i])
			} else if value.Valid {
				ar.CodeChallenge = value.String
			}
		case authrequest.FieldCodeChallengeMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code_challenge_method", values[i])
			} else if value.Valid {
				ar.CodeChallengeMethod = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AuthRequest.
// Note that you need to call AuthRequest.Unwrap() before calling this method if this AuthRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *AuthRequest) Update() *AuthRequestUpdateOne {
	return (&AuthRequestClient{config: ar.config}).UpdateOne(ar)
}

// Unwrap unwraps the AuthRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ar *AuthRequest) Unwrap() *AuthRequest {
	tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("db: AuthRequest is not a transactional entity")
	}
	ar.config.driver = tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *AuthRequest) String() string {
	var builder strings.Builder
	builder.WriteString("AuthRequest(")
	builder.WriteString(fmt.Sprintf("id=%v", ar.ID))
	builder.WriteString(", client_id=")
	builder.WriteString(ar.ClientID)
	builder.WriteString(", scopes=")
	builder.WriteString(fmt.Sprintf("%v", ar.Scopes))
	builder.WriteString(", response_types=")
	builder.WriteString(fmt.Sprintf("%v", ar.ResponseTypes))
	builder.WriteString(", redirect_uri=")
	builder.WriteString(ar.RedirectURI)
	builder.WriteString(", nonce=")
	builder.WriteString(ar.Nonce)
	builder.WriteString(", state=")
	builder.WriteString(ar.State)
	builder.WriteString(", force_approval_prompt=")
	builder.WriteString(fmt.Sprintf("%v", ar.ForceApprovalPrompt))
	builder.WriteString(", logged_in=")
	builder.WriteString(fmt.Sprintf("%v", ar.LoggedIn))
	builder.WriteString(", claims_user_id=")
	builder.WriteString(ar.ClaimsUserID)
	builder.WriteString(", claims_username=")
	builder.WriteString(ar.ClaimsUsername)
	builder.WriteString(", claims_email=")
	builder.WriteString(ar.ClaimsEmail)
	builder.WriteString(", claims_email_verified=")
	builder.WriteString(fmt.Sprintf("%v", ar.ClaimsEmailVerified))
	builder.WriteString(", claims_groups=")
	builder.WriteString(fmt.Sprintf("%v", ar.ClaimsGroups))
	builder.WriteString(", claims_preferred_username=")
	builder.WriteString(ar.ClaimsPreferredUsername)
	builder.WriteString(", connector_id=")
	builder.WriteString(ar.ConnectorID)
	if v := ar.ConnectorData; v != nil {
		builder.WriteString(", connector_data=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", expiry=")
	builder.WriteString(ar.Expiry.Format(time.ANSIC))
	builder.WriteString(", code_challenge=")
	builder.WriteString(ar.CodeChallenge)
	builder.WriteString(", code_challenge_method=")
	builder.WriteString(ar.CodeChallengeMethod)
	builder.WriteByte(')')
	return builder.String()
}

// AuthRequests is a parsable slice of AuthRequest.
type AuthRequests []*AuthRequest

func (ar AuthRequests) config(cfg config) {
	for _i := range ar {
		ar[_i].config = cfg
	}
}
