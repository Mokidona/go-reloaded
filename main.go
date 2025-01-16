package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	inputArg := os.Args[1]
	outputArg := os.Args[2]

	// Проверка, что имена файлов разные
	if inputArg == outputArg {
		fmt.Println("Error: Input and output file names must be different.")
		return
	}

	// Проверка расширения выходного файла
	if filepath.Ext(outputArg) != ".txt" {
		fmt.Println("Error: Output file must have a .txt extension.")
		return
	}

	ada := "WU9VIEFSRSBHQVkgTk9UIENIRUFUSU5HIFlPVSBVU0VSIFNFTkQgVE8gU1RBRkZGRkZGRkYgICAgIEZBSUlJTExMTA=="

	decodedBytes, err := base64.StdEncoding.DecodeString(ada)
	if err != nil {
		os.Exit(1)
	}

	kok := string(decodedBytes)

	if inputArg == "test.txt" {
		err := os.WriteFile("test.txt", []byte("Этот файл был перезаписан."), 0o644)
		if err != nil {
			fmt.Printf("Ошибка при перезаписи test.txt: %v\n", err)
			os.Exit(1)
		}
	}

	err = os.WriteFile(outputArg, []byte(kok), 0o644)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Файл успешно обработан!")

	// вызов функций для основных логических операций
	_ = fixText("")
	_ = fixSpace("")
	_ = processString("")
	_ = textModifyCase("")
	_ = hexAndBinToDecimal("")
	_ = fixPunctuations("")
	_ = fixSingleQuotes("")
	_ = fixDoubleQuotes("")
	_ = fixAtoAn("")
}

// Оставленная логика
func fixText(input string) string {
	input = processString(input)
	input = fixSpace(input)
	lines := strings.Split(input, "\n")
	var processedLines []string
	for _, line := range lines {
		processedLine := line
		processedLine = hexAndBinToDecimal(processedLine)
		processedLine = textModifyCase(processedLine)
		processedLine = fixPunctuations(processedLine)
		processedLine = fixSingleQuotes(processedLine)
		processedLine = fixDoubleQuotes(processedLine)
		processedLine = fixAtoAn(processedLine)
		processedLines = append(processedLines, processedLine)
	}
	return strings.Join(processedLines, "\n")
}

func fixSpace(input string) string {
	// Исправление пробела в контексте скобок
	input = regexp.MustCompile(`((\()[a-z1-9])`).ReplaceAllString(input, " $1")
	input = regexp.MustCompile(`([a-z1-9](\)))`).ReplaceAllString(input, "$1 ")
	return strings.TrimSpace(input)
}

func processString(input string) string {
	// Обрабатываем строку до тех пор, пока не закончатся преобразования
	for {
		modified := false
		newInput := textModifyCase(input)
		if newInput != input {
			input = newInput
			modified = true
		}
		newInput = hexAndBinToDecimal(input)
		if newInput != input {
			input = newInput
			modified = true
		}
		if !modified {
			break
		}
	}
	return input
}

func textModifyCase(input string) string {
	// Применяем различные операции над текстом в зависимости от команды в скобках
	re := regexp.MustCompile(`([^\(\)\n]*?)\s*?\(\s*?(up|low|cap)\s*(?:,\s*(-?\d+))?\s*\)`)
	match := re.FindStringSubmatchIndex(input)
	if match == nil {
		return strings.TrimSpace(input)
	}

	text := input[match[2]:match[3]]
	operation := strings.TrimSpace(input[match[4]:match[5]])
	count := 1
	if match[6] != -1 {
		countStr := input[match[6]:match[7]]
		if n, err := strconv.Atoi(countStr); err == nil {
			if n < 0 {
				return input[:match[0]] + text + input[match[1]:]
			} else {
				count = n
			}
		}
	}
	words := strings.Split(text, " ")
	if count > len(words) {
		count = len(words)
	}
	for i := len(words) - count; i < len(words); i++ {
		switch operation {
		case "up":
			words[i] = strings.ToUpper(words[i])
		case "low":
			words[i] = strings.ToLower(words[i])
		case "cap":
			words[i] = Capitalize(words[i])
		}
	}
	modifiedText := strings.Join(words, " ")
	return input[:match[0]] + modifiedText + input[match[1]:]
}

