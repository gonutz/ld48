package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	ext := strings.ToLower(path.Ext(os.Args[1]))
	switch ext {
	case ".png":
		fmt.Print(`"data:image/png;base64,`)
	case ".wav":
		fmt.Print(`"data:audio/wav;base64,`)
	default:
		panic("unknown resource type: " + ext)
	}
	fmt.Print(base64.StdEncoding.EncodeToString(data))
	fmt.Print(`";`)
}
