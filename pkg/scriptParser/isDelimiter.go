package scriptParser

func isDelimiter(char rune) bool {
    return ( char == ' ' || char == '(' || char == ')' || char == '\n' || char == '\t' )
}
