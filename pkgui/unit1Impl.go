package pkgui

import (
	"fmt"
	"process-order/global"
	"process-order/service"
	"strconv"
	"strings"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

//::private::
type TForm1Fields struct {
}

type ListColumns struct {
	Caption string
	Width   int32
}
type TTempItem struct {
	Sub1 string
	Sub2 string
	Sub3 string
	Sub4 string
	Sub5 string
	Sub6 string
	Sub7 string
}

var tempData []TTempItem

var ListColumnsList = []ListColumns{
	{Caption: "序号", Width: 50},
	{Caption: "图号", Width: 150},
	{Caption: "零件号", Width: 170},
	{Caption: "客户零件号", Width: 170},
	{Caption: "客户", Width: 150},
	{Caption: "订单日期", Width: 100},
	{Caption: "状态", Width: 150},
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ListView1.SetAlign(types.AlClient)
	f.ListView1.SetParent(f.TabSheet1)
	f.ListView1.SetViewStyle(types.VsReport)
	// f.ListView1.SetOwnerData(true)
	f.ListView1.SetGridLines(true)
	f.ListView1.SetReadOnly(true)
	f.ListView1.SetRowSelect(true)
	f.ListView1.SetOnData(f.OnListView1Data)
	// 排序箭头
	f.ListView1.SetAutoSortIndicator(true)
	f.ListView1.SetSortType(types.StBoth) // 按文本排序
	// 设置右键菜单
	pm := vcl.NewPopupMenu(f)
	// 一级菜单
	item := vcl.NewMenuItem(f)
	item.SetCaption("复制图号")
	item.SetOnClick(func(vcl.IObject) {
		// 复制图号到粘贴板
		text := f.ListView1.Selected().SubItems().CommaText()
		textSplit := strings.Split(text, ",")
		vcl.Clipboard.SetTextBuf(textSplit[0])
		fmt.Println(textSplit[0])
	})
	item2 := vcl.NewMenuItem(f)
	item2.SetCaption("复制零件号")
	item2.SetOnClick(func(vcl.IObject) {
		// 复制零件号到粘贴板
		text := f.ListView1.Selected().SubItems().CommaText()
		textSplit := strings.Split(text, ",")
		vcl.Clipboard.SetTextBuf(textSplit[1])
		fmt.Println(textSplit[1])
	})
	pm.Items().Add(item)
	pm.Items().Add(item2)
	f.ListView1.SetPopupMenu(pm)

	//
	for _, v := range ListColumnsList {
		col := f.ListView1.Columns().Add()
		col.SetCaption(v.Caption)
		col.SetWidth(v.Width)
	}
	f.ListView1.Items().BeginUpdate()
	for i, v := range global.OrderList {
		item := f.ListView1.Items().Add()
		item.SetCaption(strconv.Itoa(i + 1))
		item.SubItems().Add(v.DrawNo)
		item.SubItems().Add(v.PartNo)
		item.SubItems().Add(v.CustomerPartNo)
		item.SubItems().Add(v.Customer)
		item.SubItems().Add(v.OrderDate.Format("2006-01-02"))
		item.SubItems().Add(v.Status)
	}
	f.ListView1.Items().EndUpdate()

	f.ListView1.Items().SetCount(int32(len(tempData))) //   必须主动的设置Virtual List的行数

}

func (f *TForm1) OnMenuItem1Click(sender vcl.IObject) {
	win.MessageBox(0, "软件版本：V0.1\r\n软件开发：信息部\r\n作者：梁文杰", "工艺订单查询系统", win.MB_OK+win.MB_ICONINFORMATION)
}

func (f *TForm1) OnMenuItem2Click(sender vcl.IObject) {
	f.Close()
}

func (f *TForm1) OnLabel1Click(sender vcl.IObject) {

}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	// 查询
	f.ListView1.Items().BeginUpdate()
	f.ListView1.Items().Clear()
	index := 1
	for _, v := range global.OrderList {
		if strings.Index(v.DrawNo, strings.ToUpper(f.Edit1.Text())) != -1 {
			item := f.ListView1.Items().Add()
			item.SetCaption(strconv.Itoa(index))
			item.SubItems().Add(v.DrawNo)
			item.SubItems().Add(v.PartNo)
			item.SubItems().Add(v.CustomerPartNo)
			item.SubItems().Add(v.Customer)
			item.SubItems().Add(v.OrderDate.Format("2006-01-02"))
			item.SubItems().Add(v.Status)
			index++
		}
	}
	f.ListView1.Items().EndUpdate()
}

func (f *TForm1) OnButton2Click(sender vcl.IObject) {
	// 获取数据
	res, err := service.GetOrderList()
	if err != nil {
		fmt.Println(err)
	}
	global.OrderList = res
	f.ListView1.Items().BeginUpdate()
	f.ListView1.Items().Clear()
	index := 1
	for _, v := range global.OrderList {
		item := f.ListView1.Items().Add()
		item.SetCaption(strconv.Itoa(index))
		item.SubItems().Add(v.DrawNo)
		item.SubItems().Add(v.PartNo)
		item.SubItems().Add(v.CustomerPartNo)
		item.SubItems().Add(v.Customer)
		item.SubItems().Add(v.OrderDate.Format("2006-01-02"))
		item.SubItems().Add(v.Status)
		index++
	}
	f.ListView1.Items().EndUpdate()
}
func (f *TForm1) OnTabControl1Change(sender vcl.IObject) {

}

func (f *TForm1) OnListView1Data(sender vcl.IObject, item *vcl.TListItem) {
	data := tempData[int(item.Index())]
	// 第一列为Caption属性所管理
	item.SetCaption(data.Sub1)
	item.SubItems().Add(data.Sub2)
	item.SubItems().Add(data.Sub3)
	item.SubItems().Add(data.Sub4)
	item.SubItems().Add(data.Sub5)
	item.SubItems().Add(data.Sub6)
	item.SubItems().Add(data.Sub7)
}
