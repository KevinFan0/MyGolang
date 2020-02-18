package main

import (...)

func echo(wr http.ResponseWriter, r *http.ResponseWriter){
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil{
		wr.Writer([]byte("echo error"))
		return
	}
	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg){
		log.Println(err, "write len:", writeLen)
	}
}

func main1()  {
	http.HandleFunc("/", echo)
	err != nil{
		log.Fatal(err)
	}
}