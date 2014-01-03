package goextract

import (
    "strings"
    "regexp"
    "os/exec"
    "bytes"
)

type ExtractMetadata map[string][]string

func parse(output string) ExtractMetadata {
    metas := make(ExtractMetadata)

    split, _ := regexp.Compile("^([a-z ]*) - (.*)$")

    var results []string

    for _, line := range strings.Split(output, "\n") {

        results = split.FindStringSubmatch(line)

        if len(results) != 3 {
            continue
        }

        if  _, ok := metas[results[1]]; !ok {
            metas[results[1]] = make([]string, 0)
        } 

        metas[results[1]] = append(metas[results[1]], results[2])
    }

    return metas
}

func GetMetas(filename string) (ExtractMetadata, error) {
    path, err := exec.LookPath("extract")
    if err != nil {
        panic("you need to install `extract`")
    }

    cmd := exec.Command(path, filename)

    var out bytes.Buffer
    cmd.Stdout = &out

    err = cmd.Run()
    if err != nil {
        return nil, err
    }

    return parse(out.String()), nil
}