package main

import (
	"bytes"
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
	"os/exec"
	"strings"
)

func main() {

	goji.Get("/:words", handleSay)
	goji.Serve()

}

func handleSay(c web.C, w http.ResponseWriter, r *http.Request) {
	say(c.URLParams["words"])
	fmt.Fprintf(w, "ok")
}

func say(phrase string) {

	cmd := exec.Command("say")
	cmd.Stdin = strings.NewReader(phrase)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

}
