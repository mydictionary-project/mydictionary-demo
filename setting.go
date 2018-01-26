package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/zzc-tongji/rtoa"
)

// setting
type settingStruct struct {
	AutoSaveFile struct {
		Enable             bool `json:"enable"`
		TimeIntervalSecond int  `json:"timeIntervalSecond"`
		Notification       bool `json:"notification"`
	} `json:"autoSaveFile"`
}

// read setting
func (setting *settingStruct) read() (content string, err error) {
	const TimeIntervalSecondMinimum int = 10
	var (
		buf  []byte
		path string
	)
	// convert path
	path, err = rtoa.Convert("mydictionary-local-cli.setting.json", "")
	if err != nil {
		return
	}
	// read
	buf, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(buf, setting)
	if err != nil {
		return
	}
	// check
	if setting.AutoSaveFile.TimeIntervalSecond < TimeIntervalSecondMinimum {
		setting.AutoSaveFile.TimeIntervalSecond = TimeIntervalSecondMinimum
	}
	buf, err = json.MarshalIndent(setting, "", "\t")
	content = string(buf) + "\n\n"
	return
}

// Write : write setting
func (setting *settingStruct) Write() (err error) {
	var (
		buf  []byte
		path string
	)
	// convert path
	path, err = rtoa.Convert("mydictionary-local-cli.setting.json", "")
	if err != nil {
		return
	}
	// write
	buf, err = json.MarshalIndent(setting, "", "\t")
	if err != nil {
		return
	}
	err = os.Remove(path)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, buf, 0644)
	return
}
