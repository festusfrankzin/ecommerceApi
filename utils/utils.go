package utils 

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func ParseJson(r *http.Request, payload any )error{
	if r.Body == nil{
		return fmt.Errorf("missing request Body")
	}


	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status int , v any )error{
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)


	return json.NewEncoder(w).Encode(v)


}

func WriteError (w http.ResponseWriter, status int, err error) error{

	WriteJson(w, status, map[string]string{"error": err.Error()})

}
