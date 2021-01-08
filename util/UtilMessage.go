package util

import "github.com/bitly/go-simplejson"

// MessageError -
func MessageError(err error) []byte {

	json := simplejson.New()

	msg := "Não Autorizado!"

	if err.Error() == "Token is expired" {
		msg = "Token expirado!"
	}

	json.Set("message", msg)

	payload, _ := json.MarshalJSON()

	return payload
}

// MessageInfo -
func MessageInfo(name string, v string, msg ...string) []byte {

	json := simplejson.New()
	json.Set(name, v)

	if msg != nil {
		json.Set("message", msg[0])
	}

	payload, _ := json.MarshalJSON()

	return payload
}
