package medium

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	idToEmployee := make(map[int]*Employee)
	for _, employee := range employees {
		idToEmployee[employee.Id] = employee
	}

	totalImportance := 0

	stack := []int{id}
	for len(stack) > 0 {
		curID := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		employee := idToEmployee[curID]
		totalImportance += employee.Importance

		stack = append(stack, employee.Subordinates...)
	}

	return totalImportance
}
