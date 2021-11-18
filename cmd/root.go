/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/bradcypert/ezzip/pkg"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ezzip",
	Short: "Zip a directory, Unzip a zip file, encryption optional",
	Long: `Zip a directory, Unzip a zip file, encryption optional
	
	zip a directory: 
		ezzip my_dir
	unzip a directory: 
		ezzip my_dir.zip
	
	encrypt and zip:
		ezzip my_dir --encrypt
		...
		use key: abc123 to decrypt
	
	unzip and decrypt:
		ezzip my_dir.zip --key=abc123`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: What causes this to error?
		encrypt, _ := cmd.Flags().GetBool("encrypt")
		key, _ := cmd.Flags().GetString("key")

		if strings.HasSuffix(args[0], "zip") {
			pkg.UnzipAssets(args[0], key)
			fmt.Println("Successfully unzipped: " + args[0])
		} else {
			key, _ := pkg.ZipAssets(args[0], encrypt)
			fmt.Println("Output: " + args[0] + ".zip")

			if key != nil {
				fmt.Println("Use key: " + *key + " to decrypt")
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("encrypt", "e", false, "Encrypt the zipped files")
	rootCmd.Flags().StringP("key", "k", "", "Decryption Key")
}
