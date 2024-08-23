package stringx

func GetShortContent(content string, splitLen int) string {
	runes := []rune(content)
	if len(runes) <= splitLen {
		return content
	}

	return string(runes[:splitLen])
}
