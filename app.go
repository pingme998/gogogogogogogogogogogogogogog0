package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// 1. Download binary
	resp, err := http.Get("https://gitlab.com/developeranaz/git-hosts/-/raw/main/rclone/rclone?ref_type=heads")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create("rclone")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 2. Give permission
	err = os.Chmod("rclone", 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 3. Run binary
	cmd := exec.Command("./rclone", "run", "http")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
