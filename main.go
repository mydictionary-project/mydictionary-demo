package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/zzc-tongji/mydictionary"
	"github.com/zzc-tongji/vocabulary4mydictionary"
)

const version = "v2.1.1"

var (
	setting     settingStruct
	quitChannel chan byte
)

func main() {
	var (
		err              error
		tm               time.Time
		timeString       string
		information      string
		inputReader      *bufio.Reader
		vocabularyAsk    vocabulary4mydictionary.VocabularyAskStruct
		vocabularyResult mydictionary.VocabularyResultStruct
	)
	// get time
	tm = time.Now()
	timeString = fmt.Sprintf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	// title
	fmt.Printf("\n%smydictionary-local-cli %s\n\n", timeString, version)
	// read setting
	information, err = setting.read()
	if err != nil {
		fmt.Printf(timeString)
		fmt.Printf(err.Error() + "\n\n")
		fmt.Printf(timeString)
		fmt.Printf("Quit (press \"enter\" to continue).\n\n")
		fmt.Scanf("%s", information)
		return
	}
	// output setting
	fmt.Printf(timeString + information + "\n\n")
	// initialize
	fmt.Printf("preparing data...")
	information, err = mydictionary.Initialize()
	fmt.Printf("\r")
	if err != nil {
		fmt.Printf(timeString)
		fmt.Printf(err.Error() + "\n\n")
		fmt.Printf(timeString)
		fmt.Printf("Quit (press \"enter\" to continue).\n\n")
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
	// get time
	tm = time.Now()
	timeString = fmt.Sprintf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	// ready
	fmt.Printf(timeString)
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
				// get time
				tm = time.Now()
				timeString = fmt.Sprintf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
				fmt.Printf(timeString)
				fmt.Printf(err.Error() + "\n\n")
				fmt.Printf(timeString)
				fmt.Printf("Quit (press \"enter\" to continue).\n\n")
				fmt.Scanf("%s", information)
				return
			}
			// display
			output(vocabularyAsk, vocabularyResult)
		}
	}
}
