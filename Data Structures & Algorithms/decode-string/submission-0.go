func decodeString(s string) string {
    stringStack := []string{}
    numberStack := []int{}
    currString := ""
    currNumber := 0

    for _, char := range s{
        if char == '['{
            stringStack = append(stringStack, currString)
            numberStack = append(numberStack, currNumber)
            currString = ""
            currNumber = 0
        } else if char == ']' {
            num := numberStack[len(numberStack)-1]
            prevString := stringStack[len(stringStack)-1]
            currString = prevString + strings.Repeat(currString, num)
            stringStack = stringStack[:len(stringStack)-1]
            numberStack = numberStack[:len(numberStack)-1]
        } else if char >= '0' && char <= '9'{
            currNumber = currNumber*10+int(char-'0')
        } else{
            currString += string(char)
        }
    }

    return currString
}
