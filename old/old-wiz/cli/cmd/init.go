// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"bufio"
	"github.com/golang/protobuf/jsonpb"
	"github.com/spf13/cobra"
	"github.com/tim15/wiz/api/pkg"
	// "io"
	"os"
	"path"
	"strings"
)

// read is a utility that reads a value from stdin with a prompt and default value
func read(prompt, defaultValue string) (out string) {
	reader := bufio.NewReader(os.Stdin)
	origVal := defaultValue
	if defaultValue != "" {
		defaultValue = "(" + defaultValue + ") "
	}
	fmt.Print(prompt + ": " + defaultValue)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	text = strings.Replace(text, "\n", "", 1)
	if text == "" {
		return origVal
	}
	return text
}

/*
mu is a utility that allows getting one return value from a function as an expression.
It returns a slice of interfaces.
To use it, cast the interface to your type.

Example:
  func Getwd() (dir string, err error)
  mu(os.Getwd())[0].(string) // evaluates to string
*/
func mu(a ...interface{}) []interface{} {
	return a
}

func setupRepos(pb *pkg.Package) {
	if _, err := os.Stat(".git"); !os.IsNotExist(err) {
		// TODO: parse .git/config (INI file) and get remote url
		pb.Repository = read("git repository", "")
	}
	// TODO: add more SVNs
}

func init() {
	// TODO: add tests
	longDesc := `A utility that will walk you through creating a wiz.json file.
It will only ask for basic fields and will try to guess sensible defaults.
Use ` + "`wiz help json`" + ` for documentation on the fields. `
	// initCmd represents the init command
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Create a wiz.json file",
		Long:  longDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(longDesc)
			k := strings.Split(read("keywords", ""), ",")
			fmt.Println(k, len(k))
			pb := &pkg.Package{
				Name:        read("package name", path.Base(mu(os.Getwd())[0].(string))),
				Version:     read("version", "1.0.0"),
				Type:        pkg.Package_PackageType(pkg.Package_PackageType_value[strings.ToUpper(read("package type", "model"))]),
				Description: read("description", ""),
				Keywords:    k,
				Author:      read("author", ""),
				License:     read("license", "MIT"),
			}
			setupRepos(pb)
			marshaler := jsonpb.Marshaler{
				Indent: "    ",
			}
			json, err := marshaler.MarshalToString(pb)
			if err != nil {
				fmt.Println("Error marshalling JSON")
				return
			}
			json += "\n"
			fileLocation := mu(os.Getwd())[0].(string) + "/wiz.json"
			fmt.Println("About to write to " + fileLocation)
			fmt.Println()
			fmt.Println(json)
			ok := read("Is this ok? (yes) ", "")
			if ok != "" {
				return
			}
			file, err := os.Create(fileLocation)
			if err != nil {
				fmt.Println(err.Error())
			}
			_, err = file.WriteString(json)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	RootCmd.AddCommand(initCmd)

	// TODO: Add flags

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
