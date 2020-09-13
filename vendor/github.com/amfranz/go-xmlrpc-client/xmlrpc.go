package xmlrpc

// Struct presents hash type used in xmlprc requests and responses.
type Struct map[string]interface{}

// Base64 represents base64 data
type Base64 string

// Params represents a list of parameters to a method.
type Params struct {
	Params []interface{}
}
