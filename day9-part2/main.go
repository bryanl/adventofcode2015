package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	re = regexp.MustCompile(`^(\w+)\s+to\s+(\w+)\s+=\s+(\d+)$`)
)

func main() {
	//r := strings.NewReader(testIn)
	r := strings.NewReader(in)
	s := bufio.NewScanner(r)

	dc := newDistCalc()

	for s.Scan() {
		dc.addDist(s.Text())
	}

	// brute force all the things
	counts := dc.genCombos()
	max := counts[0]
	for _, c := range counts {
		if c > max {
			max = c
		}
	}

	fmt.Println(max)

}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("couldn't convert %q to int: %v", s, err)
	}

	return i
}

type distCalc struct {
	cList map[string]struct{}
	g     *graph
}

func newDistCalc() *distCalc {
	return &distCalc{
		cList: map[string]struct{}{},
		g:     newGraph(),
	}
}

func (dc *distCalc) addDist(inStr string) {
	res := re.FindAllStringSubmatch(inStr, -1)
	c1 := res[0][1]
	c2 := res[0][2]
	d := toInt(res[0][3])

	dc.g.addEdge(c1, c2, d)

	dc.cList[c1] = struct{}{}
	dc.cList[c2] = struct{}{}
}

func (dc *distCalc) genCombos() []int {
	a := dc.cities()
	k := len(a)

	counts := []int{}

	nextProduct := func(a []string, r int) func() []string {
		p := make([]string, r)
		x := make([]int, len(p))
		return func() []string {
			p := p[:len(x)]
			for i, xi := range x {
				p[i] = a[xi]
			}
			for i := len(x) - 1; i >= 0; i-- {
				x[i]++
				if x[i] < len(a) {
					break
				}
				x[i] = 0
				if i <= 0 {
					x = x[0:0]
					break
				}
			}
			return p
		}
	}

	np := nextProduct(a, k)
	for {
		product := np()
		if len(product) == 0 {
			break
		}

		invalid := false
		dests := map[string]bool{}

		for i := 0; i < len(product)-1; i++ {
			if product[i] == product[i+1] {
				invalid = true
				break
			}

			if dests[product[i+1]] {
				invalid = true
				break
			}
			dests[product[i+1]] = true

			sl := []string{product[i], product[i+1]}
			sort.Strings(sl)

			if e := dc.g.findEdge(product[i], product[i+1]); e == nil {
				invalid = true
				break
			}
			product[i] = fmt.Sprintf("%s-%s", sl[0], sl[1])
		}

		product = product[:len(product)-1]

		if invalid {
			continue
		}

		for i, s := range product {
			for j := 0; j < len(product); j++ {
				if i != j && product[j] == s {
					invalid = true
				}
			}
		}

		if invalid {
			continue
		}

		count := 0

		for i := 0; i < len(product); i++ {
			names := strings.Split(product[i], "-")
			e := dc.g.findEdge(names[0], names[1])
			count += e.weight
		}

		fmt.Println(count, product)

		counts = append(counts, count)
	}

	return counts

}

func (dc *distCalc) cities() []string {
	cities := []string{}

	for _, v := range dc.g.nodes {
		cities = append(cities, v.name)
	}

	return cities
}

type graph struct {
	edges []edge
	nodes map[string]*node
}

func newGraph() *graph {
	return &graph{
		edges: []edge{},
		nodes: map[string]*node{},
	}
}

func (g *graph) addEdge(from, to string, dist int) {
	fromNode := g.initNode(from)
	toNode := g.initNode(to)

	e := newEdge(fromNode, toNode, dist)
	g.edges = append(g.edges, e)
}

func (g *graph) initNode(name string) *node {
	n := g.nodes[name]
	if n == nil {
		n = newNode(name)
		g.nodes[name] = n
	}

	return n
}

func (g *graph) findEdges(n *node) []edge {
	edges := []edge{}
	for _, e := range g.edges {
		if e.a == n || e.b == n {
			edges = append(edges, e)
		}
	}

	return edges
}

func (g *graph) findEdge(nameA, nameB string) *edge {
	a := g.nodes[nameA]
	b := g.nodes[nameB]

	for _, e := range g.edges {
		if (e.a == a && e.b == b) || (e.a == b && e.b == a) {
			return &e
		}
	}

	return nil
}

type edge struct {
	a, b   *node
	weight int
}

func newEdge(a, b *node, weight int) edge {
	return edge{
		a:      a,
		b:      b,
		weight: weight,
	}
}

type node struct {
	name string
}

func newNode(name string) *node {
	return &node{
		name: name,
	}
}

var (
	testIn = `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`

	in = `Faerun to Tristram = 65
Faerun to Tambi = 129
Faerun to Norrath = 144
Faerun to Snowdin = 71
Faerun to Straylight = 137
Faerun to AlphaCentauri = 3
Faerun to Arbre = 149
Tristram to Tambi = 63
Tristram to Norrath = 4
Tristram to Snowdin = 105
Tristram to Straylight = 125
Tristram to AlphaCentauri = 55
Tristram to Arbre = 14
Tambi to Norrath = 68
Tambi to Snowdin = 52
Tambi to Straylight = 65
Tambi to AlphaCentauri = 22
Tambi to Arbre = 143
Norrath to Snowdin = 8
Norrath to Straylight = 23
Norrath to AlphaCentauri = 136
Norrath to Arbre = 115
Snowdin to Straylight = 101
Snowdin to AlphaCentauri = 84
Snowdin to Arbre = 96
Straylight to AlphaCentauri = 107
Straylight to Arbre = 14
AlphaCentauri to Arbre = 46`
)
