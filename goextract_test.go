package goextract

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)


func Test_Parse(t *testing.T) {
    
    output := `Keywords for file dmca.pdf:
mimetype - application/pdf
mimetype - audio/mpeg
format version - MPEG-2.5
creation date - 00000101000000Z
modification date - 20011017180926-03'00'
`
    metas := parse(output)

    assert.Equal(t, len(metas), 4)
    
    assert.Equal(t, metas["modification date"], []string{"20011017180926-03'00'"})
    assert.Equal(t, metas["mimetype"], []string{"application/pdf", "audio/mpeg"})
}


func Test_GetMetas(t *testing.T) {

    path, _ := os.Getwd()

    metas, _ := GetMetas(path + "/go.png")

    t.Log(metas)
}