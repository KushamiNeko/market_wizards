package datautils

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"config"
	"encoding/base64"
	"fmt"
	"imageutils"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func FileReaderExtract(content string) (*bytes.Buffer, error) {
	if !strings.Contains(content, ",") {
		return nil, fmt.Errorf("Invalid Content\n")
	}

	if !strings.Contains(content, "base64") {
		return nil, fmt.Errorf("Invalid Content\n")
	}

	s := strings.Split(content, ",")

	if len(s) < 2 {
		return nil, fmt.Errorf("Invalid Content\n")
	}

	data := s[1]

	object, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(object)

	return buffer, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func FileReaderExtractImage(content string) (*bytes.Buffer, error) {
	if !strings.Contains(content, ",") {
		return nil, fmt.Errorf("Invalid Content\n")
	}

	if !strings.Contains(content, "base64") {
		return nil, fmt.Errorf("Invalid Content\n")
	}

	s := strings.Split(content, ",")

	if len(s) < 2 {
		return nil, fmt.Errorf("Invalid Content\n")
	}

	data := s[1]

	//data := strings.Split(content, ",")[1]

	object, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	buffer, err := imageutils.ToJpeg(object, config.ImageQuality)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
