package main

import (
	"encoding/json"
	"fmt"

	"github.com/sunmyinf/go-workplace/decode"
)

func main() {
	body := decode.RequestBody{}
	jsonByte := []byte(`{
  "entry": [
    {
      "time": 1520383571,
      "changes": [
        {
          "field": "photos",
          "value":
            {
              "verb": "update",
              "object_id": "10211885744794461"
            }
        }
      ],
      "id": "10210299214172187",
      "uid": "10210299214172187"
    }
  ],
  "object": "user"
}`)
	json.Unmarshal(jsonByte, &body)
	fmt.Println(body)
}
