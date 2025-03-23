package view

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/turnerbenjamin/go_snake/entities"
	"github.com/turnerbenjamin/go_snake/game"
	"github.com/turnerbenjamin/go_snake/utilities/directions"
	"golang.org/x/term"
)

type TileDescriptionProvider interface {
	getTileDescription(byte) TileDescription
}

type TileDescription interface {
	getBackgroundColour() byte
	getChars() string
}

type Ui interface {
	Render()
}

type consoleUi struct {
	drawBuf   *bytes.Buffer
	inputReader *InputReader
	windowWidth int
}

func CreateUi() *consoleUi{
    ww, _, err := term.GetSize(int(os.Stdout.Fd()))
    if err != nil {
        panic("Unable to access terminal width")
    }
	return &consoleUi{
		drawBuf: new(bytes.Buffer),
		inputReader: newInputReader(),
		windowWidth: ww,
	}
}

func (ui *consoleUi) RenderWelcomeScreen(){
	ui.clearScreen()
	maxStringLength := getMaxStringLength(welcomeScreenTitle)
	p := ui.getPaddingToCenter(maxStringLength)
	for i, s := range welcomeScreenTitle{
		c := ui.getTitleColor(i)
    fmt.Printf("%s\033[38;5;%dm%s\033[0m\n", p, c, s)
	}

	m := "Press any key to start"
	p = ui.getPaddingToCenter(len(m))
	fmt.Println()
	fmt.Println()
	fmt.Printf("%s%s\n", p,m)

	ui.inputReader.readCharBlocking()
	ui.clearScreen()
}

func (ui *consoleUi) ShowGameOverMessage(score int) bool{
	fmt.Print(strings.Repeat("\n", 2))

	m := fmt.Sprintf("Game over: You scored %d points", score)
	p := ui.getPaddingToCenter(len(m))
	fmt.Printf("%s%s", p, m)

	fmt.Println()
	m = "Play again? [y]es or [n]o"
	p = ui.getPaddingToCenter(len(m))
	fmt.Printf("%s%s", p, m)

	for{
		char := ui.inputReader.readCharBlocking()
		if char == "y" {
			ui.clearScreen()
			return true
		}
		if char == "n"{
			ui.clearScreen()
			return false
		}
	}
}

func (ui *consoleUi) Init(){
	ui.clearScreen()
	ui.hideCursor()
}


func (ui *consoleUi) CleanUp(){
	ui.clearScreen()
	ui.showCursor()
}

func (ui *consoleUi) CheckForUserInput() (bool, directions.Direction, string){
	isInput, char := ui.inputReader.checkForCharInput()
	direction := ParseDirection(char)
	return isInput, direction, char
}


func (ui *consoleUi) RenderLevel(l game.Level, score int) {
	ui.resetBuffer()

	w := l.GetWidth()
	h := l.GetHeight()
	d := l.GetData()
	p := ui.getPaddingToCenter(w*2) // width unit is 2 chars

	fmt.Print(strings.Repeat("\n", 3))
	m := fmt.Sprintf("SCORE: %d\n", score)
	mp := ui.getPaddingToCenter(len(m))
	fmt.Printf("%s%s",mp,m)
	fmt.Print(strings.Repeat("\n", 2))

	for y := range h {
		ui.drawBuf.WriteString(p)
		for x := range w {
			tc := d[y][x]
			td := entities.GetTileDescription(tc)
			ui.drawBuf.WriteString(td)
		}
		ui.drawBuf.WriteRune('\n')
	}
	fmt.Fprint(os.Stdout, ui.drawBuf.String())
}

func (ui *consoleUi) resetBuffer(){
	ui.drawBuf.Reset()
	fmt.Print("\033[H")
}

func (ui *consoleUi) clearScreen(){
	fmt.Print("\033[H\033[2J")
}


func (ui *consoleUi) hideCursor(){
	fmt.Print("[2J\033[?25l");
}

func (ui *consoleUi) showCursor(){
	fmt.Print("\033[?25h");
}

func (ui *consoleUi)getPaddingToCenter(elWidth int) string{
	
	padding := (ui.windowWidth - elWidth) / 2
	if padding < 0{
		return ""
	}

	return strings.Repeat(" ", padding)
}

var welcomeScreenTitle = []string{
	` ______  _____        _______ __   _ _______ _     _ _______`,
	`|  ____ |     |       |______ | \  | |_____| |____/  |______`,
	`|_____| |_____| _____ ______| |  \_| |     | |    \_ |______`,
	`                                                            `,
	`           /^\/^\                                           `,
	`         _|_0|  O|                                          `,
	`\/     /~     \_/ \                                         `,
	` \____|__________/  \                                       `,
	`        \_______      \                                     `,
	`                 \     \                 \                  `,
	`                  |     |                  \                `,
	`                 /      /                    \              `,
	`                /     /                       \\            `,
	`              /      /                         \ \          `,
	`             /     /                            \  \        `,
	`           /     /             _----_            \   \      `,
	`          /     /           _-~      ~-_         |   |      `,
	`         (      (        _-~    _--_    ~-_     _/   |      `,
	`          \      ~-____-~    _-~    ~-_    ~-_-~    /       `,
	`            ~-_           _-~          ~-_       _-~        `,
	`               ~--______-~                ~-___-~           `,
}


func (ui *consoleUi) getTitleColor(i int) int{
	base := 200
	return i + base
}

func getMaxStringLength(a []string) int {
    maxLength := 0
    for _, str := range a {
        if len(str) > maxLength {
            maxLength = len(str)
        }
    }
    return maxLength
}
