package util

func ErrorsToStrings(errs []error) []string {
	var errors []string
	for _, err := range errs {
		if err.Error() != "" {
			errors = append(errors, err.Error())
		}
	}
	return errors
}
