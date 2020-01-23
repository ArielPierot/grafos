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

	arquivoEntrada := os.Args[1] // Entrada

	array, V = entraArquivo(arquivoEntrada, array)

	dijsktra(array, 0)
}

func distanciaMin(distArray []int, isLower []bool) int {
	min := math.MaxInt32
	min_index := -1

	for i := 0; i < V; i++ {
		if isLower[i] == false && distArray[i] <= min {
			operacoes++
			min = distArray[i]
			min_index = i
		}
	}
	return min_index
}

func dijsktra(grafo [][]int, origem int){

	dist := make([]int, V)
	isLower := make([]bool, V)

	for i:=1; i < V; i++ {
		dist[i] = math.MaxInt32
		isLower[i] = false
	}

	dist[origem] = 0

	for contador := 0; contador < V;  contador++ {
		u := distanciaMin(dist, isLower)
		isLower[u] = true

		for p:=0; p<V; p++ {
			if !isLower[p] && grafo[u][p] != 0 && dist[u] != math.MaxInt32 && dist[u] + grafo[u][p] < dist[p] {
				operacoes++
				dist[p] = dist[u] + grafo[u][p]
			}
		}
	}

	exibirSolucao(dist)
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

func exibirSolucao(dist []int)  {
	fmt.Println("Vértice \t Distância do vértice")
	for i:=0; i < V; i++ {
		fmt.Println(i, "\t\t",dist[i])
	}

	fmt.Println("Operações: " + strconv.Itoa(operacoes))
}