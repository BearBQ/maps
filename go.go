package main

import "fmt"

var dictionary = map[string]string{ //Создам словарь
	"Дверь":  "Door",
	"Окно":   "Window",
	"Дом":    "House",
	"Собака": "Dog",
}

func translate(word string) { //функция перевода
	fmt.Println(dictionary[word])
}

func deleteWord(word string) {
	delete(dictionary, word)
	fmt.Println("Слово", word, "удалено. Итоговый словарь", dictionary)
}

func check(word string) {
	_, ok := dictionary[word]
	if ok {
		fmt.Println("Слово", word, "содержится в словаре")
	} else {
		fmt.Println("Слово", word, "не содержится в словаре")
	}
}
func main() {
	translate("Дверь") //переводим слово
	deleteWord("Собака")
	check("Жопа")
	check("Дом")

	messagesMap := make(map[string]int)
	messagesMap["Смирнов"] = 4
	messagesMap["Иванов"] = 41
	messagesMap["петров"] = 44
	messagesMap["сидоров"] = 14
	messagesMap["козлов"] = 74
	fmt.Println(messagesMap)
}
