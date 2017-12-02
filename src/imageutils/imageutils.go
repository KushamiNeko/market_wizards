package imageutils

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func ToJpeg(imageBuffer []byte, quality int) (*bytes.Buffer, error) {

	_, err := jpeg.Decode(bytes.NewBuffer(imageBuffer))
	if err != nil {

		img, err := png.Decode(bytes.NewBuffer(imageBuffer))
		if err != nil {
			return nil, err
		}

		newImg := new(bytes.Buffer)
		err = jpeg.Encode(newImg, img, &jpeg.Options{
			Quality: quality,
		})
		if err != nil {
			return nil, err
		}

		return newImg, nil

	} else {
		return bytes.NewBuffer(imageBuffer), nil
	}

	return nil, fmt.Errorf("this should never happen\n")

}

//////////////////////////////////////////////////////////////////////////////////////////////////////
