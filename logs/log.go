//todo all
package logs

import "log"

type Errors struct {
	infoLog error
	errorLog error
}

func (e Errors) ShowInfoLog() error{
	e.infoLog = e.errorLog
	log.Printf("asd")
}