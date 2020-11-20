package arrays

type ocurrence struct {
	value     int
	ocurrence int
	indexes   []int
}

func firstDuplicate(a []int) int {
	ocurrences := make(map[int]*ocurrence)
	for i, value := range a {
		if ocurrences[value] == nil {
			oc := &ocurrence{
				value:     value,
				ocurrence: 1,
				indexes:   []int{i},
			}
			ocurrences[value] = oc
		} else {
			ocurrences[value] = &ocurrence{
				value:     value,
				ocurrence: ocurrences[value].ocurrence + 1,
				indexes:   append(ocurrences[value].indexes, i),
			}
		}
	}
	firstDup := -1
	secondIndex := len(a)
	for value, ocurrence := range ocurrences {
		if ocurrence.ocurrence < 2 {
			continue
		} else {
			if ocurrence.indexes[1] < secondIndex {
				secondIndex = ocurrence.indexes[1]
				firstDup = value
			}
		}
	}
	return firstDup
}
