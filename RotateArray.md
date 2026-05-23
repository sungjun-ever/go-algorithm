# 배열 회전
1. 전체 배열을 뒤집는다
2. 회전시키는 수 k가 주어질 때, k가 배열 길이보다 크다면 k = k % len(nums)
3. 앞에서 k개까지 뒤집고
4. 나머지를 뒤집어 준다
```
func rotate(nums []int, k int) {
    if k > len(nums) {
        k = k % len(nums)
    }

    reverse(nums, 0, len(nums)-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}
```