package pkgui

import (
	"fmt"
	"strconv"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
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
	{Caption: "零件号", Width: 180},
	{Caption: "客户零件号", Width: 180},
	{Caption: "客户", Width: 150},
	{Caption: "订单日期", Width: 150},
	{Caption: "状态", Width: 80},
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
		fmt.Println(f.ListView1.ItemIndex())
		// f.Close()
	})
	item2 := vcl.NewMenuItem(f)
	item2.SetCaption("复制零件号")
	item.SetOnClick(func(vcl.IObject) {
		fmt.Println(f.ListView1.ItemIndex())
		// f.Close()
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
	for i := 1; i <= 100; i++ {
		item := f.ListView1.Items().Add()
		item.SetCaption(strconv.Itoa(i + 1))
		item.SubItems().Add("CP492SG-A")
		item.SubItems().Add("YV2000-1003103SF1")
		item.SubItems().Add("YV2000-1003103SF1")
		item.SubItems().Add("安徽全柴动力")
		item.SubItems().Add("2022-07-16 21:41")
		item.SubItems().Add(strconv.Itoa(int(i)))
	}
	f.ListView1.Items().EndUpdate()
	// 产生100w条数据
	// tempData = make([]TTempItem, 100)
	// for i := 0; i < len(tempData); i++ {
	// 	tempData[i].Sub1 = fmt.Sprintf("%d", i+1)
	// 	tempData[i].Sub2 = "CP492SG-A"
	// 	tempData[i].Sub3 = "YV2000-1003103SF1"
	// 	tempData[i].Sub4 = "YV2000-1003103SF1"
	// 	tempData[i].Sub5 = "安徽全柴动力"
	// 	tempData[i].Sub6 = "2022-07-16 21:41"
	// 	tempData[i].Sub7 = strconv.Itoa(int(i))
	// }
	f.ListView1.Items().SetCount(int32(len(tempData))) //   必须主动的设置Virtual List的行数

}

func (f *TForm1) OnMenuItem1Click(sender vcl.IObject) {

}

func (f *TForm1) OnLabel1Click(sender vcl.IObject) {

}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {

}

func (f *TForm1) OnTabControl1Change(sender vcl.IObject) {

}

// 弹窗菜单
func (f *TForm1) OnListView1PopupMenu(sender vcl.IObject, popupmenu *vcl.TPopupMenu) {
	// 右键菜单
	// 右键菜单
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
