package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Token rappresenta un singolo elemento dell'espressione.
type Token struct {
	Tipo   string // "NUM", "OP", "LPAREN", "RPAREN"
	Valore string
	Numero float64
}

// Tokenizer converte una stringa in una lista di token.
func Tokenizza(expr string) ([]Token, error) {
	var tokens []Token
	i := 0
	for i < len(expr) {
		c := expr[i]
		if unicode.IsSpace(rune(c)) {
			i++
			continue
		}
		if c == '(' {
			tokens = append(tokens, Token{Tipo: "LPAREN", Valore: "("})
			i++
		} else if c == ')' {
			tokens = append(tokens, Token{Tipo: "RPAREN", Valore: ")"})
			i++
		} else if c == '+' || c == '-' || c == '*' || c == '/' {
			tokens = append(tokens, Token{Tipo: "OP", Valore: string(c)})
			i++
		} else if unicode.IsDigit(rune(c)) {
			j := i
			for j < len(expr) && unicode.IsDigit(rune(expr[j])) {
				j++
			}
			numStr := expr[i:j]
			num, _ := strconv.ParseFloat(numStr, 64)
			tokens = append(tokens, Token{Tipo: "NUM", Valore: numStr, Numero: num})
			i = j
		} else {
			return nil, fmt.Errorf("carattere non valido: %c", c)
		}
	}
	return tokens, nil
}

// Parser ricorsivo per valutare espressioni.
type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) Peek() Token {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	return Token{Tipo: "EOF"}
}

func (p *Parser) Consume() Token {
	t := p.tokens[p.pos]
	p.pos++
	return t
}

// Espressione = Termine { (+|-) Termine }
func (p *Parser) Espressione() (float64, error) {
	val, err := p.Termine()
	if err != nil {
		return 0, err
	}
	for p.Peek().Tipo == "OP" && (p.Peek().Valore == "+" || p.Peek().Valore == "-") {
		op := p.Consume().Valore
		destro, err := p.Termine()
		if err != nil {
			return 0, err
		}
		if op == "+" {
			val += destro
		} else {
			val -= destro
		}
	}
	return val, nil
}

// Termine = Fattore { (*|/) Fattore }
func (p *Parser) Termine() (float64, error) {
	val, err := p.Fattore()
	if err != nil {
		return 0, err
	}
	for p.Peek().Tipo == "OP" && (p.Peek().Valore == "*" || p.Peek().Valore == "/") {
		op := p.Consume().Valore
		destro, err := p.Fattore()
		if err != nil {
			return 0, err
		}
		if op == "*" {
			val *= destro
		} else {
			if destro == 0 {
				return 0, fmt.Errorf("divisione per zero")
			}
			val /= destro
		}
	}
	return val, nil
}

// Fattore = NUM | ( Espressione )
func (p *Parser) Fattore() (float64, error) {
	t := p.Peek()
	if t.Tipo == "NUM" {
		p.Consume()
		return t.Numero, nil
	}
	if t.Tipo == "LPAREN" {
		p.Consume()
		val, err := p.Espressione()
		if err != nil {
			return 0, err
		}
		if p.Peek().Tipo != "RPAREN" {
			return 0, fmt.Errorf("parentesi chiusa mancante")
		}
		p.Consume()
		return val, nil
	}
	return 0, fmt.Errorf("token inatteso: %v", t)
}

// ValutaEspressione valuta un'espressione aritmetica.
func ValutaEspressione(expr string) (float64, error) {
	tokens, err := Tokenizza(expr)
	if err != nil {
		return 0, err
	}
	if len(tokens) == 0 {
		return 0, fmt.Errorf("espressione vuota")
	}
	p := NewParser(tokens)
	result, err := p.Espressione()
	if err != nil {
		return 0, err
	}
	if p.pos < len(p.tokens) {
		return 0, fmt.Errorf("token inatteso dopo l'espressione")
	}
	return result, nil
}

// ContaOperatori conta gli operatori nell'espressione.
func ContaOperatori(expr string) map[string]int {
	count := map[string]int{"+": 0, "-": 0, "*": 0, "/": 0}
	for _, c := range expr {
		if c == '+' || c == '-' || c == '*' || c == '/' {
			count[string(c)]++
		}
	}
	return count
}

// ProfonditaParentesi restituisce la massima profondita di parentesi.
func ProfonditaParentesi(expr string) int {
	maxDepth := 0
	current := 0
	for _, c := range expr {
		if c == '(' {
			current++
			if current > maxDepth {
				maxDepth = current
			}
		} else if c == ')' {
			current--
		}
	}
	return maxDepth
}

// EValida verifica se l'espressione e sintatticamente valida.
func EValida(expr string) bool {
	depth := 0
	prevWasOp := true
	for i, c := range expr {
		if unicode.IsSpace(c) {
			continue
		}
		if c == '(' {
			depth++
			prevWasOp = true
		} else if c == ')' {
			depth--
			if depth < 0 {
				return false
			}
			prevWasOp = false
		} else if c == '+' || c == '-' || c == '*' || c == '/' {
			if prevWasOp && i > 0 {
				// controlla se c'e un operatore prima
				return false
			}
			prevWasOp = true
		} else if unicode.IsDigit(c) {
			prevWasOp = false
		} else {
			return false
		}
	}
	return depth == 0 && !prevWasOp
}

func main() {
	var expr string
	fmt.Scan(&expr)

	result, err := ValutaEspressione(expr)
	if err != nil {
		fmt.Printf("Errore: %v\n", err)
	} else {
		fmt.Printf("%s = %.2f\n", expr, result)
	}

	ops := ContaOperatori(expr)
	fmt.Println("\nOperatori:")
	for op, count := range ops {
		if count > 0 {
			fmt.Printf("%s: %d\n", op, count)
		}
	}

	fmt.Printf("\nMassima profondita parentesi: %d\n", ProfonditaParentesi(expr))
	fmt.Printf("Espressione valida: %v\n", EValida(expr))
	_ = strings.ToLower
}
