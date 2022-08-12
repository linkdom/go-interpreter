package lexer

type Lexer struct {
    input           string
    position        int     // points to current char
    readPosition    int     // current reading position in input
    ch              byte    // current char under examination
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    return l
}

func readChar(l *Lexer) {
    if l.readPosition >= len(l.input){
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}
