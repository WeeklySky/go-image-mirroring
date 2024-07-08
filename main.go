package main

import (
	"flag"
	"fmt"
	"image"

	"github.com/disintegration/imaging"
)

func main() {
	input := flag.String("input", "", "Input image file")
	output := flag.String("output", "", "Output image file")
	effect := flag.String("effect", "horizontal", "Mirror effect: horizontal or vertical")
	flag.Parse()

	if *input == "" || *output == "" {
		fmt.Println("Usage: go run main.go -input <input.jpg> -output <output.jpg> -effect <horizontal|vertical>")
		return
	}

	img, err := imaging.Open(*input)
	if err != nil {
		fmt.Printf("Failed to open image: %v\n", err)
		return
	}

	var mirroredImg *image.NRGBA
	switch *effect {
	case "horizontal":
		mirroredImg = imaging.FlipH(img)
	case "vertical":
		mirroredImg = imaging.FlipV(img)
	default:
		fmt.Println("Unknown effect. Use 'horizontal' or 'vertical'")
		return
	}

	err = imaging.Save(mirroredImg, *output)
	if err != nil {
		fmt.Printf("Failed to save image: %v\n", err)
		return
	}

	fmt.Printf("Image saved successfully: %s\n", *output)
}
