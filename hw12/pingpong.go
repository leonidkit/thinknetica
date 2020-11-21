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

type Table chan *Ball

type Ball struct {
	command string
	powered bool
}

type Player struct {
	name  string
	score int
	luck  *rand.Rand
}

func (p *Player) begin(table Table) {
	fmt.Println("begin")

	table <- &Ball{
		command: "begin",
	}
}

func (p *Player) stop(table Table) {
	fmt.Println("stop")

	table <- &Ball{
		command: "stop",
	}
}

func (p *Player) ping(table Table) {
	fmt.Println(p.name + " - ping")

	table <- &Ball{
		command: "ping",
		powered: p.boost(),
	}
}

func (p *Player) pong(table Table) {
	fmt.Println(p.name + " - pong")

	table <- &Ball{
		command: "pong",
		powered: p.boost(),
	}
}

func (p *Player) boost() bool {
	powered := false
	num := p.luck.Intn(100)
	if (num + 1) < 20 {
		powered = true
	}
	return powered
}

func (p *Player) Play(table Table, wg *sync.WaitGroup) {
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
	table   Table
	pfirst  *Player
	ptwice  *Player
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
		g.wg.Add(1)
		go g.pfirst.Play(g.table, g.wg)

		g.wg.Add(1)
		go g.ptwice.Play(g.table, g.wg)

		if g.flipCoin() {
			g.pfirst.begin(g.table)
		} else {
			g.ptwice.begin(g.table)
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
	res = fmt.Sprintf("[SCORE]:\n%s\t%d points\n%s\t%d points\n", g.pfirst.name, g.pfirst.score, g.ptwice.name, g.ptwice.score)
	if g.pfirst.score > g.ptwice.score {
		res += fmt.Sprintf("%s player win!!!\n", g.pfirst.name)
		return res
	}
	if g.pfirst.score < g.ptwice.score {
		res += fmt.Sprintf("%s player win!!!\n", g.ptwice.name)
		return res
	}
	return res + fmt.Sprint("Draw!!!")
}

func main() {
	var wg sync.WaitGroup
	table := make(Table)

	p1 := &Player{
		name: "first",
		luck: r,
	}
	p2 := &Player{
		name: "twice",
		luck: r,
	}

	g := &Game{
		table:   table,
		pfirst:  p1,
		ptwice:  p2,
		ballnum: 0,
		coin:    r,
		wg:      &wg,
	}

	g.Run()
}
