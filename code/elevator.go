package code

// CalculateDestination calculates the destinations of elevator instructions
func CalculateDestination(instructions *string) int {
	floor := 0
	for _, character := range *instructions {
		if '(' == character {
			floor++
		} else if ')' == character {
			floor--
		}
	}
	return floor
}

// CalculateBasement calculates at which instruction step Santa enters the basement
func CalculateBasement(instructions *string) int {
	position := 1
	floor := 0
	for _, character := range *instructions {
		if '(' == character {
			floor++
		} else if ')' == character {
			floor--
		}
		if floor < 0 {
			break
		}
		position++
	}
	return position
}
