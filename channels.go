package browsergo

import (
	"sync"
)

var (
	messageSync  sync.Mutex
	listenerSync sync.Mutex
)

func (b *BrowserService) sendMessage(message map[string]interface{}) {
	messageSync.Lock()
	defer messageSync.Unlock()
	b.messages <- message
}

func (b *BrowserService) sendListener(message map[string]interface{}) {
	listenerSync.Lock()
	defer listenerSync.Unlock()
	b.requests <- message
}

func (b *BrowserService) listenMessage() {
	defer close(b.done)
	for {
		select {
		case <-b.CTX.Done():
			return
		default:
			ok := <-b.messages
			for _, s := range b.messageReceivers {
				s <- ok
			}
		}
	}
}

func (b *BrowserService) listenRequests() {
	defer close(b.done)
	for {
		select {
		case <-b.CTX.Done():
			return
		default:
			ok := <-b.requests
			for _, s := range b.requestReceivers {
				s <- ok
			}
		}
	}
}

func (b *BrowserService) receiveMessage() <-chan map[string]interface{} {
	messageSync.Lock()
	defer messageSync.Unlock()
	newListeners := make(chan map[string]interface{})
	b.messageReceivers = append(b.messageReceivers, newListeners)
	return newListeners
}

func (b *BrowserService) receiveListener() <-chan map[string]interface{} {
	listenerSync.Lock()
	defer listenerSync.Unlock()
	newListeners := make(chan map[string]interface{})
	b.requestReceivers = append(b.requestReceivers, newListeners)
	return newListeners
}

func (b *BrowserService) removeMessageListener() error {
	messageSync.Lock()
	defer messageSync.Unlock()
	for i, v := range b.messageReceivers {
		if v == b.messageListener {
			b.messageReceivers = append(b.messageReceivers[:i], b.messageReceivers[i+1:]...)
		}
	}
	return errCannotFindListner
}

func (b *BrowserService) removeReceiveListener() error {
	listenerSync.Lock()
	defer listenerSync.Unlock()
	for i, v := range b.requestReceivers {
		if v == b.requestListener {
			b.requestReceivers = append(b.requestReceivers[:i], b.requestReceivers[i+1:]...)
		}
	}
	return errCannotFindListner
}
