func maxArea(heights []int) int {

    l, r := 0, len(heights)-1
    result := 0

    for l < r{
        area := (r-l) * min(heights[l], heights[r])
        if area >= result{
            result = area
        }
        if heights[l] < heights[r]{
            l++
        }else {
            r--
        }
    
    }
    return result   
}
