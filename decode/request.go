package decode

type RequestBody struct {
	Data []Data `json: data`
}

type Object string

const (
	ObjectPage     Object = "page"
	ObjectGroup    Object = "group"
	ObjectUser     Object = "user"
	ObjectSecurity Object = "security"
)

type Data struct {
	Object      object `json: object`
	CallbackUrl string `json: callback_url`
	Active      bool   `json: active`
}
