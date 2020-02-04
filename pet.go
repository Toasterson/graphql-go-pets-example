package graphql_go_pets_example

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

type PetTags []int

func (p PetTags) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *PetTags) Scan(src interface{}) error {
	switch src.(type) {
	case nil:
		return nil
	case []byte:
		return json.Unmarshal(src.([]byte), p)
	default:
		return fmt.Errorf("type %s not supported as pettag in DB", reflect.TypeOf(src).String())
	}
}

// what is needed for a pet
type Pet struct {
	ID    string  `json:"id"`
	Owner string  `json:"owner"`
	Name  string  `json:"name"`
	Tags  PetTags `json:"tags"`
}
