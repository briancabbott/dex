// Code generated by entc, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/briancabbott/dex/storage/ent/db/refreshtoken"
)

// RefreshToken is the model entity for the RefreshToken schema.
type RefreshToken struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID string `json:"client_id,omitempty"`
	// Scopes holds the value of the "scopes" field.
	Scopes []string `json:"scopes,omitempty"`
	// Nonce holds the value of the "nonce" field.
	Nonce string `json:"nonce,omitempty"`
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
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// ObsoleteToken holds the value of the "obsolete_token" field.
	ObsoleteToken string `json:"obsolete_token,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// LastUsed holds the value of the "last_used" field.
	LastUsed time.Time `json:"last_used,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RefreshToken) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case refreshtoken.FieldScopes, refreshtoken.FieldClaimsGroups, refreshtoken.FieldConnectorData:
			values[i] = new([]byte)
		case refreshtoken.FieldClaimsEmailVerified:
			values[i] = new(sql.NullBool)
		case refreshtoken.FieldID, refreshtoken.FieldClientID, refreshtoken.FieldNonce, refreshtoken.FieldClaimsUserID, refreshtoken.FieldClaimsUsername, refreshtoken.FieldClaimsEmail, refreshtoken.FieldClaimsPreferredUsername, refreshtoken.FieldConnectorID, refreshtoken.FieldToken, refreshtoken.FieldObsoleteToken:
			values[i] = new(sql.NullString)
		case refreshtoken.FieldCreatedAt, refreshtoken.FieldLastUsed:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RefreshToken", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RefreshToken fields.
func (rt *RefreshToken) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case refreshtoken.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				rt.ID = value.String
			}
		case refreshtoken.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				rt.ClientID = value.String
			}
		case refreshtoken.FieldScopes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field scopes", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rt.Scopes); err != nil {
					return fmt.Errorf("unmarshal field scopes: %w", err)
				}
			}
		case refreshtoken.FieldNonce:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nonce", values[i])
			} else if value.Valid {
				rt.Nonce = value.String
			}
		case refreshtoken.FieldClaimsUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field claims_user_id", values[i])
			} else if value.Valid {
				rt.ClaimsUserID = value.String
			}
		case refreshtoken.FieldClaimsUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field claims_username", values[i])
			} else if value.Valid {
				rt.ClaimsUsername = value.String
			}
		case refreshtoken.FieldClaimsEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field claims_email", values[i])
			} else if value.Valid {
				rt.ClaimsEmail = value.String
			}
		case refreshtoken.FieldClaimsEmailVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field claims_email_verified", values[i])
			} else if value.Valid {
				rt.ClaimsEmailVerified = value.Bool
			}
		case refreshtoken.FieldClaimsGroups:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field claims_groups", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rt.ClaimsGroups); err != nil {
					return fmt.Errorf("unmarshal field claims_groups: %w", err)
				}
			}
		case refreshtoken.FieldClaimsPreferredUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field claims_preferred_username", values[i])
			} else if value.Valid {
				rt.ClaimsPreferredUsername = value.String
			}
		case refreshtoken.FieldConnectorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field connector_id", values[i])
			} else if value.Valid {
				rt.ConnectorID = value.String
			}
		case refreshtoken.FieldConnectorData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field connector_data", values[i])
			} else if value != nil {
				rt.ConnectorData = value
			}
		case refreshtoken.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				rt.Token = value.String
			}
		case refreshtoken.FieldObsoleteToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field obsolete_token", values[i])
			} else if value.Valid {
				rt.ObsoleteToken = value.String
			}
		case refreshtoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rt.CreatedAt = value.Time
			}
		case refreshtoken.FieldLastUsed:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_used", values[i])
			} else if value.Valid {
				rt.LastUsed = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this RefreshToken.
// Note that you need to call RefreshToken.Unwrap() before calling this method if this RefreshToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (rt *RefreshToken) Update() *RefreshTokenUpdateOne {
	return (&RefreshTokenClient{config: rt.config}).UpdateOne(rt)
}

// Unwrap unwraps the RefreshToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rt *RefreshToken) Unwrap() *RefreshToken {
	tx, ok := rt.config.driver.(*txDriver)
	if !ok {
		panic("db: RefreshToken is not a transactional entity")
	}
	rt.config.driver = tx.drv
	return rt
}

// String implements the fmt.Stringer.
func (rt *RefreshToken) String() string {
	var builder strings.Builder
	builder.WriteString("RefreshToken(")
	builder.WriteString(fmt.Sprintf("id=%v", rt.ID))
	builder.WriteString(", client_id=")
	builder.WriteString(rt.ClientID)
	builder.WriteString(", scopes=")
	builder.WriteString(fmt.Sprintf("%v", rt.Scopes))
	builder.WriteString(", nonce=")
	builder.WriteString(rt.Nonce)
	builder.WriteString(", claims_user_id=")
	builder.WriteString(rt.ClaimsUserID)
	builder.WriteString(", claims_username=")
	builder.WriteString(rt.ClaimsUsername)
	builder.WriteString(", claims_email=")
	builder.WriteString(rt.ClaimsEmail)
	builder.WriteString(", claims_email_verified=")
	builder.WriteString(fmt.Sprintf("%v", rt.ClaimsEmailVerified))
	builder.WriteString(", claims_groups=")
	builder.WriteString(fmt.Sprintf("%v", rt.ClaimsGroups))
	builder.WriteString(", claims_preferred_username=")
	builder.WriteString(rt.ClaimsPreferredUsername)
	builder.WriteString(", connector_id=")
	builder.WriteString(rt.ConnectorID)
	if v := rt.ConnectorData; v != nil {
		builder.WriteString(", connector_data=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", token=")
	builder.WriteString(rt.Token)
	builder.WriteString(", obsolete_token=")
	builder.WriteString(rt.ObsoleteToken)
	builder.WriteString(", created_at=")
	builder.WriteString(rt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", last_used=")
	builder.WriteString(rt.LastUsed.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// RefreshTokens is a parsable slice of RefreshToken.
type RefreshTokens []*RefreshToken

func (rt RefreshTokens) config(cfg config) {
	for _i := range rt {
		rt[_i].config = cfg
	}
}
