package util

import (
	"image"
	"image/color"
)

func GetImageColors(img image.Image) map[color.Color]struct{} {
	/* We're using empty values for this map, because go doesn't have support for sets and we're only interested in finding out whether or not a color exists in an image,
	so we can just use empty values and we'll emulate a set by using maps.*/
	colors := make(map[color.Color]struct{})
	var empty struct{}

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			colors[img.At(x, y)] = empty
		}
	}

	return colors
}
