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
	"github.com/spf13/cobra"
	"github.com/tim15/wiz/cli/util"
	// "io"
	"os"
)

func resolve(pkg string) (name, resolve string) {
	return pkg, fmt.Sprintf("/Users/alexkreidler/registry/%s.tgz", pkg)
}

func installPackage(name, reference string) {
	fmt.Println(name, reference)
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		fmt.Errorf("Error opening file:", err.Error())
	}
	info, err := file.Stat()
	if err != nil {
		fmt.Errorf("Stat err", err)
	}
	fmt.Print(info.Name(), info.Size())
	// var _ io.Reader = (*os.File)(nil)
	archives.ExtractTarGz(file)
}

func installFromPackageFile() {
	fmt.Println("install from package file")
}

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [PACKAGE...]",
	Short: "Install packages specified by wiz.json or the input",
	Long: `Examples:
  wiz install deepmind/dnc
  wiz install tensorflow googlenet sklearn
  wiz install johnny
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
		if len(args) > 0 {
			for _, arg := range args {
				installPackage(resolve(arg))
			}
		} else {
			installFromPackageFile()
		}
	},
}

func init() {
	RootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
