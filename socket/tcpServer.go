package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var connMap = make(map[string]net.Conn)

// TCP 服务端
func process(conn net.Conn) {
	// 函数执行完之后关闭连接
	defer conn.Close()
	// 输出主函数传递的conn可以发现属于*TCPConn类型, *TCPConn类型那么就可以调用*TCPConn相关类型的方法, 其中可以调用read()方法读取tcp连接中的数据
	fmt.Printf("服务端: %T\n", conn)
	for {
		var buf [128]byte
		// 将tcp连接读取到的数据读取到byte数组中, 返回读取到的byte的数目
		n, err := conn.Read(buf[:])
		if err != nil {
			// 从客户端读取数据的过程中发生错误
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		//fmt.Println("服务端收到客户端发来的数据：", recvStr)
		if len(recvStr) > 8 {
			if recvStr[1:5] == "CMDR" {
				fmt.Println("接收到信息：", recvStr)

				instruct := strings.Split(recvStr, ",")
				switch instruct[4] {
				case "Q0": //签到
					log.Println("签到")
					//记录socket 信息
					connMap[instruct[2]] = conn
					//追踪
					ReceiveMsg(connMap[instruct[2]], instruct[2], "D1", "60")

				case "H0": //心跳
					log.Println("心跳")
				case "L0": //解锁指令
					log.Println("解锁指令")

					switch instruct[5] {
					case "0": //解锁成功
						log.Println("解锁成功")
					case "1": //解锁失败
						log.Println("解锁失败")
					}
					ReceiveMsg(connMap[instruct[2]], instruct[2], "Re", "L0")
				case "L1": //关锁指令，锁主动上报
					log.Println("关锁指令，锁主动上报")
					ReceiveMsg(connMap[instruct[2]], instruct[2], "Re", "L1")
				case "L3": //电动车开关机控制
					log.Println("电动车开关机控制")
				case "L5": //外部锁设备控制
					log.Println("外部锁设备控制")
				case "L6": //获取电池信息
					log.Println("获取电池信息")
				case "D0": //获取定位指令，单次
					log.Println("获取定位指令，单次")
					if instruct[5] == "0" {
						log.Println("指令获取定位上传标识")
					} else {
						log.Println("定位追踪上传定位标识")
					}
					log.Println("时间", instruct[6])
					if instruct[7] == "A" {
						fmt.Println("纬度:", instruct[8])
						fmt.Println("纬度半球:", instruct[9])
						fmt.Println("经度:", instruct[10])
						fmt.Println("经度半球:", instruct[11])
						fmt.Println("卫星个数:", instruct[12])
						fmt.Println("HDOP(定位精度):", instruct[13])
						fmt.Println("日期:", instruct[14])
						fmt.Println("海拔高度:", instruct[15])
						fmt.Println("高度单位:", instruct[16])
						fmt.Println("模式指示:", instruct[17])

					} else if instruct[7] == "V" {
						log.Println("无效定位")
					}

					ReceiveMsg(connMap[instruct[2]], instruct[2], "Re", "D0")
				case "D1": //定位追踪指令
					// ，定位使用D0指令上传
					log.Println("定位追踪指令，定位使用D0指令上传")
				case "S7": //锁功能设置
					log.Println("锁功能设置")
				case "S5": //获取锁信息
					log.Println("获取锁信息")
				case "S4": //电动车设置指令
					log.Println("电动车设置指令")
				case "S6": //获取电动车信息
					log.Println("获取电动车信息")
				case "S8": //找车指令
					log.Println("找车指令")
				case "G0": //获取锁固件信息
					log.Println("获取锁固件信息")
				case "G1": //获取外接设备固件版本
					log.Println("获取外接设备固件版本")
				case "W0": //报警指令
					log.Println("报警指令")
					switch instruct[5] {
					case "1": //非法移动报警
						log.Println("非法移动报警")
					case "2": //倒地报警
						log.Println("倒地报警")
					case "非法拆除报警":
						log.Println("非法拆除报警")
					case "6": //清除倒地警报
						log.Println("清除倒地警报")
					case "7": //清除非法拆除警报
						log.Println("清除非法拆除警报")
					}
					ReceiveMsg(connMap[instruct[2]], instruct[2], "Re", "W0")
				case "U0": //检测升级/启动升级
					log.Println("检测升级/启动升级")
				case "U1": //获取升级数据
					log.Println("获取升级数据")
				case "U2": //升级结果通知
					log.Println("升级结果通知")
				case "K0": //设置/获取 BLE 8字节通信 KEY
					log.Println("设置/获取 BLE 8字节通信 KEY")
				case "C0": //RFID 卡开锁 请求
					log.Println("RFID 卡开锁 请求")
					log.Println("卡号:", instruct[7])
				}
			}
		}
		// 由于是tcp连接所以双方都可以发送数据, 下面接收服务端发送的数据这样客户端也可以收到对应的数据
		inputReader := bufio.NewReader(os.Stdin)
		s, _ := inputReader.ReadString('\n')
		t := strings.Trim(s, "\r\n")
		// 向当前建立的tcp连接发送数据, 客户端就可以收到服务端发送的数据
		conn.Write([]byte(t))

	}
}

// ReceiveMsg 回复消息
func ReceiveMsg(conn net.Conn, imei, instructionType, instructionData string) {
	cmd := ""
	startBit := "0xFFFF"          //起始位
	instructionHeader := "*CMDS," //指令头
	t := time.Now().Format("060102150405")
	vendorCode := "OM," //厂商代码
	endBit := "\n"      //结束位

	cmd += startBit + instructionHeader + vendorCode + imei + "," + t + "," + instructionType + "," + instructionData + "#" + endBit
	log.Println("回复消息：", cmd)
	if len([]byte(cmd)) != 0 {
		conn.Write([]byte(cmd))
	}

}

func main() {
	// 监听当前的tcp连接
	listen, err := net.Listen("tcp", ":8732")
	fmt.Printf("服务端: %T=====\n", listen)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		fmt.Println("当前建立了tcp连接")

		// 对于每一个建立的tcp连接使用go关键字开启一个goroutine处理
		go process(conn)
		// 由于是tcp连接所以双方都可以发送数据, 下面接收服务端发送的数据这样客户端也可以收到对应的数据
		//go sendCommand(conn)
	}
}
