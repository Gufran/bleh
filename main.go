package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

type Data struct {
	Root        string
	Repository  string
	Name        string
	Description string
}

var nocomm bool

func init() {
	flag.BoolVar(&nocomm, "no-commit", false, "do not commit changes after generating the scaffolding")
	flag.Parse()
}

func main() {
	data := getInfo()
	err := writeScaffold(data, nocomm)
	if err != nil {
		log.Fatalf("Failed to generate the application. (%v)", err.Error())
	}
}

func getInfo() Data {
	reader := bufio.NewReader(os.Stdin)
	root := ""
	repo := ""
	name := ""
	desc := ""

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("GOPATH was not found, please refer to https://github.com/golang/go/wiki/GOPATH")
	}

	paths := strings.Split(gopath, string(os.PathListSeparator))

	if len(paths) > 1 {
		fmt.Println("Multiple directories found in GOPATH")
		for root == "" {
			for i, p := range paths {
				fmt.Println("  [%d] %v", i+1, p)
			}
			fmt.Print("Select the number that corresponds to your preferred directory:")

			c := readStr(reader)
			n, err := strconv.Atoi(c)
			if err != nil || n < 1 || len(paths) < n-1 {
				fmt.Println("Invalid number [" + c + "] please select a valid number from list")
				continue
			}

			root = paths[n-1]
		}
	} else {
		root = paths[0]
	}

	root = root + "/src/"

	for repo == "" {
		fmt.Printf("Using %v as root, please enter path to generate the app\n", root)
		fmt.Print("(e.g. github.com/Gufran/scaff): ")

		repo = readStr(reader)
	}

	name = path.Base(repo)
	fmt.Printf("Using '%s' as application name, change name or press enter to keep '%s': ", name, name)
	n := readStr(reader)
	if n != "" {
		fmt.Printf("Using '%s' as application name\n", n)
		name = n
	}

	fmt.Print("Enter an optional description or press enter for none: ")
	d := readStr(reader)
	if d != "" {
		fmt.Println("Application description is updated")
		desc = d
	} else {
		fmt.Println("No description provided for application")
	}

	return Data{
		Root:        root,
		Repository:  repo,
		Name:        name,
		Description: desc,
	}
}

func readStr(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.Trim(fmt.Sprintf("%s", s), "\n")
}

func writeScaffold(d Data, nocommit bool) error {
	root := path.Join(d.Root, d.Repository)
	reader := bufio.NewReader(os.Stdin)
	name := strings.ToLower(d.Name)
	app := path.Join(root, name)

	cleanup := func(e error) error {
		fmt.Printf("Encountered an error. (%v)\n", e.Error())
		fmt.Println("Trying to remove any partially generated content")
		fmt.Printf("Remove '%v'? (type 'yes' to remove): ", root)

		c := readStr(reader)
		if strings.ToLower(c) != "yes" {
			fmt.Println("Not removing anything")
			return e
		}

		err := os.RemoveAll(root)
		if err != nil {
			e = fmt.Errorf("%s and %s", e, err)
		}
		return e
	}

	dests := map[string]*template.Template{}
	dests[path.Join(app, name+".go")] = loadTpl("app.tpl")
	dests[path.Join(root, ".gitignore")] = loadTpl("gitignore")
	dests[path.Join(root, ".travis.yml")] = loadTpl("travis.yml")
	dests[path.Join(root, "Makefile")] = loadTpl("Makefile")
	dests[path.Join(root, "main.go")] = loadTpl("main.tpl")

	err := os.MkdirAll(app, 0755)
	if err != nil {
		return cleanup(err)
	}

	for path, tpl := range dests {
		f, err := os.Create(path)
		if err != nil {
			return cleanup(err)
		}

		err = tpl.Execute(f, d)
		if err != nil {
			return cleanup(err)
		}
	}

	fmt.Printf("Application scaffolding generated at %v\n", root)

	erread := &bytes.Buffer{}
	ginit := exec.Command("git", "init")
	ginit.Dir = root
	ginit.Stdout = erread
	err = ginit.Run()
	if err != nil {
		fmt.Printf("Failed to initialise the git repository. %v - %v\n", erread.String(), err.Error())
		return nil
	}

	erread = &bytes.Buffer{}
	rem := gitRepo(d.Repository)
	gremote := exec.Command("git", "remote", "add", "origin", rem)
	gremote.Dir = root
	gremote.Stdout = erread
	err = gremote.Run()
	if err != nil {
		fmt.Printf("Failed to add git remote '%v' as origin. %v - %v\n", rem, erread.String(), err.Error())
		return nil
	}

	if nocommit {
		return nil
	}

	erread = &bytes.Buffer{}
	gadd := exec.Command("git", "add", root)
	gadd.Dir = root
	gadd.Stdout = erread
	err = gadd.Run()
	if err != nil {
		fmt.Printf("Failed to stage untracked files. %v %v\n", erread.String(), err.Error())
		return nil
	}

	erread = &bytes.Buffer{}
	gcomm := exec.Command("git", "commit", "-a", "-m", "Initial commit")
	gcomm.Dir = root
	gcomm.Stdout = erread
	err = gcomm.Run()
	if err != nil {
		fmt.Printf("Failed to commit changes to the VCS. %v %v\n", erread.String(), err.Error())
		return nil
	}

	return nil
}

func loadTpl(n string) *template.Template {
	d, err := Asset("assets/" + n)
	if err != nil {
		panic(fmt.Sprintf("failed to load asset '%v'. (%v)", n, err.Error()))
	}

	return template.Must(template.New("app").Parse(string(d)))
}

func gitRepo(r string) string {
	if strings.HasPrefix(r, "github.com") {
		repo, ns := path.Base(r), path.Dir(r)
		user, base := path.Base(ns), path.Dir(ns)
		return "git@" + base + ":" + user + "/" + repo
	}

	return r
}
