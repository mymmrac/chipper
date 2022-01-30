package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mymmrac/chipper/core"
	"github.com/mymmrac/chipper/tests"
)

//go:embed config.yaml
var configData []byte

var rootCmd = &cobra.Command{
	Use:     "chipper",
	Short:   "Chipper is small tool for testing CPUs",
	Args:    cobra.NoArgs,
	Version: "v0.1.3",
	Run:     run,
}

func main() {
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(configData)); err != nil {
		fmt.Printf("Failed to read config: %v\n", err)
		os.Exit(1)
	}

	rootCmd.Flags().BoolP("simple-mode", "s", false, "Enable simple mode")
	if err := viper.BindPFlag("simple-mode", rootCmd.Flags().Lookup("simple-mode")); err != nil {
		fmt.Printf("Failed to bind flag: %v\n", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(_ *cobra.Command, _ []string) {
	var tcs tests.TestCases
	if err := viper.UnmarshalKey("test-case-list", &tcs); err != nil {
		fmt.Printf("Failed to get test cases: %v\n", err)
		os.Exit(1)
	}

	testList, err := tests.ParseTestCases(tcs)
	if err != nil {
		fmt.Printf("Failed to parce test cases: %v\n", err)
		os.Exit(1)
	}

	if len(testList) == 0 {
		fmt.Println("No test cases found")
		os.Exit(1)
	}

	progressReadInterval := viper.GetDuration("progress-read-interval")
	if progressReadInterval == 0 {
		fmt.Println("Progress read interval can't be 0")
		os.Exit(1)
	}

	if viper.GetBool("simple-mode") {
		core.ExecuteTests(testList, progressReadInterval, &simpleTerminalExecutor{})
	} else {
		bubbleExecutor := newBubbleTeaExecutor(testList, progressReadInterval)
		program := tea.NewProgram(bubbleExecutor)
		bubbleExecutor.setProgram(program)

		if err = program.Start(); err != nil {
			fmt.Printf("Failed to start bubble tea: %v", err)
			os.Exit(1)
		}
	}
}
