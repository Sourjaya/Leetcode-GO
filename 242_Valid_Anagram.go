func isAnagram(s string, t string) bool {
    if len(s)!=len(t){
        return false 
    }
    hashs:=map[int]int{}
    hasht:=map[int]int{}
    //fmt.Printf("%q",s[0])
    //hashs[s[0]]++
    //fmt.Printf("%v",hashs)
    for i:=range s{
        hashs[int(s[i])]++
        hasht[int(t[i])]++
    }

    for i,_:=range hashs{
        if hashs[i]!=hasht[i]{
            return false
        }
    }
    return true
}
