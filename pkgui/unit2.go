// 由res2go IDE插件自动生成，不要编辑。
package pkgui

import (
    "github.com/ying32/govcl/vcl"
)

type TForm2 struct {
    *vcl.TForm
    Label1 *vcl.TLabel

    //::private::
    TForm2Fields
}

var Form2 *TForm2




// vcl.Application.CreateForm(&Form2)

func NewForm2(owner vcl.IComponent) (root *TForm2)  {
    vcl.CreateResForm(owner, &root)
    return
}

var form2Bytes = []byte("\x54\x50\x46\x30\x06\x54\x46\x6F\x72\x6D\x32\x05\x46\x6F\x72\x6D\x32\x04\x4C\x65\x66\x74\x03\x60\x01\x06\x48\x65\x69\x67\x68\x74\x03\xF0\x00\x03\x54\x6F\x70\x02\x21\x05\x57\x69\x64\x74\x68\x03\x40\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x85\xB3\xE4\xBA\x8E\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xF0\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x40\x01\x0D\x44\x65\x73\x69\x67\x6E\x54\x69\x6D\x65\x50\x50\x49\x02\x78\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x68\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x60\x05\x57\x69\x64\x74\x68\x02\x5A\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x12\xE4\xBD\x9C\xE8\x80\x85\xEF\xBC\x9A\xE6\xA2\x81\xE6\x96\x87\xE6\x9D\xB0\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x00")

// 注册Form资源  
var _ = vcl.RegisterFormResource(Form2, &form2Bytes)
