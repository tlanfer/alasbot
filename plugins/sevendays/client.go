package sevendays

import (
	"github.com/tlanfer/alasbot"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
)

const (
	PlayerCount = "CurrentPlayers"
	MaxPlayers = "MaxPlayers"
)

type client struct {
	addr string
}

func New(addr string) alasbot.Game {
	return &client{
		addr: addr,
	}
}

func (c *client) PlayerCount() (int, int, error) {
	props, err := c.props()
	if err != nil {
		return 0, 0, err
	}

	count, err := strconv.Atoi(props[PlayerCount])

	if err != nil {
		return -1, -1, err
	}

	max, err := strconv.Atoi(props[MaxPlayers])

	if err != nil {
		return -1, -1, err
	}

	return count, max, nil
}

func (c *client) props() (map[string]string, error) {

	conn, err := net.Dial("tcp", c.addr)

	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(conn)

	lines := strings.Split(string(bytes), ";")

	properties := map[string]string{}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			properties[parts[0]] = parts[1]
		}

	}

	return properties, nil
}
