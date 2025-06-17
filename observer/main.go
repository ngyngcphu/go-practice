package main

import "time"

func main() {
	channel := NewYoutubeChannel("TechTutorials")

	alice := NewYoutubeSubscriber("Alice")
	bob := NewYoutubeSubscriber("Bob")
	emailSub := NewEmailSubscriber("alice@example.com")
	pushSub := NewPushNotificationSubscriber("Samsung-J7")

	channel.Subscribe(alice)
	channel.Subscribe(bob)
	channel.Subscribe(emailSub)
	channel.Subscribe(pushSub)

	channel.UploadVideo("Go Design Pattern Tutorial")
	time.Sleep(100 * time.Millisecond)

	channel.Unsubscribe(bob)
	channel.UploadVideo("Observer Pattern with Go Channels")
	time.Sleep(100 * time.Millisecond)

	channel.UploadVideo("Final Video")
	time.Sleep(50 * time.Millisecond)

	channel.Close()
	time.Sleep(100 * time.Millisecond)
}
