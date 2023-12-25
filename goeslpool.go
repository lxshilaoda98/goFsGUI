package main

import (
	"fmt"
	"github.com/fiorix/go-eventsocket/eventsocket"
)

type FsGUIInput struct {
	Command                   string
	Fshost, Fsport, Fspasswod string
}

func ConnectionFS(fsGUIInput *FsGUIInput) (client *eventsocket.Connection, err error) {
	conStr := fmt.Sprintf("%s:%s", fsGUIInput.Fshost, fsGUIInput.Fsport)
	client, err = eventsocket.Dial(conStr, fsGUIInput.Fspasswod)

	if err != nil {
		fmt.Println("ESL Connection Err: ", err)
		return
	} else {
		fmt.Println("ESL Connection Success")
	}
	return
}

func FsApi(fsGUIInput *FsGUIInput) (jsonStr string, err error) {
	client, err := ConnectionFS(fsGUIInput)
	if err != nil {
		fmt.Println("[FsApi]连接fs失败")
	} else {
		event, err := client.Send(fmt.Sprintf("api %s %s", fsGUIInput.Command, ""))
		fmt.Println("[FsApi]发送信息>>>>", fmt.Sprintf("api %s", fsGUIInput.Command))
		if err != nil {
			fmt.Println("[FsApi]运行esl comment Err..>", err)
		} else {
			jsonStr = event.Body
		}
		fmt.Println("[FsApi]相应信息<<<<<", jsonStr)
		client.Close()
	}
	return
}

func FsBgApi(fsGUIInput *FsGUIInput) (err error) {
	var jsonStr = ""
	client, err := ConnectionFS(fsGUIInput)
	if err != nil {
		fmt.Println("[FsBgApi]连接fs失败")
	} else {

		event, err := client.Send(fmt.Sprintf("api %s %s", fsGUIInput.Command, ""))
		fmt.Println("[FsApi]发送信息>>>>", fmt.Sprintf("api %s", fsGUIInput.Command))
		if err != nil {
			fmt.Println("[FsBgApi]运行esl comment Err..>", err)
		} else {
			jsonStr = event.Body
		}
		fmt.Println("[FsBgApi]相应信息<<<<<", jsonStr)
		client.Close()
	}
	return
}
