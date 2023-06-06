package sevendays

import (
	"fmt"
	"github.com/tlanfer/alasbot"
	"io/ioutil"
	"math"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	PlayerCount        = "CurrentPlayers"
	MaxPlayers         = "MaxPlayers"
	ServerTime         = "CurrentServerTime"
	DayNightLength     = "DayNightLength"
	BloodMoonFrequency = "BloodMoonFrequency"
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
		return 0, 0, fmt.Errorf("props() failed: %w", err)
	}

	count, err := strconv.Atoi(props[PlayerCount])

	if err != nil {
		return -1, -1, fmt.Errorf("strconv.Atoi[PlayerCount] failed: %w", err)
	}

	max, err := strconv.Atoi(props[MaxPlayers])

	if err != nil {
		return -1, -1, fmt.Errorf("strconv.Atoi[MaxPlayers] failed: %w", err)
	}

	return count, max, nil
}

func (c *client) GameTime() (int, int, int, int, error) {
	props, err := c.props()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	serverTime, err := strconv.Atoi(props[ServerTime])

	if err != nil {
		return -1, -1, -1, -1, err
	}

	dayLength, err := strconv.Atoi(props[DayNightLength])

	if err != nil {
		return -1, -1, -1, -1, err
	}

	bloodMoonFrequency, err := strconv.Atoi(props[BloodMoonFrequency])

	if err != nil {
		return -1, -1, -1, -1, err
	}

	totalMinutes := (float64(serverTime) / 1000.0) * float64(dayLength)

	minutesInADay := 24 * 60
	minutesIntoTheDay := math.Mod(float64(totalMinutes), float64(minutesInADay))

	days := math.Floor(float64(totalMinutes)/float64(minutesInADay)) + 1
	hours := math.Floor(minutesIntoTheDay / 60)
	minutes := math.Mod(minutesIntoTheDay, 60)

	return int(days), int(hours), int(minutes), bloodMoonFrequency, nil
}

func (c *client) props() (map[string]string, error) {

	if time.Since(c.propertiesAge) < time.Duration(10*time.Second) {
		return c.properties, nil
	}

	conn, err := net.DialTimeout("tcp", c.addr, 5*time.Second)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to %v", c.addr)
	}

	bytes, err := ioutil.ReadAll(conn)

	lines := strings.Split(string(bytes), ";")

	properties := map[string]string{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		fmt.Println(line)
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			properties[parts[0]] = parts[1]
		}

	}

	c.properties = properties
	c.propertiesAge = time.Now()

	return properties, nil
}
