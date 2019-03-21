package main

import (
	"os"
	"log"
	"bufio"
	"strings"
)

func listaPeriodicos() {

}

type Candidato struct {
	nome             string
	temProjetoFormal bool
	temPublicacao    bool
	instituicao      string
	notaTHE          float64
	N_ITC            float64
	N_Periodicos   float64
	N_Eventos      float64
	N_CV             float64
	N_F              float64
}

func (c *Candidato) setNotaITC(n float64) {
	c.N_ITC = n
}

func (c *Candidato) setN_CV(n float64) {
	c.N_CV = n
}

func (c *Candidato) setN_F(n float64) {
	c.N_F = n
}

func (c *Candidato) setN_Periodicos(n float64) {
	c.N_Periodicos = n
}

func (c *Candidato) setN_Eventos(n float64) {
	c.N_Eventos = n
}

func carregaCandidatos() map[string]Candidato {
	r := map[string]Candidato{}

	r["fabio"] = Candidato{nome: "fabio", temProjetoFormal: true, temPublicacao: true, instituicao: "", notaTHE: 0.0}
	r["george"] = Candidato{nome: "george", temProjetoFormal: false, temPublicacao: false, instituicao: "", notaTHE: 0.0}

	return r
}

func carregaIndices() map[string]float64 {
	r := map[string]float64{}

	r["A1"] = 1.0
	r["A2"] = 0.9
	r["B1"] = 0.75
	r["B2"] = 0.50
	r["B3"] = 0.20
	r["B4"] = 0.10
	r["B5"] = 0.05

	return r
}

func calculaITC(c Candidato) float64 {
	// N_ITC = N_Pro + N_Pub + N_Ran

	nota := 0.0
	if c.temProjetoFormal {
		nota += 4.0
	}

	if c.temPublicacao {
		nota += 4
	}
	nota += c.notaTHE

	return nota
}

func calculaNotaPeriodicos(c Candidato) float64 {
	r := 0.0

	return r
}

func calculaN_Eventos(c Candidato) float64 {
	r := 0.0

	var lines [] string

	// carrega arquivo de eventos

	file, err := os.Open("/Users/nsr/Dropbox/go/programacao-concorrente-e-distribuida/selecao/george_eventos_lattes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// incluir apenas linhas completas
		if len(line) >= 10 {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	linesTemp := [] string{}

	for i := 0; i < len(lines); i++ {
		s1 := strings.Split(lines[i]," . ")
		s2 := strings.Split(s1[1]," In: ")
		titulo := s2[0]
		evento := s2[1]
		titulo = titulo
		evento = evento
		//linesTemp = append(linesTemp, strings.TrimSpace(titulo + " " + evento + " "+"QualisXX"))
		//linesTemp = append(linesTemp, strings.TrimSpace(titulo))
		linesTemp = append(linesTemp, strings.TrimSpace(evento))
	}

	for i := 0; i < len(linesTemp); i++ {
		println(linesTemp[i])
		}

		os.Exit(0)
	return r
}

func calculaN_F(c Candidato) float64 {

	return c.N_CV*0.7 + c.N_ITC*0.3
}

func main() {
	//var indices map[string]int
	var candidatos map[string]Candidato

	// Carrega candidatos
	candidatos = carregaCandidatos()

	// ******** Calcula notas *****//
	for c := range candidatos {
		cTemp := candidatos[c]

		// nota ITC
		cTemp.setNotaITC(calculaITC(cTemp))

		// Nota Periodicos
		cTemp.setN_Periodicos(calculaNotaPeriodicos(cTemp))

		// Nota eventos
		cTemp.setN_Eventos(calculaN_Eventos(cTemp))

		// Nota final
		cTemp.setN_F(calculaN_F(cTemp))

		candidatos[c] = cTemp
	}

	//indices = carregaIndices()

	// Publicacoes
	//N_Periodicos := calcularNPeriodicos()
	//N_Eventos := calcularNEventos()

	//N_CV := N_Periodicos + N_Eventos

	// NOta final
	//N_F = N_CV * 0,7 + N_ITC * 0,3

	//N_F := N_CV * 0.7 + N_ITC * 0.3

	//for c := range candidatos {
	//	fmt.Printf("%s N_ITC[%f] N_CV[%f] N_F[%f] \n", c, candidatos[c].N_ITC, candidatos[c].N_CV, candidatos[c].N_F)
	//}
}
