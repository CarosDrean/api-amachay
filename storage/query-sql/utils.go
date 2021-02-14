package query_sql

func fieldString(fields []string) string {
	fieldString := ""
	for i, field := range fields {
		if i == 0 {
			fieldString = field
		} else {
			fieldString = fieldString + ", " + field
		}
	}
	return fieldString
}

func fieldStringInsert(fields []string) string {
	fieldString := ""
	for i, field := range fields {
		if i == 1 {
			fieldString = field
		} else if i != 0 {
			fieldString = fieldString + ", " + field
		}
	}
	return fieldString
}

func valuesString(fields []string) string {
	values := ""
	for i, field := range fields {
		if i == 1 {
			values = "@" + field
		} else if i != 0 {
			values = values + ", @" + field
		}
	}
	return values
}

func updatesString(fields []string) string {
	values := ""
	for i, field := range fields {
		if i == 1 {
			values = field + " = @" + field
		} else if i != 0 {
			values = values + ", " + field + " = @" + field
		}
	}
	return values
}
