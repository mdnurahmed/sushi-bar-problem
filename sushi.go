package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Restaurant struct {
	in  chan *Customer
	out chan *Customer
}

func (r *Restaurant) WaitingToBeSeated(c *Customer) {
	log.Printf("customer - %s is waiting to be seated \n", c.Name)
	r.in <- c
}
func (r *Restaurant) Sit(c *Customer) {
	<-c.seated
	time.Sleep(time.Duration(500+rand.Intn(2000)) * time.Millisecond)
}
func (r *Restaurant) Leave(c *Customer) {
	r.out <- c
}
func (r *Restaurant) Run() {
	n := 0
	in := r.in
	for {
		select {
		case c := <-in:
			//wont block at all even if that customer is not ready to recieve as it is buffered of 1 capacity
			log.Printf("customer - %s sat down and now eating\n", c.Name)
			c.seated <- true
			n++
			log.Printf("total sitting after customer %s sat down--------------------%d\n", c.Name, n)
			if n == 5 {
				log.Printf("table full .New customers have to wait now. \n")
				in = nil
			}
		case c := <-r.out:
			log.Printf("customer - %s left the restaurent\n", c.Name)
			n--
			log.Printf("total customers after customer %s left--------------------%d\n", c.Name, n)
			if n == 0 {
				log.Printf("table empty. customers can start sitting again now\n")
				in = r.in
			}
		}
	}
}
func NewRestaurant() *Restaurant {
	return &Restaurant{
		in:  make(chan *Customer, 5),
		out: make(chan *Customer, 5),
	}
}

type Customer struct {
	Name   string
	seated chan bool
}

func (c *Customer) Run(r *Restaurant) {
	r.WaitingToBeSeated(c)
	r.Sit(c)
	r.Leave(c)
}

func NewCustommer(n int) *Customer {
	return &Customer{Name: fmt.Sprintf("%d", n), seated: make(chan bool, 1)}
}

func main() {
	r := NewRestaurant()

	// Generate customers every random second
	go func() {
		id := 0
		for {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			c := NewCustommer(id)
			id++
			go c.Run(r)
		}
	}()

	r.Run()
}
