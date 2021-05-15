package rover

import (
	"fmt"
	"io"
	"strings"
)

// onCommand A SplitFunc that filters out bad commands. It returns tokens for only valid commands
func onCommand(data []byte, atEOF bool) (advance int, token []byte, err error) {

	data = []byte(strings.ToUpper(string(data)))
	var start int

	for i := 0; i < len(data); i++ {
		if isValidCommand(data[i]) {
			return i + 1, data[start : i+1], nil
		}

		// if it is not a size for the command
		if !(data[i] >= '0' && data[i] <= '9') {
			start = i + 1
		}
	}

	// no command found yet, try with longer input if not at the end already
	if !atEOF {
		return 0, nil, nil
	}

	return 0, nil, io.EOF
}

func isValidCommand(data byte) bool {
	return data == 'F' || data == 'B' || data == 'R' || data == 'L' || data == 'X'
}

func logPosition(isStart bool, prefix string, rvr *rover) {
	if isStart {
		fmt.Printf("%s: %s -> ", prefix, rvr)
	} else {
		fmt.Printf("%s\n", rvr)
	}
}
