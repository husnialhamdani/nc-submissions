func threeSum(nums []int) [][]int {

    sort.Ints(nums)
    n := len(nums)
    result := [][]int{}


    for i:=0; i<n-2; i++{

        if i>0 && nums[i]==nums[i-1]{
            continue
        }

        target := -nums[i]
        l, r := i+1, n-1

        for l < r{
            sum := nums[l] + nums[r]
            if sum==target{
                result = append(result, []int{nums[i], nums[l], nums[r]})

                // skip duplicates
                for l<r && nums[l] == nums[l+1]{
                    l++
                }

                for l<r && nums[r] == nums[r-1]{
                    r--
                }
                l++
                r--
            } else if sum < target{
                l++
            } else{
                r--
            }
        }
    }

    return result
}
