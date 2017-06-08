package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/zzc-tongji/mydictionary"
	"github.com/zzc-tongji/vocabulary4mydictionary"
)

var (
	setting     settingStruct
	quitChannel chan byte
)

func main() {
	var (
		err              error
		tm               time.Time
		information      string
		inputReader      *bufio.Reader
		vocabularyAsk    vocabulary4mydictionary.VocabularyAskStruct
		vocabularyResult mydictionary.VocabularyResultStruct
	)
	// title
	tm = time.Now()
	fmt.Printf("\n[%04d-%02d-%02d %02d:%02d:%02d]\n\nmydictionary-local-cli v1.0.0\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	// read setting
	information, err = setting.read()
	if err != nil {
		fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n%s\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), err.Error())
		fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\nQuit (press \"enter\" to continue).\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
		fmt.Scanf("%s", information)
		return
	}
	// output setting
	fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n%s\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), information)
	// initialize
	fmt.Printf("preparing data...")
	information, err = mydictionary.Initialize()
	fmt.Printf("\r")
	if err != nil {
		fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n%s\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), err.Error())
		fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\nQuit (press \"enter\" to continue).\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
		fmt.Scanf("%s", information)
		return
	}
	fmt.Printf(information)
	// check network
	fmt.Printf("checking network...")
	_, information = mydictionary.CheckNetwork()
	fmt.Printf("\r")
	fmt.Printf(information)
	// start a goroutine for automatic saving, manual saving and quitting
	quitChannel = make(chan byte, 1)
	go quitAndSave()
	// ready
	tm = time.Now()
	fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	fmt.Printf("ready\n\n")
	fmt.Printf("TIPS: press \"*\" and \"enter\" to quit at any time\n\n")
	// start
	inputReader = bufio.NewReader(os.Stdin)
	for {
		// input from stdin
		vocabularyAsk = input(inputReader)
		if vocabularyAsk.Word == "&" {
			// check network
			fmt.Printf("checking network...")
			_, information = mydictionary.CheckNetwork()
			fmt.Printf("\r")
			fmt.Printf(information)
		} else if vocabularyAsk.Word != "" {
			// query
			fmt.Printf("waiting for online paraphrase...")
			vocabularyResult, err = mydictionary.Query(vocabularyAsk)
			fmt.Printf("\r")
			if err != nil {
				tm = time.Now()
				fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n%s\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), err.Error())
				fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\nQuit (press \"enter\" to continue).\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
				fmt.Scanf("%s", information)
				return
			}
			// display
			output(vocabularyAsk, vocabularyResult)
		}
	}
}
