package utils

func Dedupe(input []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, i := range input {
		if !seen[i] {
			seen[i] = true
			result = append(result, i)
		}
	}

	return result
}
