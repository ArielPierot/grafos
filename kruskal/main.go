package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var V, operacoes int

type arestas struct {
	origem, destino, peso int
}

type grafos struct {
	a [] arestas
}

type subconjuntos struct {
	parente, classe int
}

func main() {

	var grafo grafos

	arquivoEntrada := os.Args[1] // Entrada -- O primeiro argumento entra como parametro da localização do arquivo de entrada

	V = entraArquivo(arquivoEntrada, &grafo)

	grafo.kruskal()
}

func (a arestas) comparador(e arestas)  int {
	return a.peso - e.peso
}

func procurar(scj []subconjuntos, i int) int {
	if scj[i].parente != i {
		scj[i].parente = procurar(scj, scj[i].parente)
	}

	return scj[i].parente
}

func uniao(scj []subconjuntos, x int, y int){

	raizX := procurar(scj, x)
	raizY := procurar(scj, y)

	if scj[raizX].classe < scj[raizY].classe {
		scj[raizX].parente = raizY
	} else if scj[raizX].classe > scj[raizY].classe {
		scj[raizY].parente = raizX
	}

	if scj[raizX] == scj[raizY] {
		scj[raizY].classe++
	}

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
				grafo.a = append(grafo.a, arestas{verticeOrigem, verticeDestino, s})
			}
			verticeDestino++
		}

		verticeOrigem++
	}

	arquivo.Close()

	return MAX
}

func (grafo grafos) kruskal() {

	var resultado []arestas

	scj := make([]subconjuntos, V)

	sort.Slice(grafo.a, func(i, j int) bool {
		return grafo.a[i].peso < grafo.a[j].peso
	})

	for j := 0; j < V; j++ {
		resultado = append(resultado, arestas{})
		scj = append(scj, subconjuntos{parente:j, classe:0})
	}

	e, i := 0, 0

	for e < V - 1 {

		var proximaAresta arestas
		proximaAresta = grafo.a[i]
		i++

		x := procurar(scj, proximaAresta.origem)
		y := procurar(scj, proximaAresta.destino)

		if x != y {
			resultado[e] = proximaAresta
			e++
			uniao(scj, x, y)
		}


	}

	exibirSolucao(resultado)
}

func exibirSolucao(resultado []arestas) {
	fmt.Println("Vértice de destino \t Peso")
	for i := 0; i < V; i++ {
		fmt.Println(resultado[i].destino," \t ", resultado[i].peso)
	}

	fmt.Println("Operações: " + strconv.Itoa(operacoes))
}
