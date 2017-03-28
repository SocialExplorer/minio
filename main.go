package main

import (
	"fmt"
	"os"
	"runtime"

	version "github.com/hashicorp/go-version"
	"github.com/minio/mc/pkg/console"
	minio "github.com/minio/minio/cmd"
)

const (
	// Minio requires at least Go v1.7
	minGoVersion        = "1.7"
	goVersionConstraint = ">= " + minGoVersion
)

// Check if this binary is compiled with at least minimum Go version.
func checkGoVersion(goVersionStr string) error {
	constraint, err := version.NewConstraint(goVersionConstraint)
	if err != nil {
		return fmt.Errorf("'%s': %s", goVersionConstraint, err)
	}

	goVersion, err := version.NewVersion(goVersionStr)
	if err != nil {
		return err
	}

	if !constraint.Check(goVersion) {
		return fmt.Errorf("Minio is not compiled by Go %s.  Please recompile accordingly.",
			goVersionConstraint)
	}

	return nil
}

func main() {
	// When `go get` is used minimum Go version check is not triggered but it would have compiled it successfully.
	// However such binary will fail at runtime, hence we also check Go version at runtime.
	if err := checkGoVersion(runtime.Version()[2:]); err != nil {
		console.Fatalln("Go runtime version check failed.", err)
	}

	minio.Main(os.Args)
}
