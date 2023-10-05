package crigo

import (
	"sync"
)

var (
	messageSync  sync.Mutex
	listenerSync sync.Mutex
)

func (b *CRIService) sendMessage(message map[string]interface{}) {
	messageSync.Lock()
	defer messageSync.Unlock()
	b.messages <- message
}

func (b *CRIService) sendListener(message map[string]interface{}) {
	listenerSync.Lock()
	defer listenerSync.Unlock()
	b.requests <- message
}

func (b *CRIService) listenMessage() {
	defer close(b.done)
	for {
		ok := <-b.messages
		for _, s := range b.messageReceivers {
			s <- ok
		}
	}
}

func (b *CRIService) listenRequests() {
	defer close(b.done)
	for {
		ok := <-b.requests
		for _, s := range b.requestReceivers {
			s <- ok
		}
	}
}

func (b *CRIService) receiveMessage() <-chan map[string]interface{} {
	messageSync.Lock()
	defer messageSync.Unlock()
	newListeners := make(chan map[string]interface{})
	b.messageReceivers = append(b.messageReceivers, newListeners)
	return newListeners
}

func (b *CRIService) receiveListener() <-chan map[string]interface{} {
	listenerSync.Lock()
	defer listenerSync.Unlock()
	newListeners := make(chan map[string]interface{})
	b.requestReceivers = append(b.requestReceivers, newListeners)
	return newListeners
}

func (b *CRIService) removeMessageListener() error {
	messageSync.Lock()
	defer messageSync.Unlock()
	for i, v := range b.messageReceivers {
		if v == b.messageListener {
			b.messageReceivers = append(b.messageReceivers[:i], b.messageReceivers[i+1:]...)
		}
	}
	return errCannotFindListner
}

func (b *CRIService) removeReceiveListener() error {
	listenerSync.Lock()
	defer listenerSync.Unlock()
	for i, v := range b.requestReceivers {
		if v == b.requestListener {
			b.requestReceivers = append(b.requestReceivers[:i], b.requestReceivers[i+1:]...)
		}
	}
	return errCannotFindListner
}

func (b *CRIService) closeChannels() error {
	select {
	case b.messages <- map[string]interface{}{}:
	case <-b.done:
	}
	select {
	case b.requests <- map[string]interface{}{}:
	case <-b.done:
	}

	return nil
}
