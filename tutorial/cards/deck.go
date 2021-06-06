package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)
/*
	create a new type deck, a slice of strings
*/

type deck []string

func (d deck) print(){
	fmt.Println("Printing cards in deck.")
	for i,card:=range d{
		fmt.Println(i,card)
	}
}

func newDeck() deck{
	suits :=[]string{"Hearts","Diamonds","Clubs","Spades"}
	values :=[]string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","Queen","King"}

	cards:=deck{}
	for _,suit:= range suits{
		for _,value:= range values{
			cards=append(cards,value+" of "+suit)
		}
	}
	return cards
}

func (d deck) deal(handSize int)(deck, deck){
	return d[:handSize],d[handSize:]
}

func (d deck) toString() string{
	return strings.Join([]string(d),",")
}

func (d deck) saveDeckToFile(filename string) error{
	fmt.Println("Saving to file")
	return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}

func readDeckFromFile(filename string) deck{
	fmt.Println("Reading from file")
	byteString,err:=ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Println("Error message : ", err)
		os.Exit(1)
	}
	return deck(strings.Split((string(byteString)),","))
}

func (d deck) shuffle(){
	source:=rand.NewSource(time.Now().UnixNano())
	r:=rand.New(source)
	for i:= range(d){
		newPosition:=r.Intn(len(d)-1)
		d[i],d[newPosition]=d[newPosition],d[i]
	}
}

