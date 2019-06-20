package diff

type Result struct {
	Action Action
	Index  int
	Item   string
}

type Action int

const (
	ActionAdded = Action(iota)
	ActionRemoved
	ActionChanged
)

func (a Action) String() string {
	switch a {
	case ActionAdded:
		return "added"
	case ActionRemoved:
		return "removed"
	case ActionChanged:
		return "changed"
	}
	return "?"
}

func Diff(A []string, B []string) []Result {
	result := []Result{}

	a := 0
	b := 0

	lenA := len(A)
	lenB := len(B)

compareLoop:
	for a < lenA && b < lenB {
		if A[a] != B[b] {
			for i := a + 1; i < lenA; i++ { // Check if A has extra lines
				if A[i] == B[b] {
					for j := a + 1; j <= i; j++ {
						result = append(result, Result{ActionAdded, j, A[j]})
					}
					a = i + 1 // Skip over in A
					b++
					continue compareLoop
				}
			}
			for i := b + 1; i < lenB; i++ { // Check if B has extra lines
				if B[i] == A[a] {
					for j := b + 1; j <= i; j++ {
						result = append(result, Result{ActionRemoved, j, B[j]})
					}
					b = i + 1 // Skip over in B
					a++
					continue compareLoop
				}
			}
			result = append(result, Result{ActionChanged, a, B[b]})
		}
		a++
		b++
	}
	for ; a < lenA; a++ {
		result = append(result, Result{ActionAdded, a, A[a]})
	}
	for ; b < lenB; b++ {
		result = append(result, Result{ActionRemoved, b, B[b]})
	}

	return result
}
