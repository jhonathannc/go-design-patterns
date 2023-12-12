package main

import "fmt"

func main() {
	var builder = newNotificationBuilder()

	builder.SetTitle("New notification")
	builder.SetSubtitle("This is a subtitle")
	builder.SetIcon("icon.png")
	builder.SetImage("image.png")
	builder.SetPriority(5)
	builder.SetMessage("This is a basic Notification with some text in it")
	builder.SetType("alert")

	notif, err := builder.Build()
	if err != nil {
		fmt.Println("Error creating the notification", err)
	} else {
		fmt.Printf("Notification: %+v", notif)
	}
}
