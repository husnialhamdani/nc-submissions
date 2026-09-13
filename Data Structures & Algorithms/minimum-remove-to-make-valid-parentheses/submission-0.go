func minRemoveToMakeValid(s string) string {
	bytes := []byte(s)
    stack := []int{}

    for i, char := range bytes{
        if char == '('{
            stack = append(stack, i)
        }else if char == ')'{
            if len(stack)>0{
                stack = stack[:len(stack)-1]
            }else{
                bytes[i]='*'
            }
        }
    }

    for _, openIndex := range stack {
        bytes[openIndex]='*'
    }

    var result strings.Builder
    for _, c := range bytes{
        if c != '*'{
            result.WriteByte(c)
        }
    }

    return result.String()
}
