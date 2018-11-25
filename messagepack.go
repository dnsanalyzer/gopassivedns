package main

import (
	"github.com/vmihailenco/msgpack"
)

// logEntry is the same as dnsLog without some fields which are not required
// for fluentd outputs.
type logEntry struct {
	Response_Code        int    `msgpack:"r"`
	Question             string `msgpack:"q"`
	Question_Type        string `msgpack:"qt"`
	Answer               string `msgpack:"a"`
	TTL                  uint32 `msgpack:"t"`
	Server               string `msgpack:"d"`
	Client               string `msgpack:"s"`
	Timestamp            string `msgpack:"ts"`
	Elapsed              int64  `msgpack:"e"`
	Level                string `msgpack:"level,omitempty"` // syslog level omitted if empty
	Length               int    `msgpack:"b"`
}

func (dle *dnsLogEntry) MarshalMsgpack() ([]byte, error) {
	return msgpack.Marshal(&logEntry{
		Response_Code:        dle.Response_Code,
		Question:             dle.Question,
		Question_Type:        dle.Question_Type,
		Answer:               dle.Answer,
		TTL:                  dle.TTL,
		Server:               dle.Server.String(),
		Client:               dle.Client.String(),
		Timestamp:            dle.Timestamp,
		Elapsed:              dle.Elapsed,
		Level:                dle.Level,
		Length:               dle.Length,
	})
}

// Yet to be finished UnmarshalMsgpack method.
func (dle *dnsLogEntry) UnmarshalMsgpack(data []byte) error {
	tmp := &dnsLogEntry{}
	if err := msgpack.Unmarshal(data, &tmp); err != nil {
		return err
	}

	return nil
}
