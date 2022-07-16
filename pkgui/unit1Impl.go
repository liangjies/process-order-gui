package pkgui

import (
	"fmt"
	"math/rand"

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
}

var tempData []TTempItem

var ListColumnsList = []ListColumns{
	{Caption: "序号", Width: 50},
	{Caption: "图号", Width: 100},
	{Caption: "零件号", Width: 100},
	{Caption: "客户零件号", Width: 100},
	{Caption: "客户", Width: 100},
	{Caption: "订单日期", Width: 100},
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ListView1.SetAlign(types.AlClient)
	f.ListView1.SetParent(f.TabSheet1)
	f.ListView1.SetViewStyle(types.VsReport)
	f.ListView1.SetOwnerData(true)
	f.ListView1.SetVisible(true)
	f.ListView1.SetGridLines(true)
	f.ListView1.SetReadOnly(true)

	f.ListView1.SetRowSelect(true)
	f.ListView1.SetOnData(f.OnListView1Data)

	for _, v := range ListColumnsList {
		col := f.ListView1.Columns().Add()
		col.SetCaption(v.Caption)
		col.SetWidth(v.Width)
	}

	// 产生100w条数据
	tempData = make([]TTempItem, 100)
	for i := 0; i < len(tempData); i++ {
		tempData[i].Sub1 = fmt.Sprintf("%d", i+1)
		tempData[i].Sub2 = fmt.Sprintf("CP492SG-A")
		tempData[i].Sub3 = fmt.Sprintf("YV2000-1003103SF1")
		tempData[i].Sub4 = fmt.Sprintf("子项3:%d", rand.Intn(1000000))
		tempData[i].Sub5 = fmt.Sprintf("子项4:%d", rand.Intn(1000000))
		tempData[i].Sub6 = fmt.Sprintf("子项5:%d", rand.Intn(1000000))
	}
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

func (f *TForm1) OnListView1Data(sender vcl.IObject, item *vcl.TListItem) {
	data := tempData[int(item.Index())]
	// 第一列为Caption属性所管理
	item.SetCaption(data.Sub1)
	item.SubItems().Add(data.Sub2)
	item.SubItems().Add(data.Sub3)
	item.SubItems().Add(data.Sub4)
	item.SubItems().Add(data.Sub5)
	item.SubItems().Add(data.Sub6)
}
