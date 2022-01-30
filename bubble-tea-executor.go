package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/mymmrac/chipper/core"
)

type testContext struct {
	name     string
	progress float64
	duration time.Duration
}

type bubbleTeaExecutor struct {
	tests                core.Tests
	progressReadInterval time.Duration

	program *tea.Program

	ready bool
	quit  bool
	w, h  int

	testCount         int
	testsContext      []testContext
	executionDuration time.Duration
}

func newBubbleTeaExecutor(tests core.Tests, progressReadInterval time.Duration) *bubbleTeaExecutor {
	return &bubbleTeaExecutor{
		tests:                tests,
		progressReadInterval: progressReadInterval,

		testCount:         -1,
		executionDuration: -1,
	}
}

func (b *bubbleTeaExecutor) setProgram(program *tea.Program) {
	b.program = program
}

func (b bubbleTeaExecutor) Init() tea.Cmd {
	return nil
}

type executionStart int

type executionEnd time.Duration

type testStart struct {
	name  string
	index int
}

type testProgress float64

type testEnd time.Duration

func (b bubbleTeaExecutor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+c", "q", "esc"))) {
			b.quit = true
			return b, tea.Quit
		}

	case tea.WindowSizeMsg:
		b.w = msg.Width
		b.h = msg.Height

		if !b.ready {
			go core.ExecuteTests(b.tests, b.progressReadInterval, b)
		}

		b.ready = true

	case executionStart:
		b.testCount = int(msg)

	case executionEnd:
		b.executionDuration = time.Duration(msg)
		return b, tea.Quit

	case testStart:
		b.testsContext = append(b.testsContext, testContext{
			name:     msg.name,
			progress: -1,
			duration: -1,
		})
	case testProgress:
		b.testsContext[len(b.testsContext)-1].progress = float64(msg)
	case testEnd:
		b.testsContext[len(b.testsContext)-1].duration = time.Duration(msg)
	}

	return b, nil
}

func (b bubbleTeaExecutor) View() string {
	if !b.ready {
		return "Loading...\n"
	}

	data := strings.Builder{}

	if b.testCount > 0 {
		data.WriteString(fmt.Sprintf("Tests to execute: %d\n\n", b.testCount))
	}

	for i, tc := range b.testsContext {
		data.WriteString(fmt.Sprintf("[%d/%d] %s\n", i+1, b.testCount, tc.name))

		if tc.progress >= 0 {
			data.WriteString(fmt.Sprintf("Progress: %f\n", tc.progress))
		}

		if tc.duration > 0 {
			data.WriteString(fmt.Sprintf("Done in: %s\n\n", tc.duration))
		}
	}

	if b.executionDuration > 0 {
		data.WriteString(fmt.Sprintf("Tests done in: %s\n", b.executionDuration))
	}

	if b.quit {
		data.WriteString("\n\nTerminated...\n")
	}

	return data.String()
}

func (b bubbleTeaExecutor) OnExecutionStart(count int) {
	b.program.Send(executionStart(count))
}

func (b bubbleTeaExecutor) OnExecutionEnd(duration time.Duration) {
	b.program.Send(executionEnd(duration))
}

func (b bubbleTeaExecutor) OnTestStart(name string, index int) {
	b.program.Send(testStart{
		name:  name,
		index: index,
	})
}

func (b bubbleTeaExecutor) OnTestProgress(progress float64) {
	b.program.Send(testProgress(progress))
}

func (b bubbleTeaExecutor) OnTestEnd(duration time.Duration) {
	b.program.Send(testEnd(duration))
}
