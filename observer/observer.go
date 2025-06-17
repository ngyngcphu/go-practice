package main

import "fmt"

type Subscriber interface {
	Update(video string)
	GetID() string
}

type YoutubeSubscriber struct {
	name string
}

func NewYoutubeSubscriber(name string) *YoutubeSubscriber {
	return &YoutubeSubscriber{name}
}

func (y *YoutubeSubscriber) Update(video string) {
	fmt.Printf("%s is watching the video: %s\n", y.name, video)
}

func (y *YoutubeSubscriber) GetID() string {
	return y.name
}

type EmailSubscriber struct {
	email string
}

func NewEmailSubscriber(email string) *EmailSubscriber {
	return &EmailSubscriber{email}
}

func (e *EmailSubscriber) Update(video string) {
	fmt.Printf("Sending email to %s: New video uploaded: %s\n", e.email, video)
}

func (e *EmailSubscriber) GetID() string {
	return e.email
}

type PushNotificationSubscriber struct {
	deviceID string
}

func NewPushNotificationSubscriber(deviceID string) *PushNotificationSubscriber {
	return &PushNotificationSubscriber{deviceID}
}

func (p *PushNotificationSubscriber) Update(video string) {
	fmt.Printf("Sending push notification to %s: New video uploaded %s\n", p.deviceID, video)
}

func (p *PushNotificationSubscriber) GetID() string {
	return p.deviceID
}
