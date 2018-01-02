package datautils

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"io/ioutil"
	"net/http"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

type JsonResponseBody struct {
	Etags []string `json:"etags,omitempty"`

	Keys   []string    `json:"keys,omitempty"`
	Images []string    `json:"images,omitempty"`
	Items  interface{} `json:"items,omitempty"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

type JsonBodyInterface interface {
	JsonDecode(buffer []byte) error
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func JsonRequestBodyDecode(r *http.Request, decoder JsonBodyInterface) error {
	buffer, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	err = decoder.JsonDecode(buffer)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

//func JsonRequestBodyToStruct(r *http.Request, structInterface interface{}) error {
//buffer, err := ioutil.ReadAll(r.Body)
//defer r.Body.Close()
//if err != nil {
//return err
//}

//err = json.Unmarshal(buffer, structInterface)
//if err != nil {
//return err
//}

//return nil
//}

////////////////////////////////////////////////////////////////////////////////////////////////////////

//func JsonRequestBodyToMap(r *http.Request) (map[string]interface{}, error) {
//buffer, err := ioutil.ReadAll(r.Body)
//defer r.Body.Close()
//if err != nil {
//return nil, err
//}

//var entityMap map[string]interface{}

//if err := json.Unmarshal(buffer, &entityMap); err != nil {
//return nil, err
//}

//return entityMap, nil
//}

////////////////////////////////////////////////////////////////////////////////////////////////////////

//func JsonRequestBodyToMapSlice(r *http.Request) ([]map[string]interface{}, error) {
//buffer, err := ioutil.ReadAll(r.Body)
//defer r.Body.Close()
//if err != nil {
//return nil, err
//}

//var entityMaps []map[string]interface{}

//if err := json.Unmarshal(buffer, &entityMaps); err != nil {
//return nil, err
//}

//return entityMaps, nil
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////

//func JsonInterfaceSliceToMapSlice(data []interface{}) []map[string]interface{} {
//dataMap := make([]map[string]interface{}, len(data))

//for i, v := range data {
//dataMap[i] = v.(map[string]interface{})
//}

//return dataMap
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////
