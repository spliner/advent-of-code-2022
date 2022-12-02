package day1

type meal struct {
	calories int
}

type elf struct {
	meals []meal
}

func (e *elf) TotalCalories() int {
	var calories int
	for _, meal := range e.meals {
		calories += meal.calories
	}

	return calories
}
