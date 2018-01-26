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
	tm          time.Time
	setting     settingStruct
	quitChannel chan byte
)

func main() {
	var (
		err              error
		success          bool
		information      string
		inputReader      *bufio.Reader
		vocabularyAsk    vocabulary4mydictionary.VocabularyAskStruct
		vocabularyResult vocabulary4mydictionary.VocabularyResultStruct
	)
	// title
	tm = time.Now()
	fmt.Printf("\n[%04d-%02d-%02d %02d:%02d:%02d]\n\nmydictionary-local-cli\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	// read setting
	information, err = setting.read()
	if err != nil {
		fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
		fmt.Printf(err.Error() + "\n\n")
		fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
		fmt.Printf("Quit (press \"enter\" to continue).\n\n")
		fmt.Scanf("%s", information)
		return
	}
	fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n%s", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), information)
	// initialize
	fmt.Printf("preparing data...")
	success, information = mydictionary.Initialize()
	fmt.Printf("\r")
	tm = time.Now()
	fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\nmydictionary\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n%s", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), information)
	if success == false {
		fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
		fmt.Printf("Quit (press \"enter\" to continue).\n\n")
		fmt.Scanf("%s", information)
		return
	}
	// check network
	fmt.Printf("checking network...")
	_, information = mydictionary.CheckNetwork()
	fmt.Printf("\r")
	tm = time.Now()
	fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\n%s", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), information)
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
			success, vocabularyResult = mydictionary.Query(vocabularyAsk)
			fmt.Printf("\r")
			if success == false {
				// get time
				tm = time.Now()
				fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\nMYDICTIONARY has not been initialized.\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
				fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\nQuit (press \"enter\" to continue).\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
				fmt.Scanf("%s", information)
				return
			}
			// display
			output(vocabularyAsk, vocabularyResult)
		}
	}
}
