package design

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool {
	return false
}

func (this NestedInteger) GetInteger() int {
	return 0
}
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

type NestedIterator struct {
	data []int
	pos  int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	if nestedList == nil {
		return nil
	}
	data := convertNestedList(nestedList)
	return &NestedIterator{
		data: data,
		pos:  0,
	}
}

// convertNestedList 转换 NestedInteger 数组
func convertNestedList(list []*NestedInteger) (res []int) {
	var convert func(nestedList []*NestedInteger)

	convert = func(nestedList []*NestedInteger) {
		if nestedList == nil {
			return
		}
		for _, nest := range nestedList {
			if nest.IsInteger() {
				res = append(res, nest.GetInteger())
			} else {
				convert(nest.GetList())
			}
		}
		return
	}
	convert(list)
	return
}

func (this *NestedIterator) Next() int {
	res := this.data[this.pos]
	this.pos++
	return res
}

func (this *NestedIterator) HasNext() bool {
	return this.pos < len(this.data)
}
