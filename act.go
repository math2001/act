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
    for _, action := range actions {
        if action.Done == "0" {
            fmt.Printf("%d %s\n", action.Id, action.Message)
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

func main() {
    var filename string
    var edit bool

    flag.StringVar(&filename, "file", "./act", "A path the file to store tasks")
    flag.BoolVar(&edit, "e", false, "Edit action")

    flag.Parse()

    var actions = parseActions(filename)

    if !edit {
        renderActions(actions)
    } else {
        args := flag.Args()
        id, err := strconv.ParseInt(args[0], 10, 32)
        if err != nil { log.Fatal(err) }
        message := strings.Join(args[1:], " ")
        updateActions(filename,actions, int(id), message)
    }

}


