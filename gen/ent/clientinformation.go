// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	"github.com/halcyon-org/kizuna/gen/ent/clientinformation"
)

// ClientInformation is the model entity for the ClientInformation schema.
type ClientInformation struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// APIKey holds the value of the "api_key" field.
	APIKey string `json:"api_key,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// LastUsedAt holds the value of the "last_used_at" field.
	LastUsedAt time.Time `json:"last_used_at,omitempty"`
	// LastUpdatedAt holds the value of the "last_updated_at" field.
	LastUpdatedAt time.Time `json:"last_updated_at,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ClientInformation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case clientinformation.FieldID:
			values[i] = new(pulid.ID)
		case clientinformation.FieldUsername, clientinformation.FieldAPIKey:
			values[i] = new(sql.NullString)
		case clientinformation.FieldCreatedAt, clientinformation.FieldLastUsedAt, clientinformation.FieldLastUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ClientInformation fields.
func (ci *ClientInformation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case clientinformation.FieldID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ci.ID = *value
			}
		case clientinformation.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				ci.Username = value.String
			}
		case clientinformation.FieldAPIKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api_key", values[i])
			} else if value.Valid {
				ci.APIKey = value.String
			}
		case clientinformation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ci.CreatedAt = value.Time
			}
		case clientinformation.FieldLastUsedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_used_at", values[i])
			} else if value.Valid {
				ci.LastUsedAt = value.Time
			}
		case clientinformation.FieldLastUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_updated_at", values[i])
			} else if value.Valid {
				ci.LastUpdatedAt = value.Time
			}
		default:
			ci.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ClientInformation.
// This includes values selected through modifiers, order, etc.
func (ci *ClientInformation) Value(name string) (ent.Value, error) {
	return ci.selectValues.Get(name)
}

// Update returns a builder for updating this ClientInformation.
// Note that you need to call ClientInformation.Unwrap() before calling this method if this ClientInformation
// was returned from a transaction, and the transaction was committed or rolled back.
func (ci *ClientInformation) Update() *ClientInformationUpdateOne {
	return NewClientInformationClient(ci.config).UpdateOne(ci)
}

// Unwrap unwraps the ClientInformation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ci *ClientInformation) Unwrap() *ClientInformation {
	_tx, ok := ci.config.driver.(*txDriver)
	if !ok {
		panic("ent: ClientInformation is not a transactional entity")
	}
	ci.config.driver = _tx.drv
	return ci
}

// String implements the fmt.Stringer.
func (ci *ClientInformation) String() string {
	var builder strings.Builder
	builder.WriteString("ClientInformation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ci.ID))
	builder.WriteString("username=")
	builder.WriteString(ci.Username)
	builder.WriteString(", ")
	builder.WriteString("api_key=")
	builder.WriteString(ci.APIKey)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ci.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_used_at=")
	builder.WriteString(ci.LastUsedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_updated_at=")
	builder.WriteString(ci.LastUpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ClientInformations is a parsable slice of ClientInformation.
type ClientInformations []*ClientInformation
