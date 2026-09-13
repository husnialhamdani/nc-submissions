func groupAnagrams(strs []string) [][]string {

    groups := make(map[[26]int][]string)

    for _, str := range strs{
        var count [26]int
        for _, char := range str{
            count[char-'a']++
        }

        groups[count] = append(groups[count], str)
    }

    var result [][]string
    for _, group := range groups{
        result = append(result, group)
    }

    return result
}
