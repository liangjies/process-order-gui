package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"process-order/initialize"
	"process-order/pkgui"
	"process-order/service"
	"time"

	"process-order/global"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

var (
	Server    string
	Port      string
	MainEXE   string
	UpdateEXE string
)

func init() {
	global.SYS_DB = initialize.MSSQLGorm()
}

// go build -i -ldflags="-s -w -H windowsgui" -tags tempdll
func main() {
	// 加载配置文件
	iniFile := vcl.NewIniFile(".\\Config.ini")
	defer iniFile.Free()
	// 读取配置文件
	Server = iniFile.ReadString("System", "Server", "")
	Port = iniFile.ReadString("System", "Port", "")
	MainEXE = iniFile.ReadString("System", "MainEXE", "")
	UpdateEXE = iniFile.ReadString("System", "UpdateEXE", "")
	// 检测是否有新版本
	if checkUpdate() {
		// 如果有新版本，则提示用户是否更新
		if win.MessageBox(0, "检测到新版本，是否更新？", "新版本检测", win.MB_OKCANCEL+win.MB_ICONINFORMATION) == types.MrOk {
			// 启动更新下载程序
			cmd := exec.Command(UpdateEXE, "", "", "")
			err := cmd.Start()
			if err != nil {
				fmt.Println("启动失败:", err)
				return
			} else {
				fmt.Println("启动成功!")
			}
			return
		}
	}
	// 获取数据
	res, err := service.GetOrderList()
	if err != nil {
		fmt.Println(err)
	}
	global.OrderList = res
	// GUI
	vcl.Application.SetScaled(true)
	vcl.Application.SetTitle("project1")
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&pkgui.Form1)
	vcl.Application.Run()
}

// 检测程序是否有更新
func checkUpdate() bool {
	updateURL := "http://" + Server + ":" + Port + "/updates/process-order/md5.txt"
	// 捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("检测更新异常:", err)
		}
	}()
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(updateURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// 检测是否有新版本
	f, err := os.Open(MainEXE)
	if err != nil {
		fmt.Println("Open file error:", err)
		return false
	}
	defer f.Close()
	// 获取文件的MD5值
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		fmt.Println("Copy file error:", err)
		return false
	}
	FileMd5 := h.Sum(nil)
	fmt.Println(fmt.Sprintf("FileMd5:%x", FileMd5))
	fmt.Println("body:", string(body))
	return fmt.Sprintf("%x", FileMd5) != string(body)
}
