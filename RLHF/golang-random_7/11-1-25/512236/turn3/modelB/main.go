func (msgs Messages) AnalyzeWordFrequency() map[string]int {
	wordCount := make(map[string]int)

	for _, msg := range msgs {
		// Trim any leading or trailing spaces from the message content
		trimmedMsg := strings.TrimSpace(msg.Content)

		// Exclude empty or space-only messages
		if len(trimmedMsg) > 0 {
			words := strings.Fields(trimmedMsg)
			if len(words) > 1 {
				for _, word := range words {
					wordCount[msg.User] += len(words)
				}
			}
		}
	}
	return wordCount
} 