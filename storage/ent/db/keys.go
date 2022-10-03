// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/concourse/dex/storage"
	"github.com/concourse/dex/storage/ent/db/keys"
	jose "gopkg.in/square/go-jose.v2"
)

// Keys is the model entity for the Keys schema.
type Keys struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// VerificationKeys holds the value of the "verification_keys" field.
	VerificationKeys []storage.VerificationKey `json:"verification_keys,omitempty"`
	// SigningKey holds the value of the "signing_key" field.
	SigningKey jose.JSONWebKey `json:"signing_key,omitempty"`
	// SigningKeyPub holds the value of the "signing_key_pub" field.
	SigningKeyPub jose.JSONWebKey `json:"signing_key_pub,omitempty"`
	// NextRotation holds the value of the "next_rotation" field.
	NextRotation time.Time `json:"next_rotation,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Keys) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case keys.FieldVerificationKeys, keys.FieldSigningKey, keys.FieldSigningKeyPub:
			values[i] = new([]byte)
		case keys.FieldID:
			values[i] = new(sql.NullString)
		case keys.FieldNextRotation:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Keys", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Keys fields.
func (k *Keys) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case keys.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				k.ID = value.String
			}
		case keys.FieldVerificationKeys:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field verification_keys", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &k.VerificationKeys); err != nil {
					return fmt.Errorf("unmarshal field verification_keys: %w", err)
				}
			}
		case keys.FieldSigningKey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field signing_key", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &k.SigningKey); err != nil {
					return fmt.Errorf("unmarshal field signing_key: %w", err)
				}
			}
		case keys.FieldSigningKeyPub:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field signing_key_pub", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &k.SigningKeyPub); err != nil {
					return fmt.Errorf("unmarshal field signing_key_pub: %w", err)
				}
			}
		case keys.FieldNextRotation:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field next_rotation", values[i])
			} else if value.Valid {
				k.NextRotation = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Keys.
// Note that you need to call Keys.Unwrap() before calling this method if this Keys
// was returned from a transaction, and the transaction was committed or rolled back.
func (k *Keys) Update() *KeysUpdateOne {
	return (&KeysClient{config: k.config}).UpdateOne(k)
}

// Unwrap unwraps the Keys entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (k *Keys) Unwrap() *Keys {
	_tx, ok := k.config.driver.(*txDriver)
	if !ok {
		panic("db: Keys is not a transactional entity")
	}
	k.config.driver = _tx.drv
	return k
}

// String implements the fmt.Stringer.
func (k *Keys) String() string {
	var builder strings.Builder
	builder.WriteString("Keys(")
	builder.WriteString(fmt.Sprintf("id=%v, ", k.ID))
	builder.WriteString("verification_keys=")
	builder.WriteString(fmt.Sprintf("%v", k.VerificationKeys))
	builder.WriteString(", ")
	builder.WriteString("signing_key=")
	builder.WriteString(fmt.Sprintf("%v", k.SigningKey))
	builder.WriteString(", ")
	builder.WriteString("signing_key_pub=")
	builder.WriteString(fmt.Sprintf("%v", k.SigningKeyPub))
	builder.WriteString(", ")
	builder.WriteString("next_rotation=")
	builder.WriteString(k.NextRotation.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// KeysSlice is a parsable slice of Keys.
type KeysSlice []*Keys

func (k KeysSlice) config(cfg config) {
	for _i := range k {
		k[_i].config = cfg
	}
}
