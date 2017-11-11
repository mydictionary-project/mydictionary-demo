package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/zzc-tongji/mydictionary"
)

func quitAndSave() {
	var (
		controlC chan os.Signal
		ticker   *time.Ticker
	)
	controlC = make(chan os.Signal)
	signal.Notify(controlC, syscall.SIGINT)
	if setting.AutoSaveFile.Enable {
		// enable autosave
		ticker = time.NewTicker(time.Duration(setting.AutoSaveFile.TimeIntervalSecond) * time.Second)
		for {
			select {
			case <-ticker.C:
				// ticker
				if setting.AutoSaveFile.Notification {
					save()
				} else {
					mydictionary.Save()
				}
			case <-quitChannel:
				// press "*" and "enter"
				save()
				writeSetting()
				quit()
			case <-controlC:
				// press "control-c"
				if strings.Compare(runtime.GOOS, "windows") != 0 {
					fmt.Printf("\n\n")
					save()
					writeSetting()
				}
				quit()
			}
		}
	} else {
		// disable autosave
		for {
			select {
			case <-quitChannel:
				// press "*" and "enter"
				save()
				writeSetting()
				quit()
			case <-controlC:
				// press "control-c"
				if strings.Compare(runtime.GOOS, "windows") != 0 {
					fmt.Printf("\n\n")
					save()
					writeSetting()
				}
				quit()
			}
		}
	}
}

func save() {
	var information string
	_, information = mydictionary.Save()
	fmt.Printf(information)
}

func writeSetting() {
	var err error
	err = mydictionary.Setting.Write()
	if err != nil {
		fmt.Printf(err.Error() + "\n\n")
	}
	err = setting.Write()
	if err != nil {
		fmt.Printf(err.Error() + "\n\n")
	}
}

func quit() {
	var tm time.Time
	tm = time.Now()
	fmt.Printf("[%04d-%02d-%02d %02d:%02d:%02d]\n\nQuit.\n\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	os.Exit(0)
}
