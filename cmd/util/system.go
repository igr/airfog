package util

import (
	"fmt"
	"github.com/fatih/color"
	_nethttp "net/http"
	"os"
)

const info = "i"

func Error(msg string) {
	col := color.New(color.FgHiYellow)
	col.Println(msg)
	os.Exit(2)
}

func Exit(res *_nethttp.Response) {
	col := color.New(color.FgRed)
	col.Println(res.Status)
	os.Exit(1)
}

func PrintValue(name string, value string) {
	blue := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("%s %s %s\n", blue(info), name, yellow(value))
}

func PrintInfo(value string) {
	blue := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("%s %s\n", blue(info), value)
}

func Print(value string) {
	fmt.Println(value)
}
