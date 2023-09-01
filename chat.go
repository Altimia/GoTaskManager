package main

type Chat struct {
	ID      int
	Message string
	From    User
	To      User
}

func NewChat(id int, message string, from User, to User) *Chat {
	return &Chat{
		ID:      id,
		Message: message,
		From:    from,
		To:      to,
	}
}

func (c *Chat) SendMessage() {
	// code for sending a message
}

func (c *Chat) ReceiveMessage() {
	// code for receiving a message
}
