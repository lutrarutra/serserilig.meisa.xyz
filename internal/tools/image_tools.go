package tools

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/image/draw"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// readImage reads a image file from disk. We're assuming the file will be png
// format.
func readImage(name string) (image.Image, error) {
	fd, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	// image.Decode requires that you import the right image package. We've
	// imported "image/png", so Decode will work for png files. If we needed to
	// decode jpeg files then we would need to import "image/jpeg".
	//
	// Ignored return value is image format name.

	img, _, err := image.Decode(fd)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// writeImage writes an Image back to the disk.
func writeImage(img image.Image, name string) error {
	fd, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fd.Close()

	if name[len(name)-4:] == ".png" {
		return png.Encode(fd, img)
	} else {
		return jpeg.Encode(fd, img, &jpeg.Options{75})
	}
}

// cropImage takes an image and crops it to the specified rectangle.
func cropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(crop), nil
}

func scaleImage(src_path, dst_path string) error {
	src, err := readImage(src_path)
	if err != nil {
		return err
	}

	dst, _ := os.Create(dst_path)
	temp := image.NewRGBA(image.Rect(0, 0, 256, 256))

	draw.BiLinear.Scale(temp, temp.Rect, src, src.Bounds(), draw.Over, nil)

	if dst_path[len(dst_path)-4:] == ".png" {
		return png.Encode(dst, temp)
	} else {
		return jpeg.Encode(dst, temp, &jpeg.Options{75})
	}
	return nil
}

func SaveProfilePicture(crop_x, crop_y, crop_w, crop_h int, img_w, img_h float64, src_path, file_extension string) (string, error) {
	dst_path := "static/images/profile_pics/" + RandString(8) + file_extension

	img, err := readImage(src_path)
	if err != nil {
		return "", err
	}

	bounds := img.Bounds()
	// log.Println("Hi")

	w := bounds.Dx()
	h := bounds.Dy()

	// log.Println(w)
	// log.Println(h)

	// log.Println(crop_h)
	// log.Println(crop_w)

	scale_w := float64(w) / img_w
	scale_h := float64(h) / img_h

	// log.Println(scale_w)
	// log.Println(scale_h)

	crop_x = int(float64(crop_x) * scale_w)
	crop_y = int(float64(crop_y) * scale_h)
	crop_w = int(float64(crop_w) * scale_w)
	crop_h = int(float64(crop_h) * scale_h)

	img, err = cropImage(img, image.Rect(crop_x, crop_y, crop_x+crop_w, crop_y+crop_h))
	if err != nil {
		log.Println(err)
		return "", err
	}

	if crop_w > 256 || crop_h > 256 {
		log.Println("##########################################")
		temp_path := "static/images/temp/temp-" + RandString(8) + file_extension
		err = writeImage(img, temp_path)
		if err != nil {
			return "", err
		}
		err = scaleImage(temp_path, dst_path)
		if err != nil {
			return "", err
		}
		os.Remove(temp_path)
	} else {
		err = writeImage(img, dst_path)
	}
	if err != nil {
		return "", err
	}

	return "/" + dst_path, nil
}

// func SaveProfilePicture(crop_x, crop_y, crop_w, crop_h int, src_path string) string {
// 	input, _ := os.Open(src_path)
// 	defer input.Close()

// 	name := "static/images/profile_pics/" + RandString(8) + ".png"
// 	output, _ := os.Create(name)
// 	defer output.Close()
// 	// Decode the image (from PNG to image.Image):
// 	src, _ := png.Decode(input)

// 	// Set the expected size that you want:
// 	dst := image.NewRGBA(image.Rect(0, 0, 256, 256))

// 	// Resize:
// 	draw.BiLinear.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

// 	// Encode to `output`:
// 	png.Encode(output, dst)
// 	return name
// }
