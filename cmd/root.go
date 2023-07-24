/*
Copyright © 2023 Abhijit Wakchaure<abhijitwakchaure.2014@gmail.com>

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
	"os"

	"github.com/abhijitWakchaure/git-commit-finder/git"
	"github.com/abhijitWakchaure/git-commit-finder/scanner"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile, dir, repo string

const version = "1.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-commit-finder",
	Short: "Find the latest git commit where the files on the directory matches exactly as the git repo",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting git-commit-finder v%s\n", version)
		if dir == "" || repo == "" {
			fmt.Println("Missing required args!")
			os.Exit(1)
		}
		m := scanner.Scan(dir)
		commitIDs := git.GetCommitIDs(repo)
		fmt.Printf("Found %d commits for your repo: %s\n", len(commitIDs), repo)
		for i, commitID := range commitIDs {
			git.Checkout(repo, commitID)
			fmt.Printf("%d/%d Checked out commit ID %s\n", i+1, len(commitIDs), commitID)
			err := scanner.Compare(m, repo)
			if err == nil {
				// break
				fmt.Printf("\nFound the matching commit: %s\n", commitID)
				os.Exit(0)
			}
		}
		fmt.Printf("\nNo matching commit found :(\n")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&dir, "dir", "d", "", "Path to directory for which you need to find the commit ID")
	rootCmd.Flags().StringVarP(&repo, "gitRepo", "g", "", "Path to already clonned git repo")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".git-commit-finder" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".git-commit-finder")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
