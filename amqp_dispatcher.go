package main

import "encoding/json"

type amqpDispatcher struct {
	tbd string
}

// TODO - TDD functionality for this

// DispatchMessage implementation of dispatch message interface method
func (q *amqpDispatcher) DispatchMessage(queue string, message interface{}) (err error) {
	_, err = json.Marshal(message)
	return
}
