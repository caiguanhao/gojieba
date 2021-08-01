package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/caiguanhao/gojieba"
)

func main() {
	mode := flag.String("mode", "", "Cut, CutAll, CutForSearch (default), Tag, Tokenize, Extract, ExtractWithWeight")
	noHMM := flag.Bool("nohmm", false, "set HMM to false")
	number := flag.Int("n", 10, "number of items if mode is Extract")
	search := flag.Bool("ts", false, "set tokenize mode to search")
	userDict := flag.String("user", "", "location of user dictionary")
	flag.Parse()

	var userDictContent string
	if *userDict != "" {
		ud, err := ioutil.ReadFile(*userDict)
		if err != nil {
			log.Fatal(err)
		}
		userDictContent = string(ud)
	}

	jb := gojieba.NewJieba(userDictContent)
	defer jb.Free()

	write := func(content string) {
		content = strings.TrimSpace(content)
		var i interface{}
		switch strings.ToLower(*mode) {
		case "cut":
			i = jb.Cut(content, !*noHMM)
		case "cutall":
			i = jb.CutAll(content)
		case "tag":
			i = jb.Tag(content)
		case "tokenize":
			var mode gojieba.TokenizeMode = gojieba.DefaultMode
			if *search {
				mode = gojieba.SearchMode
			}
			i = jb.Tokenize(content, mode, !*noHMM)
		case "extract":
			i = jb.Extract(content, *number)
		case "extractwithweight":
			i = jb.ExtractWithWeight(content, *number)
		default:
			i = jb.CutForSearch(content, !*noHMM)
		}
		json.NewEncoder(os.Stdout).Encode(i)
	}

	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		content, _ := ioutil.ReadAll(os.Stdin)
		write(string(content))
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			content, _ := reader.ReadString('\n')
			write(content)
		}
	}
}
