func twoSum(nums []int, target int) []int {
    hashmap:=map[int]int{}
    for i,num:=range nums{
        if key,ok:=hashmap[target-num];ok{
            return []int{key,i}
        }
        hashmap[num]=i
    }
    return []int{}
}
