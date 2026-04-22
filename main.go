package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

const imageURL = "https://i.postimg.cc/HLjqg0mG/arknights-endfield-3840x2160-26254.jpg"

func main() {
	resp, err := http.Get(imageURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	ext := getExt(imageURL)
	if ext == "" {
		ext = ".jpg"
	}

	tmpDir := os.TempDir()
	tmpFile := filepath.Join(tmpDir, "wallpaper"+ext)

	out, err := os.Create(tmpFile)
	if err != nil {
		return
	}

	_, err = io.Copy(out, resp.Body)
	out.Close()
	if err != nil {
		return
	}

	setWallpaper(tmpFile)
}

func getExt(url string) string {
	exts := []string{".jpg", ".jpeg", ".png", ".bmp", ".gif"}
	lower := strings.ToLower(url)
	for _, ext := range exts {
		if idx := strings.LastIndex(lower, ext); idx != -1 {
			return ext
		}
	}
	return ""
}

func setWallpaper(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	pathPtr, err := syscall.UTF16PtrFromString(absPath)
	if err != nil {
		return err
	}

	user32 := syscall.MustLoadDLL("user32.dll")
	SystemParametersInfoW := user32.MustFindProc("SystemParametersInfoW")

	const SPI_SETDESKWALLPAPER = 0x0014
	const SPIF_UPDATEINIFILE = 0x01
	const SPIF_SENDCHANGE = 0x02

	SystemParametersInfoW.Call(
		SPI_SETDESKWALLPAPER,
		0,
		uintptr(unsafe.Pointer(pathPtr)),
		SPIF_UPDATEINIFILE|SPIF_SENDCHANGE,
	)

	return nil
}