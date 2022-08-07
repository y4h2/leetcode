/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) (result []int) {
	hash := map[int]int{} // num: index

	for i, numi := range nums {
		if j, ok := hash[target-numi]; ok {
			if i > j {
				i, j = j, i
			}
			result = append(result, i, j)
			return
		}

		hash[numi] = i
	}
	return
}

// @lc code=end

