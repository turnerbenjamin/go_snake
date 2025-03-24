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


type inputReader struct {
	input chan string
}

func (ir *inputReader) waitForInput() string{
    return <- ir.input 
}

func (ir *inputReader) checkForInput() (bool, string){
        select {
	case char := <- ir.input:
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

func newInputReader() *inputReader{
	c := make(chan string)
	go readUserInput(c)
	return &inputReader{
		input: c,
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
    for {
        os.Stdin.Read(b)
        currentChar := string(b)
        c <- currentChar
    }
}