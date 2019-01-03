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
func (list IntList) Filter(fn predFunc) (res IntList) {
	if list.Length() == 0 {
		return list
	}
	for _, v := range list {
		if fn(v) {
			res = res.Append([]int{v})
		}
	}
	return
}

// Foldl is folding array left
func (list IntList) Foldl(fn binFunc, init int) int {
	if list.Length() == 0 {
		return init
	}
	res := init
	for _, v := range list {
		res = fn(res, v)
	}
	return res
}

// Foldr is folding array right
func (list IntList) Foldr(fn binFunc, init int) int {
	if list.Length() == 0 {
		return init
	}
	res := init
	reversed := list.Reverse()
	for _, v := range reversed {
		res = fn(v, res)
	}
	return res
}

// Map applies the given function to each arr member and returns updated array
func (list IntList) Map(fn unaryFunc) IntList {
	if list.Length() == 0 {
		return list
	}
	res := make([]int, list.Length())
	for i, v := range list {
		res[i] = fn(v)
	}
	return res
}

// Reverse changes the order of arr members
func (list IntList) Reverse() (res IntList) {
	if list.Length() == 0 {
		return list
	}
	for _, v := range list {
		res = IntList{v}.Append(res)
	}
	return
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
