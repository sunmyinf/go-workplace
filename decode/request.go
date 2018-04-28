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
	ID      string   `json:"id"`
	UID     string   `json:"uid"`
	Time    int64    `json:"time"`
	Changes []Change `json:"changes"`
}

// Change
type Change struct {
	Field Field       `json:"field"`
	Value interface{} `json:"value"`
}

// UnmarshalJSON is custom json parser for Value
func (c *Change) UnmarshalJSON(data []byte) error {

}
