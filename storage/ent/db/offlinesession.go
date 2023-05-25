// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/concourse/dex/storage/ent/db/offlinesession"
)

// OfflineSession is the model entity for the OfflineSession schema.
type OfflineSession struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// ConnID holds the value of the "conn_id" field.
	ConnID string `json:"conn_id,omitempty"`
	// Refresh holds the value of the "refresh" field.
	Refresh []byte `json:"refresh,omitempty"`
	// ConnectorData holds the value of the "connector_data" field.
	ConnectorData *[]byte `json:"connector_data,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OfflineSession) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case offlinesession.FieldRefresh, offlinesession.FieldConnectorData:
			values[i] = new([]byte)
		case offlinesession.FieldID, offlinesession.FieldUserID, offlinesession.FieldConnID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OfflineSession", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OfflineSession fields.
func (os *OfflineSession) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case offlinesession.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				os.ID = value.String
			}
		case offlinesession.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				os.UserID = value.String
			}
		case offlinesession.FieldConnID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field conn_id", values[i])
			} else if value.Valid {
				os.ConnID = value.String
			}
		case offlinesession.FieldRefresh:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field refresh", values[i])
			} else if value != nil {
				os.Refresh = *value
			}
		case offlinesession.FieldConnectorData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field connector_data", values[i])
			} else if value != nil {
				os.ConnectorData = value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this OfflineSession.
// Note that you need to call OfflineSession.Unwrap() before calling this method if this OfflineSession
// was returned from a transaction, and the transaction was committed or rolled back.
func (os *OfflineSession) Update() *OfflineSessionUpdateOne {
	return NewOfflineSessionClient(os.config).UpdateOne(os)
}

// Unwrap unwraps the OfflineSession entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (os *OfflineSession) Unwrap() *OfflineSession {
	_tx, ok := os.config.driver.(*txDriver)
	if !ok {
		panic("db: OfflineSession is not a transactional entity")
	}
	os.config.driver = _tx.drv
	return os
}

// String implements the fmt.Stringer.
func (os *OfflineSession) String() string {
	var builder strings.Builder
	builder.WriteString("OfflineSession(")
	builder.WriteString(fmt.Sprintf("id=%v, ", os.ID))
	builder.WriteString("user_id=")
	builder.WriteString(os.UserID)
	builder.WriteString(", ")
	builder.WriteString("conn_id=")
	builder.WriteString(os.ConnID)
	builder.WriteString(", ")
	builder.WriteString("refresh=")
	builder.WriteString(fmt.Sprintf("%v", os.Refresh))
	builder.WriteString(", ")
	if v := os.ConnectorData; v != nil {
		builder.WriteString("connector_data=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// OfflineSessions is a parsable slice of OfflineSession.
type OfflineSessions []*OfflineSession
