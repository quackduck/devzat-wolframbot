package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/jwalton/gchalk"
	api "github.com/quackduck/devzat/devzatapi"
)

var (
	devzatToken = os.Getenv("DEVZAT_TOKEN")
	appID       = os.Getenv("WOLFRAM_APP_ID")
	//wolfClient  = &wolfram.Client{AppID: appID}
	name = gchalk.BrightMagenta("wolfram")
)

func main() {
	s, err := api.NewSession("devzat.hackclub.com:5556", devzatToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = s.SendMessage(api.Message{Room: "#main", From: name, Data: "I am online."})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = s.RegisterCmd("wolf", "[-v/--verbose] <query>", "Ask WolframAlpha a question", func(cmdCall api.CmdCall, err error) {
		if err != nil {
			fmt.Println(err)
			s.SendMessage(api.Message{Room: cmdCall.Room, From: name, Data: "An error occurred." + err.Error()})
		}
		if len(cmdCall.Args) == 0 {
			s.SendMessage(api.Message{Room: cmdCall.Room, From: name, Data: "You need to ask a question."})
			return
		}
		f := strings.Fields(cmdCall.Args)
		verbose := false
		for i := 0; i < len(f); i++ {
			if f[i] == "-v" || f[i] == "--verbose" {
				verbose = true
				f = append(f[:i], f[i+1:]...)
			}
		}
		answer, err := getAnswer(strings.Join(f, " "), verbose)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = s.SendMessage(api.Message{Room: cmdCall.Room, From: name, Data: answer})
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = <-s.ErrorChan
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getAnswer(question string, verbose bool) (string, error) {
	response := ""
	r, err := http.Get("http://api.wolframalpha.com/v2/query?output=json&input=" + url.QueryEscape(question) + "%3F&appid=" + appID)
	if err != nil {
		return "", err
	}
	res := &WolframAPIResult{}
	err = json.NewDecoder(r.Body).Decode(res)
	if err != nil {
		return "", err
	}
	for _, pod := range res.Result.Pods {
		if !verbose {
			if pod.Title == "Result" {
				response = gchalk.BrightYellow(pod.Subpods[0].Plaintext)
				break
			}
			if pod.Title == "Input interpretation" || pod.Title == "Input" {
				continue
			}
		}
		if len(pod.Subpods) == 0 || pod.Subpods[0].Plaintext == "" {
			continue
		}
		if pod.Title != "Result" {
			response += pod.Title + ": "
		}
		for _, subpod := range pod.Subpods {
			if subpod.Plaintext == "" {
				continue
			} else if pod.Title == "Result" {
				response += gchalk.BrightYellow(subpod.Plaintext) + "  \n"
			} else {
				response += subpod.Plaintext + "  \n"
			}
		}
		if pod.Title == "Result" {
			response += "  \n"
		}
	}
	return response, nil
}

type WolframAPIResult struct {
	Result struct {
		Pods []struct {
			Title   string `json:"title"`
			Subpods []struct {
				Title     string `json:"title"`
				Plaintext string `json:"plaintext"`
			} `json:"subpods"`
		} `json:"pods"`
	} `json:"queryresult"`
}
