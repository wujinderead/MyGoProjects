package dbtest

import (
	"fmt"
	"io"
)

func closeResource(prompt string, closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			fmt.Println("close error: ", err)
			return
		}
		fmt.Println(prompt, "closed.")
	}
}
