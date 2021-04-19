package myhellomodule

import "fmt"

type Message struct {
	prefix  string
	postfix string
}

//SayHello something to say haha
func SayHello(name string) string {

	return fmt.Sprintf("Welcome to the go world %s ", name)
}

func CreateMessage(prf string, pstf string) *Message {
	return &Message{prefix: prf, postfix: pstf}
}

func (msg *Message) PrintMe() {
	fmt.Printf("%s is what you say before saying %s", msg.prefix, msg.postfix)
}
