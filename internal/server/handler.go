package server

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"

	"github.com/jackc/pgproto3/v2"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Accepted connection from %v", conn.RemoteAddr())

	handleStartupMessage(conn)

	for {
		var msgType byte
		if err := binary.Read(conn, binary.BigEndian, &msgType); err != nil {
			log.Printf("Failed to read message type: %v", err)
			return
		}

		var length int32
		if err := binary.Read(conn, binary.BigEndian, &length); err != nil {
			log.Printf("Failed to read message length: %v", err)
			return
		}

		message := make([]byte, length-4)
		if _, err := conn.Read(message); err != nil {
			log.Printf("Failed to read message: %v", err)
			return
		}

		log.Printf("Received message: %v messageType: %v", string(message), msgType)

		switch msgType {
		case QueryMessage:
			handleQuery(conn, message)
		case StartupMessage:
			handleStartupMessage(conn)
		case ParseMessage:
			handleParseMessage(conn, message)
		case ParseCompleteMessage:
			handleParseCompleteMessage(conn)
		case DataRow:
			handleDataRow(conn, message, nil)
		case ParameterStatus:
			handleParameterStatus(conn, message)
		case Failed:
			log.Printf("Failed message: %v", string(message))
		case Terminate:
			log.Printf("Terminated message: %v", string(message))
			return
		default:
			log.Printf("Unknown message type: %v", msgType)
			return
		}
	}
}

func handleParameterStatus(conn net.Conn, message []byte) {

	_, err := conn.Write(mustEncode((&pgproto3.ParameterStatus{
		Name:  "server_version",
		Value: "1.0",
	}).Encode(nil)))

	if err != nil {
		log.Printf("Failed to write ParameterStatus message: %v", err)
		return
	}
}

func handleStartupMessage(conn net.Conn) {

	buf := mustEncode((&pgproto3.AuthenticationOk{}).Encode(nil))
	buf = mustEncode((&pgproto3.ReadyForQuery{TxStatus: 'I'}).Encode(buf))

	_, err := conn.Write(buf)

	if err != nil {
		log.Printf("Failed to write ReadyForQuery message: %v", err)
		return
	}
}

func handleDataRow(conn net.Conn, message []byte, values []interface{}) {
	dataRow := &pgproto3.DataRow{
		Values: make([][]byte, len(values)),
	}

	for i, value := range values {
		switch v := value.(type) {
		case int32:
			dataRow.Values[i] = []byte(strconv.Itoa(int(v)))
		case string:
			dataRow.Values[i] = []byte(v)
		default:
			dataRow.Values[i] = nil
		}
	}

	buf := mustEncode(dataRow.Encode(nil))

	_, err := conn.Write(buf)

	if err != nil {
		log.Printf("Failed to write Parse message: %v", err)
		return
	}
}

func handleParseCompleteMessage(conn net.Conn) {
	buf := mustEncode((&pgproto3.ParseComplete{}).Encode(nil))

	_, err := conn.Write(buf)

	if err != nil {
		log.Printf("Failed to write Parse message: %v", err)
		return
	}

}

func handleQuery(conn net.Conn, message []byte) {

	buf := mustEncode((&pgproto3.RowDescription{Fields: []pgproto3.FieldDescription{
		{
			Name:                 []byte("lock_status"),
			TableOID:             0,
			TableAttributeNumber: 0,
			DataTypeOID:          25,
			DataTypeSize:         -1,
			TypeModifier:         -1,
			Format:               0,
		},
	}}).Encode(nil))
	buf = mustEncode((&pgproto3.DataRow{Values: [][]byte{[]byte("f")}}).Encode(buf))
	buf = mustEncode((&pgproto3.CommandComplete{CommandTag: []byte("SELECT 1")}).Encode(buf))
	buf = mustEncode((&pgproto3.ReadyForQuery{TxStatus: 'I'}).Encode(buf))

	_, err := conn.Write(buf)
	if err != nil {
		log.Printf("Failed to write CommandComplete message: %v", err)
		return
	}
}

func handleParseMessage(conn net.Conn, message []byte) {
	buf := mustEncode((&pgproto3.Parse{}).Encode(nil))

	_, err := conn.Write(buf)

	if err != nil {
		log.Printf("Failed to write Parse message: %v", err)
		return
	}

}

func mustEncode(buf []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return buf
}
