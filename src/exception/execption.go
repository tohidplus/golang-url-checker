package exception

import (
	"log"
)

func LogPrint(err error) {
	if err != nil {
		log.Println(err)
	}
}
func LogFatal(err error)  {
	if err != nil {
		log.Fatalln(err)
	}
}
func LogPanic(err error)  {
	if err != nil {
		log.Panicln(err)
	}
}