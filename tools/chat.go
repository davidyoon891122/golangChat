package tools

type Chat struct {
	msg string
}

func InitChat() *Chat {
	chat := &Chat{}

	return chat
}

func (c *Chat) chatPacker(msg string) {
	c.putMessage(msg)
}

func (c *Chat) putMessage(msg string) {
	c.msg = msg
}

func readChat() string {
	chat := readString()

	return chat
}
