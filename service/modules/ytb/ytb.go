package ytb

import (
	"log"
	"os"
	"os/exec"
)

func DownloadVideo(url string) {
	cmd := exec.Command("powershell", "yt-dlp", url)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Fail to run command: %v", err)
	}
}
