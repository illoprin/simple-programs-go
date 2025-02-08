package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

func WriteImageToPNG(img *image.RGBA, filename string) error {
	// Define file
	file, err := os.Create(filename) // Create new file stream
	if err != nil {                  // Process errors
		return err
	}
	defer file.Close()

	var img_err error = png.Encode(file, img) // Write image data to file with png encoding
	if img_err != nil {                       // Process errors
		return err
	}
	return nil
}

var (
	box_bg_color     [3]uint8 = [3]uint8{206, 215, 224}
	box_border_color [3]uint8 = [3]uint8{115, 121, 128}
	grass_bg_color   [3]uint8 = [3]uint8{72, 148, 53}
	default_color    [3]uint8 = [3]uint8{160, 69, 196}
)

const (
	box_margin uint16 = 16 // margin in pixels
)

func gen_noise(width uint16) []float32 {
	var noise []float32 = make([]float32, width*width)

	var val float64
	for i := uint16(0); i < width*width; i++ {
		val = math.Min(float64(rand.Float32()+.5), 1.0)
		noise[i] = float32(val)
	}
	return noise
}

func getBoxColor(
	i, width, height uint32,
	grain_16 []float32,
	grain_32 []float32) [3]uint8 {

	var res [3]uint8 = [3]uint8{0, 0, 0}

	var x, y uint32 = i % uint32(width), i / uint32(width)
	var gi_16, gi_32 uint32 = x/16 + y/16, x/32 + y/32

	if x <= uint32(box_margin) ||
		x >= width-uint32(box_margin) ||
		y <= uint32(box_margin) ||
		y >= height-uint32(box_margin) {
		// Set border pixel
		res[0] = uint8(float32(box_border_color[0]) * grain_32[gi_32])
		res[1] = uint8(float32(box_border_color[1]) * grain_32[gi_32])
		res[2] = uint8(float32(box_border_color[2]) * grain_32[gi_32])
	} else {
		// Set background pixel
		res[0] = uint8(float32(box_bg_color[0]) * grain_16[gi_16])
		res[1] = uint8(float32(box_bg_color[1]) * grain_16[gi_16])
		res[2] = uint8(float32(box_bg_color[2]) * grain_16[gi_16])
	}

	return res
}

func getNoiseColor(
	i, width uint32,
	grain_16 []float32) [3]uint8 {

	var res [3]uint8 = [3]uint8{0, 0, 0}
	var x, y uint32 = i % uint32(width), i / uint32(width)
	var gi_16 uint32 = x/16 + y/16

	res[0] = uint8(float32(default_color[0]) * grain_16[gi_16])
	res[1] = uint8(float32(default_color[1]) * grain_16[gi_16])
	res[2] = uint8(float32(default_color[2]) * grain_16[gi_16])
	return res
}

func getGrassColor(i uint32, width uint16, grain_32 []float32) [3]uint8 {
	var res [3]uint8 = [3]uint8{0, 0, 0}
	var x, y uint16 = uint16(i) % width, uint16(i) / width
	var gi_32 uint16 = x/32 + y/32
	res[0] = uint8(float32(grass_bg_color[0]) * grain_32[gi_32])
	res[1] = uint8(float32(grass_bg_color[1]) * grain_32[gi_32])
	res[2] = uint8(float32(grass_bg_color[2]) * grain_32[gi_32])

	return res
}

func CreateTexture(width, height uint16, tex_type string) (*image.RGBA, error) {
	// Define new image
	var img_obj *image.RGBA = image.NewRGBA(
		image.Rect(0, 0, int(width), int(height)))

	// Generate grain textures with different sizes
	grain_16 := gen_noise(width / 16)
	grain_32 := gen_noise(width / 32)

	// Fill image pixels
	for i := uint32(0); i < uint32(width)*uint32(height); i++ {
		var res_color [3]uint8

		// Get pixel color
		switch tex_type {
		case "box":
			res_color = getBoxColor(i, uint32(width), uint32(height), grain_16, grain_32)
			break
		case "noise":
			res_color = getNoiseColor(i, uint32(width), grain_16)
			break
		case "grass":
			res_color = getGrassColor(i, width, grain_32)
			break
		default:
			return nil, fmt.Errorf("Unknown texture type")
		}

		// Set result color to pixel
		img_obj.Set(int(i%uint32(width)), int(i/uint32(width)),
			color.RGBA{
				res_color[0],
				res_color[1],
				res_color[2], 255})
	}

	return img_obj, nil
}
