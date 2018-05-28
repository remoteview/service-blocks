package blocks

import (
	"fmt"
	"time"
)

// Block - working block
type Block struct {
	ID        string   `json:"id,omitempty"`
	StartTime JSONTime `json:"start_time,omitempty"`
	EndTime   JSONTime `json:"end_time,omitempty"`
	Timezone  string   `json:"timezone,omitempty"`
	UserID    string   `json:"user_id,omitempty"`
}

// Marshaler -
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// JSONTime -
type JSONTime time.Time

// MarshalJSON -
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("15:04"))
	return []byte(stamp), nil
}
