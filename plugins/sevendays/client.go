package sevendays

import (
	"github.com/tlanfer/alasbot"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	PlayerCount = "CurrentPlayers"
	MaxPlayers = "MaxPlayers"
	ServerTime = "CurrentServerTime"
)

type client struct {
	addr          string
	properties    map[string]string
	propertiesAge time.Time
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

func (c *client) GameTime() (int, error) {
	props, err := c.props()
	if err != nil {
		return 0, err
	}

	minutes, err := strconv.Atoi(props[ServerTime])

	if err != nil {
		return 0, err
	}

	//duration := time.Minute * time.Duration(minutes)

	return minutes, nil
}

func (c *client) props() (map[string]string, error) {

	if time.Since(c.propertiesAge) < time.Duration(10*time.Second) {
		return c.properties, nil
	}

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

	c.properties = properties
	c.propertiesAge = time.Now()

	return properties, nil
}
