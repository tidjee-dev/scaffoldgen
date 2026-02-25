package tui

import "github.com/charmbracelet/lipgloss"

//
// Palette (balanced for dark terminals)
//

var (
	colorSuccess = lipgloss.Color("10")  // green
	colorInfo    = lipgloss.Color("12")  // blue
	colorWarn    = lipgloss.Color("11")  // yellow
	colorError   = lipgloss.Color("9")   // red
	colorHeader  = lipgloss.Color("6")   // cyan
	colorDim     = lipgloss.Color("242") // softer gray (was 240)
)

//
// Styles
//

var (
	headerStyle = lipgloss.NewStyle().
			Foreground(colorHeader).
			Bold(true)

	subHeaderStyle = lipgloss.NewStyle().
			Foreground(colorHeader).
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(colorSuccess).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(colorInfo)

	warnStyle = lipgloss.NewStyle().
			Foreground(colorWarn).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(colorError).
			Bold(true)

	dimStyle = lipgloss.NewStyle().
			Foreground(colorDim)
)

//
// Core Helpers
//

func Header(s string) string {
	return headerStyle.Render("\n─── " + s + " ───\n")
}

func SubHeader(s string) string {
	return subHeaderStyle.Render(s)
}

func Success(s string) string { return successStyle.Render(s) }
func Info(s string) string    { return infoStyle.Render(s) }
func Warn(s string) string    { return warnStyle.Render(s) }
func Error(s string) string   { return errorStyle.Render(s) }
func Dim(s string) string     { return dimStyle.Render(s) }

//
// Preview Labels (Improved)
//
// Rules:
// - existing → calm gray, readable
// - added → blue accent
// - ignored → dim + italic feel (simulated)
//

func Added(name string) string {
	return infoStyle.Render("+ " + name)
}

func Exists(name string) string {
	return "✓ " + name
}

func Ignored(name string) string {
	return dimStyle.Render("… " + name + " (ignored)")
}
