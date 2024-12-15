package config

import (
	"fmt"
	"time"
)

var (
	// Commit git commit hash of the HEAD in the executable application's source code in compile-time
	Commit string

	// Branch denotes the source code branch
	Branch string

	// Version latest git tag or a pseudo-version indicating git tag and commit fingerprint
	// in the executable application's source code in compile-time
	Version = "v0.0.0"

	// BuildDate reveals date when the executable application was build
	BuildDate = time.Now().Format("2006-01-02 15:04:05-07:00")
)

func PrintVars() {
	fmt.Println("Configurations")
	fmt.Printf("Commit: %s\n", Commit)
	fmt.Printf("Branch: %s\n", Branch)
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Build Date: %s\n", BuildDate)
}
