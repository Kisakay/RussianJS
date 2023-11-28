package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Main struct {
	extensionFile string
	keyWords      []struct {
		og string
		to string
	}
}

func NewMain() *Main {
	m := &Main{
		extensionFile: "rjs",
		keyWords: []struct {
			og string
			to string
		}{
			{og: "extends", to: "расширяет"},
			{og: "function", to: "функция"},
			{og: "if", to: "если"},
			{og: "else", to: "иначе"},
			{og: "for", to: "для"},
			{og: "while", to: "пока"},
			{og: "do", to: "делай"},
			{og: "switch", to: "переключить"},
			{og: "case", to: "случай"},
			{og: "break", to: "прервать"},
			{og: "continue", to: "продолжить"},
			{og: "return", to: "вернуть"},
			{og: "new", to: "новый"},
			{og: "class", to: "класс"},
			{og: "constructor", to: "конструктор"},
			{og: "super", to: "супер"},
			{og: "this", to: "этот"},
			{og: "try", to: "попробовать"},
			{og: "catch", to: "поймать"},
			{og: "throw", to: "бросить"},
			{og: "finally", to: "наконец"},
			{og: "async", to: "асинхронно"},
			{og: "await", to: "ожидать"},
			{og: "typeof", to: "тип"},
			{og: "instanceof", to: "экземпляр"},
			{og: "delete", to: "удалить"},
			{og: "default", to: "по умолчанию"},
			{og: "in", to: "в"},
			{og: "with", to: "с"},
			{og: "debugger", to: "отладчик"},
			{og: "let", to: "пусть"},
			{og: "var", to: "переменная"},
			{og: "const", to: "конwithтанта"},
			{og: "yield", to: "уступать"},
			{og: "eval", to: "оценить"},
			{og: "arguments", to: "аргументы"},
			{og: "console", to: "консоль"},
			{og: "log", to: "логгер"},
			{og: "warn", to: "предупреждение"},
			{og: "error", to: "ошибка"},
		},
	}
	return m
}

func (m *Main) loadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *Main) replaceKeywords(str string) string {
	for _, kw := range m.keyWords {
		str = strings.ReplaceAll(str, kw.og, kw.to)
	}
	return str
}

func (m *Main) convertCode(str string) string {
	return m.replaceKeywords(str)
}

func (m *Main) writeFile(path, str string) error {
	err := ioutil.WriteFile(path, []byte(str), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (m *Main) russianConvert(str string) string {
	for _, kw := range m.keyWords {
		str = strings.ReplaceAll(str, kw.to, kw.og)
	}
	return str
}

func main() {
	mainObj := NewMain()

	// Utilisation du code
	filePath := "test/in.js"
	loadedContent, err := mainObj.loadFile(filePath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	// Conversion du code JavaScript en russe
	convertedToRussian := mainObj.convertCode(loadedContent)
	// fmt.Println("Code JavaScript converti en russe :", convertedToRussian)

	// Convertir le code russe en JavaScript
	convertedToJS := mainObj.russianConvert(convertedToRussian)
	fmt.Println("Code russe converti en JavaScript :", convertedToJS)

	// Écriture du code russe dans un fichier
	russianFilePath := "test/out.rjs"
	err = mainObj.writeFile(russianFilePath, convertedToRussian)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier russe :", err)
		return
	}
}
