package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"encoding/json"
)

var (
	cmd string
	env string
)

func main() {
	flag.StringVar(&cmd, "cmd", "", "help message for flagname")
	flag.StringVar(&env, "env", "", "help message for flagname")
	flag.Parse()

	getPages(env)

	exec.Command(cmd).Run()
}


func getPages(env string)  {
	raw, err := ioutil.ReadFile(env)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}


	var dat []map[string][]string

	if err := json.Unmarshal(raw, &dat); err != nil {
		panic(err)
	}

	for _, element := range dat {
		for k,v := range element {
			for _, path := range v {
				os.Setenv(k, path)
			}
		}
	}
}