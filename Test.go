package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var result string

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Введите выражение в формате \n\"Text\" + \"Text\" или \"Text\" - \"Text\"\n\"Text\" * (1-10) или \"Text\" / (1-10)\n")
	userText, _ := reader.ReadString('\n')
	userText = strings.TrimSpace(userText)

	result, err := calculator(userText)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else if len(result) > 40 {
		fmt.Println(result)
		fmt.Printf("%v…\"", result[:41])
	} else {
		fmt.Println(result)
	}
}

func calculator(text string) (string, error) {
	//Проверка условия что есть арифметический знак
	if strings.ContainsAny(text, "+-*/") {
		// Обработка сложения
		if strings.Contains(text, "+") {
			part := strings.Split(text, "+")
			for i, _ := range part {
				part[i] = strings.TrimSpace(part[i])
			}
			//Проверка что обе части в кавычках
			if isQuotedString(part[0]) && isQuotedString(part[1]) {
				part[0] = strings.Trim(part[0], `"`)
				part[1] = strings.Trim(part[1], `"`)
				if textLenght(part[0]) {
					return "", fmt.Errorf("Длина первой строки больше 10 символов")
				} else if textLenght(part[1]) {
					return "", fmt.Errorf("Длина второй строки больше 10 символов")
				}
				result = (part[0] + part[1])
				return fmt.Sprintf("\"%s\"", result), nil
			} else {
				return "", fmt.Errorf("Неверный формат ввода: строки должны быть в кавычках")
			}
		}
		// Обработка вычитания
		if strings.Contains(text, "-") {
			part := strings.Split(text, "-")
			for i, _ := range part {
				part[i] = strings.TrimSpace(part[i])
			}
			if isQuotedString(part[0]) && isQuotedString(part[1]) {
				part[0] = strings.Trim(part[0], `"`)
				part[1] = strings.Trim(part[1], `"`)
				if textLenght(part[0]) {
					return "", fmt.Errorf("Длина первой строки больше 10 символов")
				} else if textLenght(part[1]) {
					return "", fmt.Errorf("Длина второй строки больше 10 символов")
				}
				if strings.Contains(part[0], part[1]) {
					result = strings.Replace(part[0], part[1], "", 1)
					return fmt.Sprintf("\"%s\"", result), nil
				} else {
					result = part[0]
					return fmt.Sprintf("\"%s\"", result), nil
				}
			} else {
				return "", fmt.Errorf("Неверный формат ввода: строки должны быть в кавычках")
			}
		}
		//Обработка умножения
		if strings.ContainsAny(text, "*") {
			part := strings.Split(text, "*")
			for i, _ := range part {
				part[i] = strings.TrimSpace(part[i])
			}
			if isQuotedString(part[0]) {
				part[0] = strings.Trim(part[0], `"`)
				if textLenght(part[0]) {
					return "", fmt.Errorf("Длина первой строки больше 10 символов")
				}
				count, err := strconv.Atoi(strings.TrimSpace(part[1]))
				if err != nil || count < 1 || count > 10 {
					return "", fmt.Errorf("вторая часть должна быть числом от 1 до 10")
				}
				result := strings.Repeat(part[0], count)
				return fmt.Sprintf("\"%s\"", result), nil
			} else {
				return "", fmt.Errorf("неверный формат ввода: первая часть должна быть строкой в кавычках")
			}
		} // Обработка деления
		if strings.Contains(text, "/") {
			part := strings.Split(text, "/")
			for i := range part {
				part[i] = strings.TrimSpace(part[i])
			}
			if isQuotedString(part[0]) {
				part[0] = strings.Trim(part[0], `"`)
				if textLenght(part[0]) {
					return "", fmt.Errorf("Длина первой строки больше 10 символов")
				}
				count, err := strconv.Atoi(part[1])
				if err != nil || count < 1 || count > 10 {
					return "", fmt.Errorf("вторая часть должна быть числом от 1 до 10")
					//} else if parts[1] == "0" {
					//	return "", fmt.Errorf("Деление на 0")
				}
				partLength := len(part[0]) / count
				if partLength == 0 {
					return "", fmt.Errorf("результат деления слишком мал")
				}
				result := part[0][:partLength]
				return fmt.Sprintf("\"%s\"", result), nil
			} else {
				return "", fmt.Errorf("неверный формат ввода: первая часть должна быть строкой в кавычках")
			}
		}
	}
	return "", fmt.Errorf("отсутствует оператор")
}

func isQuotedString(s string) bool {
	return len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"'
}

func textLenght(s string) bool {
	return len(s) > 10
}
