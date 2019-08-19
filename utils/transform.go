package utils

func Pluralize(singularForm string, pluralForm string, count int) string {
	if count == 1 {
		return singularForm
	}

	return pluralForm
}
