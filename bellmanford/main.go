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

type aresta struct {
	origem, destino, peso int
}

type grafos struct {
	a [] aresta
}

func main() {

	var grafo grafos

	arquivoEntrada := os.Args[1] // Entrada -- O primeiro argumento entra como parametro da localização do arquivo de entrada

	V = entraArquivo(arquivoEntrada, &grafo)

	bellmanFord(grafo, 0) // entrada do array e o vértice de origem
}

func entraArquivo(arquivoEntrada string, grafo *grafos) (int) {
	arquivo, _ := os.Open(arquivoEntrada)
	scanner := bufio.NewScanner(arquivo)
	var MAX int
	verticeOrigem := 0

	for scanner.Scan() {
		MAX++
		input := scanner.Text()
		arraySeparado := strings.Split(input, ",")

		verticeDestino := 0

		for _, v := range arraySeparado {
			s, _ := strconv.Atoi(v)
			if s != 0 {
				operacoes++
				grafo.a = append(grafo.a, aresta{verticeOrigem, verticeDestino, s})
			}
			verticeDestino++
		}

		verticeOrigem++
	}

	arquivo.Close()

	return MAX
}

func bellmanFord(grafo grafos, origem int) {

	dist := make([]int, V)

	for i := 1; i < V; i++ {
		dist[i] = math.MaxInt32
	}

	dist[origem] = 0

	for i := 1; i < V - 1; i++{
		for j := 0; j < V; j++ {
			u := grafo.a[j].origem
			v := grafo.a[j].destino
			peso := grafo.a[j].peso

			if dist[u] != math.MaxInt32 && dist[u] + peso < dist[v] {
				dist[v] = dist[u] + peso
				operacoes++
			}

		}
	}

	exibirSolucao(dist)
}

func exibirSolucao(dist []int) {
	fmt.Println("Vértice \t Distância da origem")
	for i := 0; i < V; i++ {

		if dist [i] == math.MaxInt32 {
			fmt.Println(i, "\t\t inf.")
		} else {
			fmt.Println(i, "\t\t", dist[i])
		}

	}

	fmt.Println("Operações: " + strconv.Itoa(operacoes))
}
