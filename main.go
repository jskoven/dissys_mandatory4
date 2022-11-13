package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	token "github.com/jskoven/mandatory_handin_4_dissys/grpc"
	"google.golang.org/grpc"
)

func main() {
	f, err := os.OpenFile("Logs.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	log.SetOutput(f)

	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &peer{
		id:            ownPort,
		index:         int32(arg1),
		amountOfPings: make(map[int32]int32),
		clients:       make(map[int32]token.ExclusionClient),
		ctx:           ctx,
		token:         false,
	}

	//Seeding random generator, since we got the same random results before we did this. Also making sure first node starts with token.
	if p.id == 5000 {
		p.token = true
		p.random = 4
	} else if p.id == 5001 {
		p.random = 2
	} else if p.id == 5002 {
		p.random = 3
	}

	// Create listener tcp on port ownPort
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	grpcServer := grpc.NewServer()
	token.RegisterExclusionServer(grpcServer, p)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()

	for i := 0; i < 3; i++ {
		port := int32(5000) + int32(i)

		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		fmt.Printf("Trying to dial: %v\n", port)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := token.NewExclusionClient(conn)
		p.clients[port] = c
	}
	rand.Seed(int64(p.random))
	go p.hasToken()
	go p.calculateIfWorkToDo()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		/*Make this better lmao*/
		indextouse := int(p.id) + 1
		if indextouse == 5003 {
			indextouse = 5000
		}
		p.sendTokenToNext(indextouse)
	}
}

type peer struct {
	token.UnimplementedExclusionServer
	id            int32
	index         int32
	amountOfPings map[int32]int32
	clients       map[int32]token.ExclusionClient
	ctx           context.Context
	token         bool
	hasWorkToDo   bool
	random        int32
}

func (p *peer) hasToken() {
	for {
		if p.token {
			indextouse := int(p.id) + 1
			if indextouse == 5003 {
				indextouse = 5000
			}
			if p.hasWorkToDo {
				log.Printf("Node #%d has the token and wishes to work on critical section\n", p.id)
				time.Sleep(3 * time.Second)
				p.writeToCriticalSection("writing to critical section.")
				p.hasWorkToDo = false
				log.Printf("Node #%d has finished their work on critical section and is sending token to node #%d\n", p.id, indextouse)

			} else {
				log.Printf("Node #%d has the token\n", p.id)
			}

			time.Sleep(1 * time.Second)
			p.sendTokenToNext(indextouse)
			//emptybox := token.Empty{}
			//p.clients[int32(indextouse)].GiveToken(p.ctx, &emptybox)
		}

	}
}

func (p *peer) sendTokenToNext(index int) {
	if p.token {
		emptybox := token.Empty{}
		p.clients[int32(index)].GiveToken(p.ctx, &emptybox)
		p.token = false
	}

}

func (p *peer) GiveToken(context.Context, *token.Empty) (*token.Empty, error) {
	p.token = true
	emptybox := &token.Empty{}
	return emptybox, nil
}

func (p *peer) calculateIfWorkToDo() {
	for {
		switch p.random {
		case 1:
			p.hasWorkToDo = true
			p.random = rand.Int31n(12)
			time.Sleep(1 * time.Second)
		default:
			p.random = rand.Int31n(12)
			time.Sleep(1 * time.Second)
		}
	}
}

func (p *peer) writeToCriticalSection(toWrite string) {
	f, err := os.OpenFile("CriticalSection.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	_ = err
	log.SetOutput(f)

	log.Println("Node #", p.id, " ", toWrite)

	s, err := os.OpenFile("Logs.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	_ = err
	log.SetOutput(s)

}
