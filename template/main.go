package main

import "fmt"

// 模版方法模式

/*
让我们来考虑一个一次性密码功能 （OTP） 的例子。 将 OTP 传递给用户的方式多种多样 （短信、 邮件等）。 但无论是短信还是邮件， 整个 OTP 流程都是相同的：

	生成随机的 n 位数字。
	在缓存中保存这组数字以便进行后续验证。
	准备内容。
	发送通知。

后续引入的任何新 OTP 类型都很有可能需要进行相同的上述步骤。

因此， 我们会有这样的一个场景， 其中某个特定操作的步骤是相同的， 但实现方式却可能有所不同。 这正是适合考虑使用模板方法模式的情况。

首先， 我们定义一个由固定数量的方法组成的基础模板算法。 这就是我们的模板方法。 然后我们将实现每一个步骤方法， 但不会改变模板方法。
*/

func main() {
	// otp := otp{}

	// smsOTP := &sms{
	//  otp: otp,
	// }

	// smsOTP.genAndSendOTP(smsOTP, 4)

	// emailOTP := &email{
	//  otp: otp,
	// }
	// emailOTP.genAndSendOTP(emailOTP, 4)
	// fmt.Scanln()
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)

}
