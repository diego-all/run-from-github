package codexgen

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type generator struct {
	projectName string
	techStack   []string
}

func Generate(projectName string, techStack []string) {
	g := generator{projectName, techStack}
	g.generate()
}

func (g generator) generate() {
	fmt.Println("Generating", g.projectName)

	err := g.initGitRepo()
	if err != nil {
		fmt.Println("Failed to setup a Git repository:", err)
	}
	fmt.Println("Generation completed.")
}

func (g generator) processFile(path string, content []byte) {
	if !g.needed(path) {
		return
	}
	for _, tech := range g.techStack {
		path = strings.Replace(path, tech+".", "", 1)
	}
	pathElements := strings.Split(path, "/")
	separator := string(os.PathSeparator)
	pathElements = append([]string{g.projectName}, pathElements...)
	_ = os.MkdirAll(
		strings.Join(pathElements[:len(pathElements)-1], separator),
		os.ModePerm,
	)
	fmt.Println("creating: " + strings.Join(pathElements, separator))
	err := ioutil.WriteFile(
		strings.Join(pathElements, separator),
		content,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (g generator) initGitRepo() error {
	fmt.Println("setting up Git repository")
	cmd := exec.Command("git", "init", "-b", "main", ".")
	cmd.Dir = g.projectName
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = g.projectName
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("git", "commit", "-m", "Initial commit from Goxygen")
	cmd.Dir = g.projectName
	return cmd.Run()
}

func (g generator) needed(path string) bool {
	if !hasTechPrefix(path) {
		return true
	}
	for _, tech := range g.techStack {
		if strings.HasPrefix(path, tech+".") {
			return true
		}
	}
	return false
}

func hasTechPrefix(path string) bool {
	for _, tech := range []string{
		"angular", "react", "vue", "mongo", "mysql", "postgres",
	} {
		if strings.HasPrefix(path, tech+".") {
			return true
		}
	}
	return false
}
