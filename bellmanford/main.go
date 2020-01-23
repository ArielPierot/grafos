package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var V int


func main() {

	var array [][]int
	arquivoEntrada := os.Args[1] // Entrada

	array, V = entraArquivo(arquivoEntrada, array)

	//dijsktra(array, 0)
}

func exibirSolucao(dist []int)  {
	fmt.Println("Vértice \t Distância do vértice")
	for i:=0; i < V; i++ {
		fmt.Println(i, "\t\t",dist[i])
	}
}

func entraArquivo(arquivoEntrada string, array [][]int) ([][]int, int) {
	arquivo, _ := os.Open(arquivoEntrada)
	scanner := bufio.NewScanner(arquivo)
	var MAX int

	for scanner.Scan() {
		MAX++
		input := scanner.Text()
		arraySeparado := strings.Split(input,",")

		var linha []int

		for _, v := range arraySeparado {
			s, _ := strconv.Atoi(v);
			linha = append(linha, s)
		}
		array = append(array, linha)
	}

	arquivo.Close()
	return array, MAX
}
