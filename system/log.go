package system

import (
	"log"
)

//Init Init
func setLogOutput(file string) {
	if file != "" {
		outfile := &TimeWriter{
			Dir:        file,
			Compress:   true,
			ReserveDay: 30,
		}

		log.SetOutput(outfile)
		//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名
	}
}
