package main  
import (  
    "fmt"
    "reflect"
    "encoding/binary"
    "bytes"
    "hash/crc32"
)

type TCPFrame struct {
    SourcePort      uint16
    DestinationPort uint16
    SequenceNumber  uint32
    AcknowledgmentNumber uint32
    DataOffset      uint8
    Reserved        uint8
    Flags           uint8
    WindowSize      uint16
    Checksum        uint16
    UrgentPointer   uint16
    Options         []byte
    Data            []byte
}

func (f *TCPFrame) CalculateChecksum() {
    data := f.Serialize()
    checksum := crc32.ChecksumIEEE(data)
    f.Checksum = uint16(checksum)
}

func (f *TCPFrame) Serialize() []byte {
    var buf bytes.Buffer
    v := reflect.ValueOf(f)
    t := reflect.TypeOf(f)
    numFields := t.NumField()

    for i := 0; i < numFields; i++ {
        field := v.Field(i)
        fieldType := t.Field(i)

        switch field.Kind() {
        case reflect.Uint8, reflect.Uint16, reflect.Uint32:
            value := uint32(field.Uint())
            if err := binary.Write(&buf, binary.BigEndian, value); err != nil {
                panic(err)
            }
        case reflect.Slice:
            if fieldType.Name == "Options" || fieldType.Name == "Data" {
                buf.Write(field.Bytes())
            }
        default:
            panic(fmt.Sprintf("Unsupported field type: %v", field.Kind()))
        }
    }

    return buf.Bytes()
}

func (f *TCPFrame) Deserialize(data []byte) {
    buf := bytes.NewBuffer(data)
    v := reflect.ValueOf(f).Elem()
    t := reflect.TypeOf(f).Elem()
    numFields := t.NumField()

    for i := 0; i < numFields; i++ {
        field := v.Field(i)
        fieldType := t.Field(i)

        switch field.Kind() {
        case reflect.Uint8, reflect.Uint16, reflect.Uint32:
            var value uint32
            if err := binary.Read(buf, binary.BigEndian, &value); err != nil {
                panic(err)
            }
            field.SetUint(uint64(value))
        case reflect.Slice:
            if fieldType.Name == "Options" || fieldType.Name == "Data" {
                field.SetBytes(buf.Next(field.Len()))
            }
        default:
            panic(fmt.Sprintf("Unsupported field type: %v", field.Kind()))
        }
    }
}

func main() {
    frame := &TCPFrame{
        SourcePort:      1234,
        DestinationPort: 5678,
        SequenceNumber:  1000,
        AcknowledgmentNumber: 2000,
        DataOffset:      5,
        Reserved:        0,
        Flags:           0x02, // SYN flag
        WindowSize:      8192,
        Checksum:        0,  // Initialize with 0 for calculation
        UrgentPointer:   0,
        Options:         []byte{0x01, 0x03, 0x03, 0x07}, // MSS option
        Data:            []byte("Hello, World!"),
    }

    // Calculate the checksum
    frame.CalculateChecksum()

    // Inspect the frame using reflection
    v := reflect.ValueOf(frame).Elem()
    t := reflect.TypeOf(frame).Elem()