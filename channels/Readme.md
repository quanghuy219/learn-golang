# Channels and Go Routines

## 1. Routines
Create a new thread and run the function within it
```go
go <function call>
```
- A main routine is create when program is launched
- Then child routines are created by the `go` keyword

## 2. Channels
- Channels are means of communication between go routines
- Messages sent through a channel must be of the same type
- Usage
```go
// Create a new channel of type string
channel := make(chan string)

// Send value into channel
channel <- "value"

// Wait for a value to be sent into the channel and then assign the value to a variable
myStr <- channel

// Use the value sent into the channel as function argument
fmt.Println(<- channel)
```
