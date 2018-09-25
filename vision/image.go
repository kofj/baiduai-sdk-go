package vision

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	_ "image/jpeg" // jpeg
	"image/png"
	_ "image/png"
	"io"
	"net/url"
	"os"
)

const (
	// MaxSize 图形最长边 4096 px
	MaxSize = 4096
	// MinSize 图形最短边 15 px
	MinSize = 15
)

// Image 图形
type Image struct {
	Data string
	URL  string
	Size *Size
}

// Size 图形尺寸
type Size struct {
	Height int
	Width  int
}

// GetImage 获取图形尺寸与 base64 编码数据
func GetImage(reader io.Reader) (string, *Size, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return "", nil, err
	}
	bounds := img.Bounds()
	size := &Size{
		Width:  bounds.Dx(),
		Height: bounds.Dy(),
	}

	emptyBuff := bytes.NewBuffer(nil) //开辟一个新的空buff
	png.Encode(emptyBuff, img)        //img写入到buff

	return base64.StdEncoding.EncodeToString(emptyBuff.Bytes()), size, nil
}

// FromReader 从 io.Reader 读取图形
func FromReader(r io.Reader) (*Image, error) {
	// 检查图形尺寸
	data, size, err := GetImage(r)
	if err != nil {
		return nil, err
	}
	if size.Height > MaxSize || size.Height < MinSize || size.Width > MaxSize || size.Width < MinSize {
		return nil, errors.New("image size is invalid")
	}

	return &Image{
		Data: data,
		Size: size,
	}, nil
}

// MustFromReader Panic
func MustFromReader(reader io.Reader) *Image {
	img, err := FromReader(reader)
	if err != nil {
		panic(err)
	}
	return img
}

// FromFile 从文件读取图形
func FromFile(file string) (*Image, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	// defer f.Close()

	return FromReader(f)
}

// MustFromFile Panic
func MustFromFile(file string) *Image {
	img, err := FromFile(file)
	if err != nil {
		panic(err)
	}
	return img
}

// FromBytes 从字符串读取图形
func FromBytes(raw []byte) (*Image, error) {
	buf := bytes.NewReader(raw)
	data, size, err := GetImage(buf)
	if err != nil {
		return nil, err
	}
	if size.Height > MaxSize || size.Height < MinSize || size.Width > MaxSize || size.Width < MinSize {
		return nil, errors.New("image size is invalid")
	}

	return &Image{
		Data: data,
		Size: size,
		URL:  "",
	}, nil

}

// MustFromBytes Panic
func MustFromBytes(bts []byte) *Image {
	img, err := FromBytes(bts)
	if err != nil {
		panic(err)
	}
	return img
}

// FromURL 使用远程图形
func FromURL(link string) (*Image, error) {
	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "https" {
		return nil, errors.New("not support https scheme")
	}
	return &Image{
		Data: "",
		URL:  link,
		Size: nil,
	}, nil
}

// MustFromURL Panic
func MustFromURL(link string) *Image {
	img, err := FromURL(link)
	if err != nil {
		panic(err)
	}
	return img
}
