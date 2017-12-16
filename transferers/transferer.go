package transferers

// Transferer interface to allow binding a Request to a Struct
type Transferer interface {
	GetValues() map[string]interface{}
}
