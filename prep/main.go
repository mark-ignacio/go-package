package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

var (
	flagVersion string
	flagArch    string
)

func main() {
	flag.Parse()
	log.Printf("preparing v%s", flagVersion)
	if err := os.MkdirAll("package", 0o750); err != nil {
		log.Fatalf("error creating package directory: %v", err)
	}
	if err := os.MkdirAll("dist", 0o750); err != nil {
		log.Fatalf("error creating dist directory: %v", err)
	}
	prep(flagVersion, flagArch)
}

func prep(version, architecture string) {
	const packageDir = "package"
	if err := os.MkdirAll(packageDir, 0o750); err != nil {
		log.Fatalf("error creating %s: %v", packageDir, err)
	}
	// download
	uri := fmt.Sprintf("https://go.dev/dl/go%s.linux-%s.tar.gz", version, architecture)
	log.Printf("downloading %s", uri)
	r, err := http.Get(uri)
	if err != nil {
		log.Fatalf("error performing GET: %v", err)
	}
	defer r.Body.Close()
	f, err := os.CreateTemp("", fmt.Sprintf("*-%s-%s.tar.gz", version, architecture))
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	written, err := io.Copy(f, r.Body)
	if err != nil {
		log.Fatalf("error downloading: %v", err)
	}
	log.Printf("downloaded ~%dkB", written/1024)
	// extract (use a subprocess, who cares)
	cmd := exec.Command("tar", "-C", packageDir, "-xzf", f.Name())
	cmd.Stdout = os.Stdout
	cmd.Stdout = os.Stderr
	log.Printf("exec: %v", cmd.Args)
	if err := cmd.Run(); err != nil {
		log.Fatalf("error extracting: %v", err)
	}
}

func init() {
	flag.StringVar(&flagVersion, "version", "", "Go binary version")
	flag.StringVar(&flagArch, "arch", runtime.GOARCH, "defaut: runtime.GOARCH")
}
