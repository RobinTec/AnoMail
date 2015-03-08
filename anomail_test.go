package email

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	//__send_mail_test()
	//_private_test()
	//__add_address_test()
	//send_simple_test()
	send_with_bcc_test()
}

//=========Test Function========================
func send_simple_test() {
	err := SendSimple("测试SendSimple", "sendsimple@app.xiaomi.com", []string{"gaojiasheng@xiaomi.com", "gaojiasheng.him@foxmail.com"}, "我是内容：SendSimpleTest")
	fmt.Println(err)
}

func send_with_cc_test() {
	err := SendWithCc("测试SendWithCc", "SendWithCc@app.xiaomi.com", []string{"gaojiasheng@xiaomi.com"}, []string{"gaojiasheng.him@foxmail.com"}, "我是内容：SendWithCc")
	fmt.Println(err)
}
func send_with_bcc_test() {
	err := SendWithBcc("测试SendWithBcc", "SendWithBcc@app.xiaomi.com", []string{"gaojiasheng.him@foxmail.com"}, []string{}, []string{"gaojiasheng@xiaomi.com"}, "我是内容：SendWithBcc")
	fmt.Println(err)
}

//=========Test Base Class======================
func _add_address_test() {
	receiver := []string{"gaojiasheng", "chenzijun@xiaomi.com"}
	ret := __add_mail_postfix(receiver, "@xiaomi.com")
	fmt.Println(ret)
}

func _private_test() {
	Mail := New("我是标题", "memeda@app.xiaomi.com", []string{"gaojiasheng@xiaomi.com"})
	Mail.SetBcc([]string{"gaojiasheng@xiaomi.com"})
	err := Mail.Send("你好，我是一封新邮件！<hr/>我高家升的foxmail邮箱，密送给了高家升小米邮箱,试一下^_^<hr/>如果你看到我就说明你MIME没白看，嗯！<h1>我支持html哦</h1><h2 style='color:red'>我还有颜色呢</h2><a href='http://www.baidu.com'>我还能跳转呢～</a>")
	fmt.Println(err)
}
