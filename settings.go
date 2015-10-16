// configuration.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Settings struct {
	Url                string
	LoginData          string
	FilterSearchString string
	CheckPeriodMinutes int
}

func (settings *Settings) Load() error {
	file, err := os.Open("settings.json")
	if err != nil{
		return err
	}	
	
	defer file.Close()
	buffer, err := ioutil.ReadAll(file)
	if err != nil{
		fmt.Println(err.Error())
		return err
	}	
	
	err = json.Unmarshal(buffer, &settings)
	if err != nil{
		fmt.Println(err.Error())
		return err
	}	

	fmt.Println("Url: " + settings.Url)
	fmt.Println("LoginData: " + settings.LoginData)
	fmt.Println("FilterSearchString: " + settings.FilterSearchString)
	fmt.Println("CheckPeriodMinutes: " + strconv.Itoa(settings.CheckPeriodMinutes))
	
	return nil
}
