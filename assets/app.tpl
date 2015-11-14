package {{.Name}}

import (
	"fmt"
	"io"
)

var (
	// Name of the application as exposed to a CLI client
	AppName string

	// AppVersion is the commit hash on which the build was generated
	AppVersion string

	// AppBuild is the time of build in UTC
	AppBuild string

	// AppBuildBranch is the git branch on which the build was generated
	AppBuildBranch string

	// AppGoVersion is the Go compiler version that generated the build
	AppGoVersion string

	// A short description of application to print along with help
	AppDescription = `{{.Description}}`
)

type Info struct {
	Name        string
	Description string
	Version     string
	Build       string
	BuildBranch string
	GoVersion   string
}

func GetInfo() Info {
	return Info{
		Name:        AppName,
		Description: AppDescription,
		Version:     AppVersion,
		Build:       AppBuild,
		BuildBranch: AppBuildBranch,
		GoVersion:   AppGoVersion,
	}
}

func (i Info) PrettyPrint(w io.Writer) {
	banner := "  %v\n"
	banner += "    %v\n"
	banner += "    build:       %v\n"
	banner += "    version:     %v\n"
	banner += "    branch:      %v\n"
	banner += "    go version:  %v\n\n"

	fmt.Fprintf(w, banner, i.Name, i.Description, i.Build, i.Version, i.BuildBranch, i.GoVersion)
}
