package view

import (
	"os"

	"github.com/turnerbenjamin/go_snake/utilities/directions"
	"golang.org/x/sys/windows"
)

var directionMap = map[string]directions.Direction{
    "k": directions.Up,
	"l": directions.Right,
	"j": directions.Down,
	"h": directions.Left,
}


type InputReader struct {
	lastChar chan string
}

func (ir *InputReader) readCharBlocking() string{
    return <- ir.lastChar 
}

func (ir *InputReader) checkForCharInput() (bool, string){
        select {
	case char := <- ir.lastChar:
        return true, char
	default:
        return false, ""
	}
}

func ParseDirection(input string) directions.Direction{
	d, exists := directionMap[input]
	if !exists {
		return directions.InvalidDirection
	}
	return d
}

func newInputReader() *InputReader{
	c := make(chan string)
	go readUserInput(c)
	return &InputReader{
		lastChar: c,
	}
}

func configureConsole() func() {
	var mode uint32
    h := windows.Handle(os.Stdin.Fd())
    windows.GetConsoleMode(h, &mode)
    newMode := mode &^ (windows.ENABLE_ECHO_INPUT | windows.ENABLE_LINE_INPUT)
    windows.SetConsoleMode(h, newMode)
    return func(){
        defer windows.SetConsoleMode(h, mode)
    }
}


func readUserInput(c chan string){    
    var resetConsole = configureConsole();
    defer resetConsole()

	var b []byte = make([]byte, 1)

    // var lastChar string
    for {
        os.Stdin.Read(b)
        currentChar := string(b)
        c <- currentChar
        // if currentChar != lastChar {
            // c <- currentChar
            // lastChar = currentChar
        // }
    }
}