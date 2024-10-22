package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {
    style := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FF0000")).
        Background(lipgloss.Color("#0000FF")).
        Bold(true).
        Italic(true).
        Width(20).
        Align(lipgloss.Center)

    fmt.Println(style.Render("Hello, World!"))

}