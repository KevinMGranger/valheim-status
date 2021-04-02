package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"log"
	"os/exec"

	"github.com/coreos/go-systemd/v22/activation"
)

var libex = os.Getenv("libexec")

func init() {
	if libex == "" {
		libex = "/usr/libexec/valheim"
	}
}

func who(server string) (who []string, err error) {
	cmd := exec.Command(libex+"/who", server)
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(pipe)

	err = cmd.Start()
	if err != nil {
		return
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)

		if len(parts) == 2 {
			who = append(who, parts[1])
		} else {
			who = append(who, "Unknown user with ID "+parts[0])
		}
	}

	err = cmd.Wait()
	return
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	base := path.Base(req.URL.Path)

	w.Header().Add("Content-Type", "text/plain")

	if base == "/" {
		io.WriteString(w, "append the world instance name to view who's online!\n")
		return
	}

	who, err := who(base)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	if len(who) == 0 {
		io.WriteString(w, "No one's online")
	}

	for _, user := range who {
		io.WriteString(w, user+"\n")
	}
}

func main() {
	listeners, err := activation.Listeners()
	if err != nil {
		log.Fatal(err)
	}

	if len(listeners) != 1 {
		log.Fatal("Unexpected number of socket activation fds")
	}

	http.HandleFunc("/", HelloServer)
	http.Serve(listeners[0], nil)
}
