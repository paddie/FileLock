package lock

import (
	"fmt"
	"exec"
	// "os"
)

func main() {
	// Gets the path of the foremost folder in Finder
	out, err := exec.Command("/usr/bin/osascript", "-e", "tell application \"Finder\" to get the POSIX path of (target of front window as alias)").Output()

	if err != nil {
		fmt.Printf("Couldn't determine foremost window\n\t- most likely: No window is actually open in Finder\n\t- Error: %v", err)
		return
	}
	// remove '\n' at the end of the line 
	path := string(out)[0:len(out)-1]

	fmt.Printf("Locking files in: '%v'\n", path)

	// using chflags to unlog files in directory
	out, err = exec.Command("/usr/bin/chflags", "-R", "uchg", path).Output()
	if err != nil {
		fmt.Println("Issue locking files - error: ", err)
		return
	}
	fmt.Println("Files are now locked")
}