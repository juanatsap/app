package main

import (
	"fmt"

	"github.com/juanatsap/go-toolkit"
)

func main() {
	fmt.Println("🔄 Generating random string...")
	var t = toolkit.Tools{
		MaxFileSize: 1024 * 1024 * 1024,
		AllowedFileTypes: []string{
			"image/jpeg",
			"image/png",
			"image/gif",
		},
	}
	randomString := t.RandomString(16)
	fmt.Println(randomString)

	fmt.Println("✅ Done.")
}
