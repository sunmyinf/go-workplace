package decode

// Object is wrap type for object name
type Object string

const (
	ObjectApplication             Object = "application"
	ObjectCertificateTransparency Object = "certificate_transparency"
	ObjectGroup                   Object = "group"
	ObjectPage                    Object = "page"
	ObjectPermissions             Object = "permissions"
	ObjectUser                    Object = "user"
	ObjectSecurity                Object = "workplace_security"
)

// RequestBody is struct for json of webhook post request payload
type RequestBody struct {
	Entry  []Entry `json:"entry"`
	Object Object  `json:"object"`
}

// Entry
type Entry struct {
	UID     string   `json:"uid"`
	Changes []Change `json:"changes"`
}

// Change
type Change struct {
	Field string                 `json:"field"`
	Value map[string]interface{} `json:"value"`
}
