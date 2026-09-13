func searchRange(nums []int, target int) []int {
	first:=-1
    last:=-1

    left, right := 0, len(nums)-1

    for left<=right{
        mid:=left+(right-left)/2
        if nums[mid]==target{
            first=mid
            right=mid-1
        }else if nums[mid]<target{
            left=mid+1
        } else{
            right=mid-1
        }
    }

    left, right = 0, len(nums)-1
    for left<=right{
        mid:=left+(right-left)/2
        if nums[mid]==target{
            last=mid
            left=mid+1
        }else if nums[mid]<target{
            left=mid+1
        }else{
            right=mid-1
        }
    }

    return []int{first,last}
}
