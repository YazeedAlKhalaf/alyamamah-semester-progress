package internal

import "fmt"

func CreateProgressBar(value float32, maxValue float32, size int) string {
	percentage := value / maxValue
	var progressBar string

	doneProgress := int(float32(size) * percentage)
	emptyProgress := size - doneProgress

	for i := 0; i < doneProgress; i++ {
		progressBar += "▓"
	}
	for i := 0; i < emptyProgress; i++ {
		progressBar += "░"
	}

	percentageText := fmt.Sprintf("%d%%%s", int(percentage*100), Ternary(percentage == 1, " 🚀", ""))
	progressBar = fmt.Sprintf("%s %s", progressBar, percentageText)

	return progressBar
}
