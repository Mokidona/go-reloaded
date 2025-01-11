package main

import (
	"fmt"
	"os"
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

	// Чтение входного файла
	content, err := os.ReadFile(inputArg)
	if err != nil {
		fmt.Printf("Error reading arg %v\n", err)
		os.Exit(1)
	}

	// Обработка текста
	replacedContent := fixText(string(content))

	// Запись в выходной файл
	err = os.WriteFile(outputArg, []byte(replacedContent), 0o644)
	if err != nil {
		fmt.Printf("Error writing arg %v\n", err)
		os.Exit(1)
	}
}

// fixText обрабатывает входной текст через серию преобразований
func fixText(input string) string {
	input = processString(input)
	input = fixSpace(input)
	lines := strings.Split(input, "\n")
	var processedLines []string

	// Обработка каждой строки текста
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

// fixSpace корректирует пробелы вокруг скобок
func fixSpace(input string) string {
	input = regexp.MustCompile(`((\()[a-z1-9])`).ReplaceAllString(input, " $1")
	input = regexp.MustCompile(`([a-z1-9](\)))`).ReplaceAllString(input, "$1 ")
	return strings.TrimSpace(input)
}

// processString выполняет несколько итеративных преобразований текста
func processString(input string) string {
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

// textModifyCase обрабатывает команды изменения регистра текста (up, low, cap)
func textModifyCase(input string) string {
	re := regexp.MustCompile(`([^\(\)\n]*?)\s*?\(\s*?(\s*up|low|cap\s*)(?:,\s*(-?\d+))?\s*\)`)
	match := re.FindStringSubmatchIndex(input)
	if match == nil {
		return strings.TrimSpace(input)
	}

	textStart := match[2]
	textEnd := match[3]
	opStart := match[4]
	opEnd := match[5]
	countStart := match[6]
	countEnd := match[7]

	text := input[textStart:textEnd]
	operation := strings.TrimSpace(input[opStart:opEnd])
	count := 1

	// Извлечение параметра count, если он присутствует
	if countStart != -1 {
		countStr := input[countStart:countEnd]
		if n, err := strconv.Atoi(countStr); err == nil {
			if n < 0 {
				newInput := input[:match[0]] + text + input[match[1]:]
				return textModifyCase(newInput)
			} else {
				count = n
			}
		}
	}

	words := strings.Split(text, " ")
	if count > len(words) {
		count = len(words)
	}
	startWord := len(words) - count
	if startWord < 0 {
		startWord = 0
	}

	// Изменение регистра слов в зависимости от операции
	for i := startWord; i < len(words); i++ {
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
	newInput := input[:match[0]] + modifiedText + input[match[1]:]

	return textModifyCase(newInput)
}

// Capitalize делает первую букву слова заглавной
func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	if len(word) == 1 {
		return strings.ToUpper(string(word[0]))
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

// hexAndBinToDecimal преобразует числа в двоичном и шестнадцатеричном формате в десятичный
func hexAndBinToDecimal(input string) string {
	for {
		converted := false
		binPattern := regexp.MustCompile(`\b(\w+)\s*\(\s*bin\s*\)`)
		input = binPattern.ReplaceAllStringFunc(input, func(match string) string {
			parts := binPattern.FindStringSubmatch(match)
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
		hexPattern := regexp.MustCompile(`\b(\w+)\s*\(\s*hex\s*\)`)
		input = hexPattern.ReplaceAllStringFunc(input, func(match string) string {
			parts := hexPattern.FindStringSubmatch(match)
			if len(parts) < 2 {
				return match
			}
			hexStr := parts[1]
			decimal, err := strconv.ParseInt(hexStr, 16, 64)
			if err != nil || hexStr == "" {
				return match
			}
			converted = true
			return strconv.FormatInt(decimal, 10)
		})
		if !converted {
			break
		}
	}

	leftoverPattern := regexp.MustCompile(`\(\s*(bin|hex)\s*\)`)
	input = leftoverPattern.ReplaceAllString(input, "")
	return input
}

// fixPunctuations корректирует пробелы вокруг знаков препинания
func fixPunctuations(input string) string {
	re := regexp.MustCompile(`\s*([.,!?:;])\s*`)
	input = re.ReplaceAllString(input, "$1")
	re = regexp.MustCompile(`([.,!?:;])([a-zA-Z0-9-])`)
	input = re.ReplaceAllString(input, "$1 $2")
	input = strings.Join(strings.Fields(input), " ")
	return strings.TrimSpace(input)
}

// fixDoubleQuotes корректирует пробелы внутри двойных кавычек
func fixDoubleQuotes(input string) string {
	re := regexp.MustCompile(`"\s*(.*?)\s*"`)
	input = re.ReplaceAllString(input, `"$1"`)
	re = regexp.MustCompile(`(["\s]+)|([\s+"])`)
	input = re.ReplaceAllString(input, "$1$2")
	return strings.TrimSpace(input)
}

// fixSingleQuotes корректирует пробелы внутри одиночных кавычек
func fixSingleQuotes(input string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	input = re.ReplaceAllString(input, "'$1'")
	return strings.TrimSpace(input)
}

// fixAtoAn корректирует "a" и "an" перед словами
func fixAtoAn(input string) string {
	silentHWords := map[string]bool{
		"honest":    true,
		"heir":      true,
		"honorific": true,
		"honor":     true,
		"herb":      true,
		"hotel":     true,
		"hour":      true,
		"homage":    true,
	}

	words := strings.Fields(input)
	for i := 0; i < len(words)-1; i++ {
		if len(words[i+1]) > 1 {
			if words[i] == "a" || words[i] == "A" {
				if silentHWords[words[i+1]] || strings.ContainsRune("aeiouAEIOU", rune(words[i+1][0])) {
					if words[i] == "a" {
						words[i] = "an"
					} else {
						words[i] = "An"
					}
				}
			} else if words[i] == "an" || words[i] == "An" {
				if !silentHWords[words[i+1]] && !strings.ContainsRune("aeiouAEIOU", rune(words[i+1][0])) {
					if words[i] == "an" {
						words[i] = "a"
					} else {
						words[i] = "A"
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}
