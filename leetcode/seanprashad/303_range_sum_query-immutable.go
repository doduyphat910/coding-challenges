package seanprashad

type NumArray struct {
	prefixNums []int
}

func Constructor(nums []int) NumArray {
	var prefixNums = make([]int, len(nums), len(nums))
	prefixNums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixNums[i] = nums[i] + prefixNums[i-1]
	}
	return NumArray{prefixNums: prefixNums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > 0 {
		return this.prefixNums[right] - this.prefixNums[left-1]
	}
	return this.prefixNums[right]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
