package ytb

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type DownloadVideoResult struct {
	VideoId  string
	HlsPath  string
	FilePath string
}

type ConvertToHlsResult struct {
	metadata    DownloadVideoResult
	StoragePath string
}

func DownloadVideo(url string) (DownloadVideoResult, error) {
	id := gonanoid.MustGenerate("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 8)
	dest := filepath.Join(".", "storages", id)
	cmd := exec.Command("powershell", "yt-dlp", "-o", dest, url)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("yt-dlp comamnd failed: %v", err)
		return DownloadVideoResult{}, err
	}

	return DownloadVideoResult{VideoId: id, HlsPath: dest, FilePath: dest + ".webm"}, nil
}

func ConvertToHls(metadata DownloadVideoResult) (ConvertToHlsResult, error) {
	dest := metadata.HlsPath
	err := os.Mkdir(dest, os.ModeDir)
	if err != nil {
		log.Printf("mkdir command failed: %v", err)
	}

	cmd := exec.Command(
		"ffmpeg",
		"-i", metadata.FilePath,
		"-c:v", "h264",
		"-c:a", "aac",
		"-f", "hls",
		"-hls_time", "10",
		"-hls_list_size", "0",
		"-hls_segment_filename", filepath.Join(dest, "segment_%03d.ts"),
		filepath.Join(dest, "index.m3u8"),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("ffmpeg command failed: %v", err)
		return ConvertToHlsResult{}, err
	}

	return ConvertToHlsResult{metadata: metadata, StoragePath: dest}, nil
}
