package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	invalid := false
	args := os.Args[1:]
	if len(args) >= 1 {
		switch args[0] {
		case "firefox":
			fmt.Println("updating firefox...")
			err := download("tmp/firefox-developer.tar.bz", "https://download.mozilla.org/?product=firefox-devedition-latest-ssl&os=linux64&lang=en-US")
			if err != nil {
				fmt.Println(err)
				break
			}
			_, err = exec.Command(os.Getenv("SHELL"), "-c", "sudo mv /opt/firefox tmp/firefox.bak").Output()
			if err != nil {
				fmt.Println(err)
				break
			}
			_, err = exec.Command(os.Getenv("SHELL"), "-c", "sudo find tmp -type d -exec chmod 777 {} \\;").Output()
			if err != nil {
				fmt.Println(err)
				break
			}
			_, err = exec.Command(os.Getenv("SHELL"), "-c", "sudo find tmp -type f -exec chmod 777 {} \\;").Output()
			if err != nil {
				fmt.Println(err)
				break
			}
			_, err = exec.Command(os.Getenv("SHELL"), "-c", "sudo tar -xjf tmp/firefox-developer.tar.bz -C /opt").Output()
			if err != nil {
				fmt.Println(err)
				break
			}
			err = os.Remove("tmp/firefox-developer.tar.bz")
			if err != nil {
				fmt.Println(err)
				break
			}
			err = os.RemoveAll("tmp/firefox.bak")
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("update complete")
			break
		default:
			invalid = true
		}
	}
	if invalid {
		fmt.Println("Invalid arguments.")
		fmt.Println("\nOptions:")
		fmt.Println("  firefox")
	}
}

func download(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
