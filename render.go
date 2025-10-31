package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

// RenderTextToPNG renders text into a PNG at outPath.
func RenderTextToPNG(text, outPath string) error {
	const (
		fontSize = 72
		dpi      = 300
		padding  = 16
	)

	fontBytes := goregular.TTF

	ttf, err := opentype.Parse(fontBytes)
	if err != nil {
		return err
	}

	face, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}
	defer face.Close()

	lines := strings.Split(text, "\n")
	var maxWidth fixedInt
	lineHeight := face.Metrics().Height.Ceil()
	for _, ln := range lines {
		w := measureStringWidth(face, ln)
		if w > maxWidth {
			maxWidth = w
		}
	}
	if len(lines) == 0 {
		lines = []string{""}
	}

	imgW := int(maxWidth) + padding*2
	imgH := len(lines)*lineHeight + padding*2
	if imgW <= 0 {
		imgW = 1
	}
	if imgH <= 0 {
		imgH = 1
	}

	img := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: face,
	}

	x := padding
	y := padding + face.Metrics().Ascent.Ceil()
	for _, ln := range lines {
		d.Dot = fixedPoint26_6(x, y)
		d.DrawString(ln)
		y += lineHeight
	}

	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := &png.Encoder{
		CompressionLevel: png.NoCompression,
	}

	return encoder.Encode(f, img)
}

// helper type & helpers
type fixedInt int

func fixedPoint26_6(x, y int) fixed.Point26_6 {
	return fixed.Point26_6{
		X: fixed.Int26_6(x << 6),
		Y: fixed.Int26_6(y << 6),
	}
}

func measureStringWidth(face font.Face, s string) fixedInt {
	adv := font.MeasureString(face, s)
	return fixedInt(adv.Ceil())
}
