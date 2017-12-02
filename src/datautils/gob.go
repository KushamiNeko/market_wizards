package datautils

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func ToGob64(data interface{}) (string, error) {
	bytesBuffer := new(bytes.Buffer)
	gobEncoder := gob.NewEncoder(bytesBuffer)

	err := gobEncoder.Encode(data)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytesBuffer.Bytes()), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func FromGob64(b64Str string, dst interface{}) error {
	byStr, err := base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		return err
	}

	bytesBuffer := new(bytes.Buffer)
	bytesBuffer.Write(byStr)

	gobDecoder := gob.NewDecoder(bytesBuffer)
	err = gobDecoder.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
