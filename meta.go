package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Script struct {
	name     string
	commands []string
	args     []string
}

type Meta struct {
	scripts map[string]Script
}

func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	return string(content), err
}

func readConfigFile() string {
	fileNames := []string{"Metafile"}

	for _, fileName := range fileNames {
		_, err := os.Stat(fileName)

		if err != nil {
			continue
		}

		content, err := readFile(fileName)

		if err != nil {
			continue
		}

		return content
	}

	panic("Could not find config file.")
}

func parseConfig(content string) Meta {
	scripts := make(map[string]Script)

	var i int
	lines := strings.Split(content, "\n")

	for i != len(lines)-1 {
		fields := strings.Fields(lines[i])
		i += 1

		if len(fields) == 0 {
			continue
		}

		name := strings.TrimRight(fields[0], ":")
		args := fields[1:]

		var commands []string

		for ; len(lines[i]) > 0 && lines[i][0] == '\t'; i++ {
			command := strings.TrimLeft(lines[i], "\t")
			commands = append(commands, command)
		}

		script := Script{name: name, args: args, commands: commands}
		scripts[name] = script

	}

	return Meta{scripts: scripts}
}

func runScript(meta Meta) {
	args := os.Args[1:]
	name, args := args[0], args[1:]

	switch name {
	case "help":
		fmt.Printf("Meta - %d script(s)\n", len(meta.scripts))

		fmt.Println("\nusage: meta {script} [...args]")

		fmt.Println("\nscripts:")
		for scriptName, script := range meta.scripts {
			fmt.Printf("  %s ", scriptName)
			for _, arg := range script.args {
				fmt.Printf("{%s} ", arg)
			}
			fmt.Println()

			for _, command := range script.commands {
				fmt.Printf("    > %s\n", command)
			}
		}
	default:
		script, ok := meta.scripts[name]

		if !ok {
			panic(fmt.Sprintf("The script `%s` is not a know script.", name))
		}
		
		env := os.Environ()

		if len(script.args) != len(args) {
			panic(fmt.Sprintf(
				"The script `%s` needs %d args, but %d were passed.",
				name,
				len(script.args),
				len(args)))
		}

		for i, arg_name := range script.args {
			arg := fmt.Sprintf("%s=%s", arg_name, args[i])
			env = append(env, arg)
		}

		for _, command := range script.commands {
			cmd := exec.Command("/bin/bash", "-c", command)
			cmd.Env = env

			fmt.Printf("> %s\n", command)

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()

			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	content := readConfigFile()
	meta := parseConfig(content)
	runScript(meta)
}
