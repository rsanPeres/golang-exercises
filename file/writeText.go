package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	size, err := file.WriteString("\nHello, world!\n")
	if err != nil {
		panic(err)
	}

	size2, err := file.Write([]byte("escrevendo outros dados\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho : %d bytes\n", size)
	fmt.Printf("Arquivo criado com sucesso! Tamanho : %d bytes", size2)

	file.Close()

	//read

	fileOpen, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fileOpen))
	file.Close()

	//read lines

	fileLines, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(fileLines)
	buffer := make([]byte, 10)

	for {
		line, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:line]))
	}
}
