package ascii

import (
	"github.com/nfnt/resize"

	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"reflect"
)

const asciistr = "MND8OX$7I?+=~:,..  "

// ScaleImage with the provided with
func ScaleImage(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return img, w, h
}

// Convert2Ascii convert the image to Ascii
func Convert2Ascii(img image.Image, w, h int) ([]byte, error) {
	table := []byte(asciistr)
	tableSize := uint64(len(asciistr)) - 1
	buf := new(bytes.Buffer)

	var err error
	for i := 0; i < h && err == nil; i++ {
		for j := 0; j < w && err == nil; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * tableSize / 255)
			err = buf.WriteByte(table[pos])
		}
		err = buf.WriteByte('\n')
	}

	return buf.Bytes(), err
}
