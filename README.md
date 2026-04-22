# Auto Wallpaper Changer

A Windows desktop application that automatically downloads an image from a URL and sets it as your wallpaper. Runs silently in the background with no visible window.

## Features

- Automatically download image from URL
- Set as Windows wallpaper
- Runs silently (no console window)
- Simple configuration

## Configuration

Edit the `imageURL` constant in `main.go` to change the wallpaper URL:

```go
const imageURL = "https://your-image-url.jpg"
```

## Supported Image Formats

- JPG/JPEG
- PNG
- BMP
- GIF

## Building

### Prerequisites

- [Go](https://go.dev/dl/) (1.21 or higher)

### Build Commands

```bash
go mod tidy
go build -ldflags="-H=windowsgui" -o wallpaper.exe
```

The `-ldflags="-H=windowsgui"` flag removes the console window.

## Running

Just double-click `wallpaper.exe` - it will:
1. Download the image from the URL
2. Save to temp folder
3. Set as your wallpaper

No window will appear.

## License

MIT License