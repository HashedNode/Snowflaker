package structs

import "fmt"

type JsonError struct {
	original []byte
}

func (js JsonError) Error() string {
	return fmt.Sprintf("Received an invalid Snowflake ID %q", string(js.original))
}
