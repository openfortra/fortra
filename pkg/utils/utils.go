package utils

import (
	"bufio"
	"log"
	"fmt"
	"os"
	"io/ioutil"
	"runtime"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"github.com/openfortra/fortra/internal/constants"
	"strings"
)

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// Check if File / Directory Exists
/*
func exists(path string) (bool, error) {
	_, err := v.fs.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}*/

func runtimeEnv() string {
	if runtime.GOOS == "windows" {
		return ""
	}
	return constants.DefaultUserDirLinux
}

func UserConfigFile() string {
	configPath := filepath.Join(userHomeDir(), runtimeEnv())
	configFile := filepath.Join(configPath, constants.DefaultUserConfigFile)
	return configFile
}

func ConfigFileDirExists() {
	configPath := filepath.Join(userHomeDir(), runtimeEnv())
	//configFile := filepath.Join(configPath, constants.DefaultUserConfigFile)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.Mkdir(configPath, 0700)
	}
}

// CliReader handles getting input from command line interactively
func CliReader(msg string, optVal string) string {
	if optVal == "" {
	    scanner := bufio.NewScanner(os.Stdin)
	    fmt.Printf(" %s: ", msg)
		for scanner.Scan() {
			return scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
	return optVal
}

func CliQuestion(msg string) string {
	lines := []string{}
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Printf("%s (Y/N): ", msg)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		break
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	line := strings.ToLower(strings.Join(lines, ""))
	if line == "yes" {
		line = "y"
	} else if line == "no" {
		line = "n"
	}
	return line
}

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func YamlReader(content interface{}, filename string) interface{} {
	yamlContent, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(yamlContent), &content)
    if err != nil {
         log.Fatalf("error: %v", err)
    }
    return content
}

func YamlWriter(content interface{}, filename string) {
    yamlContent, err := yaml.Marshal(content)
    if err != nil {
            log.Fatalf("error: %v", err)
    }
	ioutil.WriteFile(filename, yamlContent, 0600)
}

