package modbusconnectionmodel

type ModbusRtuConnection struct {
	Address            string // example "/dev/ttyUSB0"
	BaudRate           uint16 // example "9600"
	DataBits           uint8  // example "8"
	StopBits           uint8  // example "1"
	Parity             string // example "N"
	Timeout            uint32 // example 60000000000 # 1 min
	Enabled            bool   // example "true"
	DelayRtsBeforeSend uint32 // example "2000000" #2 milisec
	DelayRtsAfterSend  uint32 // example "3000000" #3 milisec
	RtsHighDuringSend  bool   // example "false"
	RtsHighAfterSend   bool   // example "false"
	RxDuringTx         bool   // example "false"
}

type ModbusTcpConnection struct {
	AddressTcp string // exmaple "10.10.16.85"
	PortTcp    string // exmaple "1502"
}

type Connection uint8

const (
	RtuConnection Connection = iota
	TcpConnection
)

// struct for dataStore
type ConnectionConfig struct {
	ConnectionType      Connection
	ModbusRtuConnection ModbusRtuConnection
	ModbusTcpConnection ModbusTcpConnection
}
