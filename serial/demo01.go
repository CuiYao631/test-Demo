package main

import (
	serial "github.com/tarm/goserial"
	"io"
	"log"
	"time"
)

//串口连接器
type SerialConnection struct {
	S      *io.ReadWriteCloser
	Ch     *chan []byte   //开启线程用于往外发送数据的channel
	StopCh *chan struct{} //用于管理线程的channel，关闭这个channel就关闭了串口的读取线程
}

func main() {
	sc := &SerialConnection{}
	err := sc.ConnectToSerial("/dev/cu.wchusbserial542A0409181", 115200)
	if err != nil {
		log.Println(err)
	}
	sc.ReadSerialLoop()
	sc.ReadSerialData()

}

//连接串口
func (sc *SerialConnection) ConnectToSerial(name string, baud int) (err error) {

	//设置串口编号
	c := &serial.Config{Name: name, Baud: baud, ReadTimeout: time.Millisecond * 100}
	ch := make(chan []byte, 128)
	c2 := make(chan struct{}, 10)
	sc.Ch = &ch

	//打开串口，初始化指针
	s, err := serial.OpenPort(c)
	if err != nil {
		return err
	}
	sc.S = &s
	sc.StopCh = &c2
	return nil
}
func (sc *SerialConnection) ReadSerial() {

	var num int
	for {

		select {
		case <-(*sc.StopCh): //关闭线程
			return
		default:
			buffer := make([]byte, 1024) //优化，如何在运行过程中清空数组从而不用重复分配内存
			num, _ = (*sc.S).Read(buffer)
			if num > 0 {
				*sc.Ch <- buffer
				//fmt.Println(string(buffer))
			}
		}
	}
}

func (sc *SerialConnection) ReadSerialLoop() {
	go func() {
		sc.ReadSerial()
	}()
}

// ReadSerialData 接收数据
func (sc *SerialConnection) ReadSerialData() {
	for {
		select {
		case data := <-*sc.Ch:
			log.Println("data:", string(data))

		}
	}
}
