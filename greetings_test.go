package main

import "testing"

func TestGreetings(t *testing.T) {
    message := greetings("Code Education Rocks!")
    if message == "" {
       t.Errorf("Erro ao gerar a mensagem de boas vindas!")
    }
}