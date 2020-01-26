package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var V, operacoes int

func main() {

	var array [][]int

	arquivoEntrada := os.Args[1] // Entrada -- O primeiro argumento entra como parametro da localização do arquivo de entrada

	array, V = entraArquivo(arquivoEntrada, array)

	prim(array) // entrada do array e o vértice de origem
}

func distanciaMin(distArray []int, isLower []bool) int {

	min := math.MaxInt32
	min_index := -1

	for i := 0; i < V; i++ {
		operacoes++
		if isLower[i] == false && distArray[i] < min {
			min = distArray[i]
			min_index = i
		}
	}
	return min_index
}

func prim(grafo [][]int) {

	prt := make([]int, V)
	dist := make([]int, V)
	isLower := make([]bool, V)
	for i := 1; i < V; i++ {
		dist[i] = math.MaxInt32
		isLower[i] = false
	}

	dist[0] = 0
	prt[0] = -1

	for c := 0; c < V - 1; c++ {
		u := distanciaMin(dist, isLower)
		isLower[u] = true

		for p := 0; p < V; p++ {
			operacoes++
			if grafo[u][p] != 0 && isLower[p] == false && grafo[u][p] < dist[p]{
				prt[p] = u
				dist[p] = grafo[u][p]
			}
		}
	}

	exibirSolucao(prt, grafo)
}

func entraArquivo(arquivoEntrada string, array [][]int) ([][]int, int) {
	arquivo, _ := os.Open(arquivoEntrada)
	scanner := bufio.NewScanner(arquivo)
	var MAX int

	for scanner.Scan() {
		MAX++
		input := scanner.Text()
		arraySeparado := strings.Split(input, ",")

		var linha []int

		for _, v := range arraySeparado {
			s, _ := strconv.Atoi(v)
			linha = append(linha, s)
		}
		array = append(array, linha)
	}

	arquivo.Close()
	return array, MAX
}

func exibirSolucao(prt []int, grafo [][]int) {
	fmt.Println("V. Origem \tV. Destino \tPeso")
	for i := 1; i < V; i++ {
		fmt.Println(i, " \t\t", prt[i], " \t\t", grafo[i][prt[i]])
	}

	fmt.Println("Operações: " + strconv.Itoa(operacoes))
}
