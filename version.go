package torus

import "fmt"

var (
	Version   string
	Build     string
	BuildTime string
)

func Revision() string {
	return fmt.Sprintf("rev. %v (v%v)", Build, Version)
}
