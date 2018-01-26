package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/zzc-tongji/vocabulary4mydictionary"
)

// input
func input(inputReader *bufio.Reader) (vocabularyAsk vocabulary4mydictionary.VocabularyAskStruct) {
	var err error
	// input
	vocabularyAsk.Word, err = inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("%s\n\n", err.Error())
		quit()
		save()
	}
	vocabularyAsk.Word = strings.TrimSpace(vocabularyAsk.Word)
	// quit
	if strings.Compare(vocabularyAsk.Word, "*") == 0 {
		vocabularyAsk.Word = ""
		fmt.Printf("\n")
		quitChannel <- 0
		return
	}
	// check network
	if strings.Compare(vocabularyAsk.Word, "&") == 0 {
		vocabularyAsk.Word = "&"
		fmt.Printf("\n")
		return
	}
	// empty input
	if strings.Compare(vocabularyAsk.Word, "") == 0 {
		vocabularyAsk.Word = ""
		fmt.Printf("\n")
		return
	}
	// set word and option
	if strings.Contains(vocabularyAsk.Word, "!") || strings.Contains(vocabularyAsk.Word, "！") {
		vocabularyAsk.Advance = true
	} else {
		vocabularyAsk.Advance = false
	}
	if strings.Contains(vocabularyAsk.Word, "@") {
		vocabularyAsk.Online = true
	} else {
		vocabularyAsk.Online = false
	}
	if strings.Contains(vocabularyAsk.Word, "#") {
		vocabularyAsk.DoNotRecord = true
	} else {
		vocabularyAsk.DoNotRecord = false
	}
	vocabularyAsk.Word = strings.Replace(vocabularyAsk.Word, "!", "", -1)
	vocabularyAsk.Word = strings.Replace(vocabularyAsk.Word, "！", "", -1)
	vocabularyAsk.Word = strings.Replace(vocabularyAsk.Word, "@", "", -1)
	vocabularyAsk.Word = strings.Replace(vocabularyAsk.Word, "#", "", -1)
	vocabularyAsk.Word = strings.TrimSpace(vocabularyAsk.Word)
	fmt.Printf("\n")
	return
}

// output
func output(vocabularyAsk vocabulary4mydictionary.VocabularyAskStruct, vocabularyResult vocabulary4mydictionary.VocabularyResultStruct) {
	const (
		separator1 = "============================================================\n"
		separator2 = "------------------------------------------------------------\n"
	)
	var content string
	// word
	content = separator1
	content += convertVocabularyAsk(vocabularyAsk)
	// basic
	if vocabularyResult.Basic != nil {
		content += separator2
		content += "* BASIC\n"
		for i := 0; i < len(vocabularyResult.Basic); i++ {
			content += "\n"
			content += convertVocabularyAnswer(vocabularyResult.Basic[i])
		}
	}
	// advance
	if vocabularyResult.Advance != nil {
		content += separator2
		content += "* ADVANCE\n"
		for i := 0; i < len(vocabularyResult.Advance); i++ {
			content += "\n"
			content += convertVocabularyAnswer(vocabularyResult.Advance[i])
		}
	}
	content += separator1
	// new line
	content += "\n"
	fmt.Printf("%s", content)
}

// convert struct "VocabularyAsk" to string
func convertVocabularyAsk(vocabularyAsk vocabulary4mydictionary.VocabularyAskStruct) (result string) {
	result = vocabularyAsk.Word
	if vocabularyAsk.Advance || vocabularyAsk.Online || vocabularyAsk.DoNotRecord {
		result += " ("
		if vocabularyAsk.Advance {
			result += "advance, "
		}
		if vocabularyAsk.Online {
			result += "online, "
		}
		if vocabularyAsk.DoNotRecord {
			result += "do not record, "
		}
		result = strings.TrimSuffix(result, ", ")
		result += ")"
	}
	result += "\n"
	return
}

// convert struct "VocabularyAnswer" to string
func convertVocabularyAnswer(vocabularyAnswer vocabulary4mydictionary.VocabularyAnswerStruct) (result string) {
	result = vocabularyAnswer.Word
	if strings.Compare(result, "") == 0 {
		result = "{null}"
	}
	result += "\n"
	if vocabularyAnswer.Location.TableType == vocabulary4mydictionary.Online {
		result += fmt.Sprintf("  [%s]\n", vocabularyAnswer.SourceName)
	} else {
		result += fmt.Sprintf("  [%s: %d] (%d)\n", vocabularyAnswer.SourceName, vocabularyAnswer.SerialNumber, vocabularyAnswer.QueryCounter)
	}
	for i := 0; i < len(vocabularyAnswer.Definition); i++ {
		result += "  "
		result += vocabularyAnswer.Definition[i]
		result += "\n"
	}
	for i := 0; i < len(vocabularyAnswer.Note); i++ {
		result += "  # "
		result += vocabularyAnswer.Note[i]
		result += "\n"
	}
	if strings.Compare(vocabularyAnswer.Status, vocabulary4mydictionary.Basic) != 0 &&
		strings.Compare(vocabularyAnswer.Status, vocabulary4mydictionary.Advance) != 0 {
		result += "  {"
		result += vocabularyAnswer.Status
		result += "}\n"
	}
	return
}
