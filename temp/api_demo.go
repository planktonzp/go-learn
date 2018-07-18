package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/containous/traefik/log"
	"io/ioutil"
	"net/http"
)

type MSG struct {
	Mobiles []string `json:"mobiles"`
	Content string   `json:"content"`
}

type RESPMSG struct {
	Status string  `json:"status"`
	Msg    string `json:"msg"`
}

func main() {
	var msg MSG
	tel := "13261309023"
	if len(tel) != 0 {
		msg.Mobiles = append(msg.Mobiles, tel)
		msg.Content = "又挂啦,修不了啦"

		bytesData, _ := json.Marshal(msg)
		/*
			if err != nil {
				fmt.Println(err.Error())
				log.Error(err)
				return string("")
			}
		*/
		reader := bytes.NewReader(bytesData)
		url := "http://10.161.35.65:1821/octopus/rest/api/message/send/watchu"
		req, err := http.NewRequest("POST", url, reader)
		if err != nil {
			log.Error(err)
			fmt.Println("网址有问题，请检查")
            return
		}
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		cli := http.Client{}
		resp, err := cli.Do(req)
		if err != nil {
			log.Error(err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
			return
		}
		var respmsg RESPMSG
		e := json.Unmarshal(body, &respmsg)
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		fmt.Println(respmsg.Msg)
        return
	}
	fmt.Println(string("没写联系人，我也不知道联系谁"))
}
