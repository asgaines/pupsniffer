package null

import (
	"encoding/json"
	"fmt"
)

// Null is how the humane society's API represents a JSON null
const Null = `{"0":"\n  "}`

// String allows for a typically string-typed field to be null (as defined by Null const)
type String struct {
	Value string
	Valid bool
}

// UnmarshalJSON satisfies the json.Unmarshaler interface, used to unmarshal a NullString field
func (n *String) UnmarshalJSON(data []byte) error {
	var s string

	err := json.Unmarshal(data, &s)
	if err != nil {
		if string(data) == Null {
			(*n).Valid = false
			return nil
		} else {
			return fmt.Errorf("Invalid value for NullString: %s. Error: %s", data, err)
		}
	}

	(*n).Value = s
	(*n).Valid = true

	return nil
}
