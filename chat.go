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
	fmt.Printf("%s sent a message: %s\n", c.From.Username, c.Message)
}

func (c *Chat) ReceiveMessage() {
	fmt.Printf("%s received a message: %s\n", c.To.Username, c.Message)
}
