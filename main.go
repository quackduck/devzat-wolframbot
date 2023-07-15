package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	api "github.com/quackduck/devzat/devzatapi"
)

var (
	devzatToken string
	appID       string
)

func main() {
	devzatToken = os.Getenv("DEVZAT_TOKEN")
	appID = os.Getenv("WOLFRAM_APP_ID")

	s, err := api.NewSession("devzat.hackclub.com:5556", devzatToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = s.SendMessage(api.Message{Room: "#main", From: "wolfram", Data: "I am online."})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = s.RegisterCmd("wolf", "<query>", "Ask WolframAlpha a question", func(cmdCall api.CmdCall, err error) {
		if err != nil {
			fmt.Println(err)
			s.SendMessage(api.Message{Room: cmdCall.Room, From: "wolfram", Data: "An error occurred." + err.Error()})
		}
		answer, err := getAnswer(cmdCall.Args)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = s.SendMessage(api.Message{Room: cmdCall.Room, From: "wolfram", Data: answer})
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

func getAnswer(question string) (string, error) {
	r, err := http.Get("http://api.wolframalpha.com/v1/result?output=json&i=" + url.QueryEscape(question) + "%3F&appid=" + appID)
	if err != nil {
		return "", err
	}
	response, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