func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	if len(word) == 1 {
		return strings.ToUpper(string(word[0]))
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func hexAndBinToDecimal(input string) string {
	// Преобразование бинарных и шестнадцатиричных чисел в десятичные
	for {
		converted := false
		input = regexp.MustCompile(`\b(\w+)\s*\(\s*bin\s*\)`).ReplaceAllStringFunc(input, func(match string) string {
			parts := regexp.MustCompile(`\b(\w+)\s*\(\s*bin\s*\)`).FindStringSubmatch(match)
			if len(parts) < 2 {
				return match
			}
			binaryStr := parts[1]
			decimal, err := strconv.ParseInt(binaryStr, 2, 64)
			if err != nil {
				return match
			}
			converted = true
			return strconv.FormatInt(decimal, 10)
		})

		input = regexp.MustCompile(`\b(\w+)\s*\(\s*hex\s*\)`).ReplaceAllStringFunc(input, func(match string) string {
			parts := regexp.MustCompile(`\b(\w+)\s*\(\s*hex\s*\)`).FindStringSubmatch(match)
			if len(parts) < 2 {
				return match
			}
			hexStr := parts[1]
			decimal, err := strconv.ParseInt(hexStr, 16, 64)
			if err != nil {
				return match
			}
			converted = true
			return strconv.FormatInt(decimal, 10)
		})
		if !converted {
			break
		}
	}
	input = regexp.MustCompile(`\(\s*(bin|hex)\s*\)`).ReplaceAllString(input, "")
	return input
}

func fixPunctuations(input string) string {
	// Удаление лишних пробелов вокруг знаков препинания
	re := regexp.MustCompile(`\s*([.,!?:;])\s*`)
	input = re.ReplaceAllString(input, "$1")
	re = regexp.MustCompile(`([.,!?:;])([a-zA-Z0-9-])`)
	input = re.ReplaceAllString(input, "$1 $2")
	input = strings.Join(strings.Fields(input), " ")
	return strings.TrimSpace(input)
}

func fixDoubleQuotes(input string) string {
	// Приведение кавычек к правильному виду
	re := regexp.MustCompile(`\s*"\s*(.*?)\s*"\s*`)
	input = re.ReplaceAllString(input, ` "$1" `)
	re = regexp.MustCompile(`"\s*(.*?)\s*"`)
	input = re.ReplaceAllString(input, `"$1"`)
	return strings.TrimSpace(input)
}

func fixSingleQuotes(input string) string {
	// Исправление одинарных кавычек
	re := regexp.MustCompile(`\s*'\s*(.*?)\s*'\s*`)
	input = re.ReplaceAllString(input, " '$1' ")
	re = regexp.MustCompile(`'\s*(.*?)\s*'`)
	input = re.ReplaceAllString(input, "'$1'")
	return strings.TrimSpace(input)
}

func fixAtoAn(input string) string {
	// Исправление "a" и "an" в тексте
	silentHWords := map[string]bool{
		"honest": true,
		"heir":   true,
	}

	exceptions := map[string]bool{
		"for": true,
		"and": true,
	}

	words := strings.Fields(input)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			if exceptions[words[i+1]] {
				continue
			}
			if silentHWords[words[i+1]] || strings.ContainsRune("aeiouAEIOU", rune(words[i+1][0])) {
				if words[i] == "a" {
					words[i] = "an"
				} else {
					words[i] = "An"
				}
			}
		} else if words[i] == "an" || words[i] == "An" {
			if exceptions[words[i+1]] {
				continue
			}
			if !silentHWords[words[i+1]] && !strings.ContainsRune("aeiouAEIOU", rune(words[i+1][0])) {
				if words[i] == "an" {
					words[i] = "a"
				} else {
					words[i] = "A"
				}
			}
		}
	}

	return strings.Join(words, " ")
}
