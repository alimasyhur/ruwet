package utils

import (
	"fmt"
	"os"
	"strings"
)

type ProgressBar struct {
	Total int
	Width int
}

func NewProgressBar(total int) *ProgressBar {
	return &ProgressBar{
		Total: total,
		Width: 30,
	}
}

func (p *ProgressBar) Render(current int, message string) {
	percent := float64(current) / float64(p.Total)
	filled := int(percent * float64(p.Width))

	bar := strings.Repeat("█", filled) + strings.Repeat("░", p.Width-filled)

	// Buat string output dulu
	output := fmt.Sprintf("[%s] %3.0f%% (%d/%d) %s",
		bar,
		percent*100,
		current,
		p.Total,
		message,
	)

	// \r untuk ke awal baris, lalu print output
	// Gunakan spaces untuk menimpa sisa baris lama (jika ada)
	fmt.Fprintf(os.Stdout, "\r%-80s", output)
	os.Stdout.Sync()
}
