package utils

import "log"

type errorStruct struct {
}

func (this *errorStruct) Report(err error) error {
	log.Printf(err.Error())
	return err
}

var Error = errorStruct{}
