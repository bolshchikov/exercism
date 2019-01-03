package listops

// IntList is a list of integers
type IntList []int

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Length returns the amount of members in the list
func (list IntList) Length() (length int) {
	for range list {
		length++
	}
	return
}

// Filter returns updates list with applied function to each member
func (list IntList) Filter(fn predFunc) IntList {
	return list
}

// Foldl is folding array left
func (list IntList) Foldl(fn binFunc, arg int) int {
	return 0
}

// Foldr is folding array right
func (list IntList) Foldr(fn binFunc, arg int) int {
	return 0
}

// Map applies the given function to each arr member and returns updated array
func (list IntList) Map(fn unaryFunc) IntList {
	return list
}

// Reverse changes the order of arr members
func (list IntList) Reverse() (res IntList) {
	return list
}

// Append adds items to list
func (list IntList) Append(appendee IntList) IntList {
	var res = make([]int, list.Length()+appendee.Length())
	index := 0
	for _, v := range list {
		res[index] = v
		index++
	}
	for _, v := range appendee {
		res[index] = v
		index++
	}
	return res
}

//Concat joins two or more arrays
func (list IntList) Concat(parts []IntList) IntList {
	concatLength := 0
	for _, arr := range parts {
		concatLength += arr.Length()
	}
	var res IntList
	res = res.Append(list)
	for _, part := range parts {
		res = res.Append(part)
	}
	return res
}
