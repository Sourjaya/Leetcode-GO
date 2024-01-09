func containsDuplicate(nums []int) bool {
    hashset:=map[int]bool{}
    for _,i:=range nums{
        if _,ok:=hashset[i];ok{
            return true
        } else
        {
            hashset[i]=true
        }
    }
    return false
}
