package util

import (
	"fmt"
	"io"
)

func ToString(stringer fmt.Stringer) string {
	if stringer != nil {
		return stringer.String()
	} else {
		return "<nil>"
	}
}

func Close(name string, closer io.Closer) {
	err := closer.Close()
	if err != nil {
		fmt.Println("close", name, "err:", err)
	} else {
		fmt.Println(name, "closed")
	}
}
