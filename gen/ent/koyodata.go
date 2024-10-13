// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	"github.com/halcyon-org/kizuna/gen/ent/koyodata"
)

// KoyoData is the model entity for the KoyoData schema.
type KoyoData struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
	// KoyoID holds the value of the "koyo_id" field.
	KoyoID pulid.ID `json:"koyo_id,omitempty"`
	// Scale holds the value of the "scale" field.
	Scale float64 `json:"scale,omitempty"`
	// Params holds the value of the "params" field.
	Params map[string]string `json:"params,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// EntryAt holds the value of the "entry_at" field.
	EntryAt time.Time `json:"entry_at,omitempty"`
	// TargetAt holds the value of the "target_at" field.
	TargetAt              time.Time `json:"target_at,omitempty"`
	koyo_information_data *pulid.ID
	selectValues          sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KoyoData) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case koyodata.FieldParams:
			values[i] = new([]byte)
		case koyodata.FieldID, koyodata.FieldKoyoID:
			values[i] = new(pulid.ID)
		case koyodata.FieldScale:
			values[i] = new(sql.NullFloat64)
		case koyodata.FieldVersion:
			values[i] = new(sql.NullString)
		case koyodata.FieldEntryAt, koyodata.FieldTargetAt:
			values[i] = new(sql.NullTime)
		case koyodata.ForeignKeys[0]: // koyo_information_data
			values[i] = &sql.NullScanner{S: new(pulid.ID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KoyoData fields.
func (kd *KoyoData) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case koyodata.FieldID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				kd.ID = *value
			}
		case koyodata.FieldKoyoID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field koyo_id", values[i])
			} else if value != nil {
				kd.KoyoID = *value
			}
		case koyodata.FieldScale:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field scale", values[i])
			} else if value.Valid {
				kd.Scale = value.Float64
			}
		case koyodata.FieldParams:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field params", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &kd.Params); err != nil {
					return fmt.Errorf("unmarshal field params: %w", err)
				}
			}
		case koyodata.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				kd.Version = value.String
			}
		case koyodata.FieldEntryAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field entry_at", values[i])
			} else if value.Valid {
				kd.EntryAt = value.Time
			}
		case koyodata.FieldTargetAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field target_at", values[i])
			} else if value.Valid {
				kd.TargetAt = value.Time
			}
		case koyodata.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field koyo_information_data", values[i])
			} else if value.Valid {
				kd.koyo_information_data = new(pulid.ID)
				*kd.koyo_information_data = *value.S.(*pulid.ID)
			}
		default:
			kd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the KoyoData.
// This includes values selected through modifiers, order, etc.
func (kd *KoyoData) Value(name string) (ent.Value, error) {
	return kd.selectValues.Get(name)
}

// Update returns a builder for updating this KoyoData.
// Note that you need to call KoyoData.Unwrap() before calling this method if this KoyoData
// was returned from a transaction, and the transaction was committed or rolled back.
func (kd *KoyoData) Update() *KoyoDataUpdateOne {
	return NewKoyoDataClient(kd.config).UpdateOne(kd)
}

// Unwrap unwraps the KoyoData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kd *KoyoData) Unwrap() *KoyoData {
	_tx, ok := kd.config.driver.(*txDriver)
	if !ok {
		panic("ent: KoyoData is not a transactional entity")
	}
	kd.config.driver = _tx.drv
	return kd
}

// String implements the fmt.Stringer.
func (kd *KoyoData) String() string {
	var builder strings.Builder
	builder.WriteString("KoyoData(")
	builder.WriteString(fmt.Sprintf("id=%v, ", kd.ID))
	builder.WriteString("koyo_id=")
	builder.WriteString(fmt.Sprintf("%v", kd.KoyoID))
	builder.WriteString(", ")
	builder.WriteString("scale=")
	builder.WriteString(fmt.Sprintf("%v", kd.Scale))
	builder.WriteString(", ")
	builder.WriteString("params=")
	builder.WriteString(fmt.Sprintf("%v", kd.Params))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(kd.Version)
	builder.WriteString(", ")
	builder.WriteString("entry_at=")
	builder.WriteString(kd.EntryAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("target_at=")
	builder.WriteString(kd.TargetAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// KoyoDataSlice is a parsable slice of KoyoData.
type KoyoDataSlice []*KoyoData