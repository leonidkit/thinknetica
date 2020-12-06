package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type table chan *ball

type ball struct {
	command string
	powered bool
}

type player struct {
	name  string
	score int
	luck  *rand.Rand
}

func (p *player) begin(table table) {
	fmt.Println("begin")

	table <- &ball{
		command: "begin",
	}
}

func (p *player) stop(table table) {
	fmt.Println("stop")

	table <- &ball{
		command: "stop",
	}
}

func (p *player) ping(table table) {
	fmt.Println(p.name + " - ping")

	table <- &ball{
		command: "ping",
		powered: p.boost(),
	}
}

func (p *player) pong(table table) {
	fmt.Println(p.name + " - pong")

	table <- &ball{
		command: "pong",
		powered: p.boost(),
	}
}

func (p *player) boost() bool {
	powered := false
	num := p.luck.Intn(100)
	if (num + 1) < 21 {
		powered = true
	}
	return powered
}

func (p *player) Play(table table, wg *sync.WaitGroup) {
	defer wg.Done()
OUT:
	for b := range table {
		time.Sleep(time.Millisecond * 10)

		if b.powered {
			p.stop(table)
			break
		}

		switch b.command {
		case "begin":
			p.ping(table)
		case "ping":
			p.pong(table)
		case "pong":
			p.ping(table)
		case "stop":
			p.score += 1
			break OUT
		}
	}
}

type Game struct {
	table   table
	pfirst  *player
	psecond *player
	ballnum int
	coin    *rand.Rand
	wg      *sync.WaitGroup
}

func (g *Game) Run() {
	for {
		if g.ballnum == 11 {
			fmt.Print(g)
			break
		}
		g.wg.Add(2)
		go g.pfirst.Play(g.table, g.wg)
		go g.psecond.Play(g.table, g.wg)

		if g.flipCoin() {
			g.pfirst.begin(g.table)
		} else {
			g.psecond.begin(g.table)
		}

		g.wg.Wait()

		g.ballnum++
	}
}

func (g *Game) flipCoin() bool {
	if g.coin.Int()%2 == 0 {
		return true
	}
	return false
}

func (g *Game) String() string {
	var res string
	res = fmt.Sprintf("[SCORE]:\n%s\t%d points\n%s\t%d points\n", g.pfirst.name, g.pfirst.score, g.psecond.name, g.psecond.score)
	if g.pfirst.score > g.psecond.score {
		res += fmt.Sprintf("%s player win!!!\n", g.pfirst.name)
		return res
	}
	if g.pfirst.score < g.psecond.score {
		res += fmt.Sprintf("%s player win!!!\n", g.psecond.name)
		return res
	}
	return res + fmt.Sprint("Draw!!!")
}

func main() {
	var wg sync.WaitGroup
	table := make(table)
	defer close(table)

	p1 := &player{
		name: "first",
		luck: r,
	}
	p2 := &player{
		name: "second",
		luck: r,
	}

	g := &Game{
		table:   table,
		pfirst:  p1,
		psecond: p2,
		ballnum: 0,
		coin:    r,
		wg:      &wg,
	}

	g.Run()
}
