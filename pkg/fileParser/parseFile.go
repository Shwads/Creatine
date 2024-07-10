/*
This is incredibly hacky. Should've just used and tweaked someone elses yaml parser but I wanted to do it myself. Maybes you can write a better one in a future project
but for now this will have to do.
*/

package fileParser

type Mode int8

const (
	Normal Mode = iota
	Headers
	List
	MultiLineVal
)

var parserMode Mode

