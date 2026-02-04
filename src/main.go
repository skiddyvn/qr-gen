package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func main() {
	// Khai báo các tham số đầu vào
	link := flag.String("link", "https://google.com", "Link muốn tạo QR")
	logoPath := flag.String("logo", "", "Đường dẫn tới file logo (PNG)")
	fgColor := flag.String("fg", "#000000", "Màu mã vạch (Hex)")
	bgColor := flag.String("bg", "#FFFFFF", "Màu nền (Hex)")
	logoSize := flag.Int("size", 20, "Kích thước logo (20 - 100)")
	outputPath := flag.String("out", "qrcode.png", "Đường dẫn file đầu ra")

	flag.Parse()

	// 1. Khởi tạo QR code
	qrc, err := qrcode.New(*link)
	if err != nil {
		log.Fatalf("Không thể khởi tạo QR: %v", err)
	}

	// 2. Thiết lập Writer với các tùy chọn tùy chỉnh
	options := []standard.ImageOption{
		standard.WithQRWidth(20),
		standard.WithFgColorRGBHex(*fgColor),
		standard.WithBgColorRGBHex(*bgColor),
	}

	// Kiểm tra nếu người dùng có truyền logo
	if *logoPath != "" {
		options = append(options,
			standard.WithLogoImageFilePNG(*logoPath),
			standard.WithLogoSizeMultiplier(int(*logoSize)),
		)
	}

	w, err := standard.New(*outputPath, options...)
	if err != nil {
		log.Fatalf("Lỗi thiết lập writer: %v", err)
	}

	// 3. Lưu file
	if err = qrc.Save(w); err != nil {
		log.Fatalf("Lỗi lưu file: %v", err)
	}

	fmt.Printf("✅ Đã tạo QR code tại: %s\n", *outputPath)
}
