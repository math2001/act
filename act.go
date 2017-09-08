package main

import (
    "fmt"
    "flag"
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
    "bytes"
    "golang.org/x/crypto/ssh/terminal"
)

/*

Usage:

$ alias act="act -file=~/act"

$ act Go for a run
$ act Fix #12
$ act
2 Go for a run
7 Fix #12
$ act -e 7 Fix #11
$ act
2 Go for a run
7 Fix #11

$ alias todo="act -file=./act"

*/

type Action struct {
    Id int
    Done string
    Message string
}

func getLines(filename string) []string {
    file, err := os.Open(filename)
    var lines []string

    if err != nil {
        return lines
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return lines
}


func parseActions(filename string) []Action {
    lines := getLines(filename)
    var actions []Action
    for _, line := range lines {
        var infos = strings.SplitN(line, " ", 3)
        id, err := strconv.ParseInt(infos[0], 10, 32)
        if err != nil { log.Fatal(err) }
        actions = append(actions, Action{int(id), infos[1], infos[2]})
    }
    return actions
}

func renderActions(actions []Action) {
    var isTerminal bool = terminal.IsTerminal(int(os.Stdout.Fd()))
    for _, action := range actions {
        if action.Done == "0" {
            if isTerminal {
                fmt.Printf("\033[0;30m%d\033[0m %s\n", action.Id, action.Message)
            } else {
                fmt.Printf("%d %s\n", action.Id, action.Message)
            }
        }
    }
}

func updateActions(filename string, actions []Action, id int, newMessage string) {
    var buffer bytes.Buffer
    for _, action := range actions {
        if action.Id == id {
            action.Message = newMessage
        }
        buffer.WriteString(strconv.FormatInt(int64(action.Id), 10))
        buffer.WriteString(" ")
        buffer.WriteString(action.Done)
        buffer.WriteString(" ")
        buffer.WriteString(action.Message)
        buffer.WriteString("\n")
    }

    f, err := os.Create(filename)
    if err != nil { log.Fatal(err) }
    _, err = buffer.WriteTo(f)
    if err != nil { log.Fatal(err) }
}

func addAction(filename string, message string, lastid int) {
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0775)
    if err != nil { log.Fatal(err) }
    defer f.Close()
    var buffer bytes.Buffer
    buffer.WriteString(strconv.FormatInt(int64(lastid + 1), 10))
    buffer.WriteString(" 0 ")
    buffer.WriteString(message)
    buffer.WriteString("\n")
    buffer.WriteTo(f)
}

func main() {
    var filename string
    var editId int

    flag.StringVar(&filename, "file", "./act", "A path the file to store tasks")
    flag.IntVar(&editId, "e", -1, "Action id you want to edit")

    flag.Parse()

    actions := parseActions(filename)
    args := flag.Args()
    nargs := flag.NArg()

    if editId == -1 {
        if nargs > 0 {
            var id int
            if len(actions) != 0 {
                id = actions[len(actions)-1].Id
            } else {
                id = 0
            }
            addAction(filename, strings.Join(args, " "), id)
        } else {
            renderActions(actions)
        }
    } else {
        message := strings.Join(args, " ")
        updateActions(filename, actions, int(editId), message)
    }

}


