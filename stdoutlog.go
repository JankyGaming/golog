package golog

import (
	"fmt"
)

//NewLogClientStdOut returns a client for writing logs to StdOut
func NewLogClientStdOut() LogClient {
	return &stdOutLogClient{}
}

func (c *stdOutLogClient) Log(t logType, l logLevel, desc string, data map[string]interface{}) {
	newLog := buildLog(t, l, desc, data)

	fmt.Println(newLog.StdOutPrint)
	if data != nil {
		fmt.Printf("%+v", "Data: ")
		fmt.Printf("%+v\n", data)
	}
}

type stdOutLogClient struct {
}
