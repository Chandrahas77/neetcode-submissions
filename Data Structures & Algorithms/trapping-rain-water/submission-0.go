func trap(height []int) int {
left, right := 0, len(height)-1
    maxL, maxR := height[left],height[right]
    totalWater := 0

    for left < right{
        if maxL < maxR{
            left++
            if height[left] > maxL{
                maxL = height[left]
            }else{
                totalWater += maxL - height[left] 
            }
        }else{
            right--
            if height[right] > maxR{
                maxR = height[right]
            }else{
                totalWater += maxR - height[right]
            }
        }
    }
    return totalWater
}
