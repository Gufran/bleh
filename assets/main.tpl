package main

import (
    "os"
    "{{.Repository}}/{{.Name}}"
)

func main() {
    {{.Name}}.GetInfo().PrettyPrint(os.Stderr)
}
