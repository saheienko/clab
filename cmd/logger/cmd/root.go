// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/saheienko/clab/pkg/logger"
)

var (
	cfgFile string

	speed      int32
	bufferSize int32
	endpoint   string
	filePath   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "logger",
	Short: "Simple application that store numbers to a file",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := logger.New(endpoint, filePath, speed, bufferSize)
		if err != nil {
			return fmt.Errorf("create logger: %v", err)
		}

		return l.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().Int32VarP(&speed, "flow_speed", "s", 0, "speed of the fibogen (number/s)")
	rootCmd.PersistentFlags().Int32VarP(&bufferSize, "buffer_size", "b", 128, "writer's buffer size")
	rootCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "", "writer's address (stdout is used by default)")
	rootCmd.PersistentFlags().StringVarP(&filePath, "file_path", "f", "test.out", "file for storing data")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fibogen.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".logger" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".logger")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
