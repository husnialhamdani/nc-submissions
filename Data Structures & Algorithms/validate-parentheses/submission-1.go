func isValid(s string) bool {
    
    stack := []rune{}

    bracketMap := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }

    for _, char := range s{
        if _, open := bracketMap[char]; open {
            stack = append(stack, char)
        }else{
            if len(stack)==0 || bracketMap[stack[len(stack)-1]] != char{
                return false
            }

            stack = stack[:len(stack)-1]
        }
    }


    return len(stack)==0
}
