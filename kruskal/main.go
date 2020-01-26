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
	aresta [] arestas
}

type subconjuntos struct {
	adj, nivel int
}

func main() {

	var grafo grafos

	arquivoEntrada := os.Args[1] // Entrada -- O primeiro argumento entra como parametro da localização do arquivo de entrada

	V = entraArquivo(arquivoEntrada, &grafo)

	grafo.kruskal()
}

func procurar(scj []subconjuntos, i int) int {

	operacoes++
	if scj[i].adj != i {
		scj[i].adj = procurar(scj, scj[i].adj)
	}

	return scj[i].adj
}

func uniao(scj []subconjuntos, x int, y int){

	resX := procurar(scj, x)
	resY := procurar(scj, y)

	operacoes++
	if scj[resX].nivel < scj[resY].nivel {
		scj[resX].adj = resY
	} else if scj[resX].nivel > scj[resY].nivel {
		scj[resY].adj = resX
	} else {
		scj[resY].adj = resX
		scj[resX].nivel++
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
				grafo.aresta = append(grafo.aresta, arestas{verticeOrigem, verticeDestino, s})
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

	var scj []subconjuntos

	sort.Slice(grafo.aresta, func(i, j int) bool {
		return grafo.aresta[i].peso < grafo.aresta[j].peso
	})

	for j := 0; j < V; j++ {
		resultado = append(resultado, arestas{})
		scj = append(scj, subconjuntos{adj: j, nivel:0})
	}

	e, i := 0, 0

	for e < V - 1 {

		var proximaAresta arestas
		proximaAresta = grafo.aresta[i]
		i++

		x := procurar(scj, proximaAresta.origem)
		y := procurar(scj, proximaAresta.destino)

		operacoes++
		if x != y {
			resultado[e] = proximaAresta
			e++
			uniao(scj, x, y)
		}


	}

	exibirSolucao(resultado)
}

func exibirSolucao(resultado []arestas) {
	fmt.Println("V. Origem \tV. Destino \tPeso")
	for i := 1; i < V; i++ {
		 if resultado[i].peso != 0 {
			fmt.Println(resultado[i].origem, " \t\t", resultado[i].destino, " \t\t", resultado[i].peso)
		}
	}

	fmt.Println("Operações: " + strconv.Itoa(operacoes))
}
