package serial

import (
	"bytes"
	"strings"

	sp "github.com/tarm/serial"
)

// TODO: implement io.Reader / io.Writer on connection

type Connection struct {
	Port   *sp.Port
	Config *sp.Config
}

func NewConnection(device string, baudrate int) (*Connection, error) {
	c := &sp.Config{
		Name: device,
		Baud: baudrate,
	}

	s, err := sp.OpenPort(c)
	if err != nil {
		return nil, err
	}

	if err := s.Flush(); err != nil {
		return nil, err
	}

	return &Connection{
		Port:   s,
		Config: c,
	}, nil
}

func (c *Connection) Disconnect() error {
	return c.Port.Close()
}

func (c *Connection) Write(data []byte) error {
	_, err := c.Port.Write(data)

	return err
}

func (c *Connection) WriteString(data string) error {
	return c.Write([]byte(data))
}

func (c *Connection) Read() ([]byte, error) {
	buf := bytes.NewBuffer([]byte(nil))

	for {
		data := make([]byte, 512)
		_, err := c.Port.Read(data)
		if err != nil {
			return []byte(nil), err
		}

		data = bytes.Trim(data, "\x00")

		buf.Write(data)

		if strings.Contains(string(data), "ok\n") {
			break
		}
	}

	return buf.Bytes(), nil
}

func (c *Connection) ReadString() (string, error) {
	d, err := c.Read()

	return string(d), err
}
