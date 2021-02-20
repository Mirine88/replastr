package replast

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Config is a struct that is able to use to make.
type Config struct {
	// replace Prefix + key + Suffix to value when call a replace function.
	KeyAndValue map[string]string
	// replace Prefix + key + Suffix to value when call a replace function.
	Prefix string
	// replace Prefix + key + Suffix to value when call a replace function.
	Suffix string
	// if a count of Prefix + key + Suffix is two or more and ReplaceOne is true,
	// replace only the first Prefix + key + Suffix to value.
	ReplaceOne bool
	// if call a replace file function(ReplaceFileWithConfig) and MakeBuildFolder is true,
	// when it is finished, make the folder that the name is BuildFolderName.
	MakeBuildFolder bool
	// if call a replace file function(ReplaceFileWithConfig) and MakeBuildFolder is true,
	// when it is finished, make the folder that the name is BuildFolderName.
	BuildFolderName string
}

// VERSION is this package version.
const VERSION = "v0"

func checkErr(err error) bool {
	return err != nil
}

// NewConfig returns a Config struct with default values.
func NewConfig() Config {
	return Config{
		Prefix: "${",
		Suffix: "}",
		ReplaceOne: false,
		MakeBuildFolder: true,
		BuildFolderName: "build",
	}
}

// Replace returns a string type value that was replaced by default config.
func Replace(str string, keyAndValue map[string]string) string {
	result := str
	for key, value := range keyAndValue {
		result = strings.ReplaceAll(result, fmt.Sprintf("${%s}", key), value)
	}
	return result
}

// ReplaceWithConfig returns a string type value that was replaced by an inputted config.
func ReplaceWithConfig(str string, config Config) string {
	keyAndValue := config.KeyAndValue
	result := str
	for key, value := range keyAndValue {
		k := fmt.Sprintf("%s%s%s", config.Prefix, key, config.Suffix)
		if config.ReplaceOne {
			result = strings.Replace(result, k, value, 1)
			continue
		}
		result = strings.ReplaceAll(result, k, value)
	}
	return result
}

// ReplaceFile replaces a data of an inputted file.
func ReplaceFile(filename string, keyAndValue map[string]string) error {
	bytes, err := ioutil.ReadFile(filename)
	if checkErr(err) {
		return err
	}

	data := Replace(string(bytes), keyAndValue)

	err = ioutil.WriteFile(filename, []byte(data), os.ModePerm)
	if checkErr(err) {
		return err
	}
	return nil
}

// ReplaceFileWithConfig replaces a data of an inputted file by an inputted config.
func ReplaceFileWithConfig(filename string, config Config) error {
	bytes, err := ioutil.ReadFile(filename)
	if checkErr(err) {
		return err
	}

	data := ReplaceWithConfig(string(bytes), config)

	if config.MakeBuildFolder {
		os.Mkdir("build", os.ModePerm)
		filename = "build/" + filename
	}

	err = ioutil.WriteFile(filename, []byte(data), os.ModePerm)
	if checkErr(err) {
		return err
	}
	return nil
}