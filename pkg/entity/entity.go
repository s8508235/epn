package entity

import "go.uber.org/dig"

type PhoneFileArgs struct {
	dig.In
	Input  string `name:"input_file"`
	Output string `name:"output_file"`
}
type CheckFileArgs struct {
	dig.In
	Input string `name:"check_file"`
}
