package main

import (
	"fmt"
	"sync"
)

type YoutubeChannel struct {
	name        string
	subscribers map[string]Subscriber
	mu          sync.RWMutex
	notifyChan  chan string
	active      bool
	wg          sync.WaitGroup
}

func NewYoutubeChannel(name string) *YoutubeChannel {
	channel := &YoutubeChannel{
		name:        name,
		subscribers: make(map[string]Subscriber),
		notifyChan:  make(chan string, 100),
		active:      true,
	}

	go channel.notificationWorker()
	return channel
}

func (yc *YoutubeChannel) Subscribe(subscriber Subscriber) {
	yc.mu.Lock()
	defer yc.mu.Unlock()

	yc.subscribers[subscriber.GetID()] = subscriber
	fmt.Printf("✅ %s subscribed to %s\n", subscriber.GetID(), yc.name)
}

func (yc *YoutubeChannel) Unsubscribe(subscriber Subscriber) {
	yc.mu.Lock()
	defer yc.mu.Unlock()

	if _, exists := yc.subscribers[subscriber.GetID()]; exists {
		delete(yc.subscribers, subscriber.GetID())
		fmt.Printf("❌ %s unsubscribed from %s\n", subscriber.GetID(), yc.name)
	}
}

func (yc *YoutubeChannel) UploadVideo(video string) {
	fmt.Printf("\n🎥 %s uploaded: %s\n", yc.name, video)

	if yc.active {
		select {
		case yc.notifyChan <- video:
		default:
			fmt.Println("Notification channel is full, skipping notification")
		}
	}
}

func (yc *YoutubeChannel) Close() {
	fmt.Printf("\n🔒 Closing %s channel...\n", yc.name)
	yc.active = false
	close(yc.notifyChan)
}

func (yc *YoutubeChannel) notificationWorker() {
	for video := range yc.notifyChan {
		yc.notifySubscribers(video)
	}
}

func (yc *YoutubeChannel) notifySubscribers(video string) {
	yc.mu.RLock()
	subscribers := make([]Subscriber, 0, len(yc.subscribers))
	for _, sub := range yc.subscribers {
		subscribers = append(subscribers, sub)
	}
	yc.mu.RUnlock()

	for _, subscriber := range subscribers {
		yc.wg.Add(1)
		go func(sub Subscriber) {
			defer yc.wg.Done()
			sub.Update(video)
		}(subscriber)
	}

	yc.wg.Wait()
}
