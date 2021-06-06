package main

import (
	"testing"
)
func TestNewDeck(t *testing.T){
	d:=newDeck()
	if len(d)!=52{
		t.Errorf("Expected deck length is 52, got %v",len(d))
	}
	if d[0]!="Ace of Hearts"{
		t.Errorf("Expected first card Ace of Hearts %v",d[0])
	}
	if d[len(d)-1]!="King of Spades"{
		t.Errorf("Expected first card King of Spades %v",d[len(d)-1])
	}
}