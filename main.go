package main

import (
	"fmt"
	"net/http"
	"time"
)

var website = []string{"/members", "/news_admin", "/guanli", "/newsadmin", "/fuck", "/manage", "/Sql", "/GuanLi", "/File", "/main", "/login", "/Sys_Login", "/manage", "/houtai", "/users", "/2006", "/Test", "/user_login", "/Mgr", "/User_Admin", "/Fuck", "/root", "/News_Admin", "/2019", "/Admin_Bak", "/2008", "/sqlconf", "/2001", "/user_admin", "/hello", "/myadmin", "/admin2", "/2007", "/Login", "/web2", "/admin_login", "/2005", "/Manage", "/Conf", "/sys", "/Upload", "/Hello", "/mgr", "/System", "/Site", "/Web", "/sys_login", "/web1", "/Administrator", "/Sys", "/test", "/MyAdmin", "/administrator", "/backup", "/Master", "/system", "/Web2", "/admin1", "/Root", "/ok", "/auto", "/file", "/log", "/master", "/Admin2", "/Auto", "/admin_bak", "/sql", "/adm", "/conf", "/Admin1", "/admin_user", "/Manager_Login", "/2002", "/manager_login", "/HouTai", "/config", "/Admin_Login", "/web", "/SqlConf", "/Members", "/Log", "/site", "/user", "/User", "/admin", "/Users", "/User_Login", "/upload", "/Admin", "/Adm", "/Web1", "/BackUp", "/NewsAdmin", "/check", "/Config", "/other", "/Main", "/inc", "/Inc", "/add", "/administrator", "/web_master", "/User", "/login", "/Users", "/Manage", "/Config", "/System", "/Administrator", "/config", "/admin1", "/Upload", "/system", "/log", "/NewsAdmin", "/Admin_user", "/Check", "/fuck", "/Fuck", "/register", "/user_admin", "/news_admin", "/2005", "/user_login", "/Test", "/houtai", "/administration", "/User_Login", "/mgr", "/Add", "/myadmin", "/Reg", "/Admin_Login", "/Root", "/desk", "/Login", "/sys", "/User_Admin", "/Local", "/MyAdmin", "/admin_login", "/2002", "/Manager_Login", "/user", "/test", "/Sys_Login", "/Sys", "/users", "/2001", "/News_Admin1404", "/Admin2", "/admin2", "/Mgr", "/checklogin", "/Edit", "/2003", "/newsadmin", "/Admin1", "/backup", "/edit", "/guanli", "/root", "/2007", "/Adm", "/Log", "/adm", "/master", "/upload", "/Web_Master", "/Master", "/BackUp", "/2008", "/check", "/local", "/CheckLogin", "/memberlogin", "/WebMaster", "/Administration", "/MemberLogin", "/Manager", "/reg", "/manage", "/Register", "/Conf", "/Admin", "/sys_login", "/conf", "/webmaster", "/Ok", "/Check", "/Other", "/Member", "/Web2", "/admin_bak", "/members", "/sql", "/Auto", "/auto", "/sqlconf", "/File", "/Admin_Bak", "/Web", "/site", "/Site", "/hello", "/GuanLi", "/HouTai", "/Web1", "/Check", "/Sql", "/Ok", "/web1", "/Main", "/include", "/check", "/Include", "/check", "/other", "/Reg", "/web", "/member", "/Administration", "/local", "/file", "/main", "/CheckLogin", "/checklogin", "/Edit", "/web2", "/Add", "/admin_user", "/Members", "/Local", "/Inc", "/Hello", "/Check", "/reg", "/inc", "/web_master", "/Admin_user", "/ok", "/Register", "/Other", "/webmaster", "/Web_Master", "/register", "/administration", "/edit", "/memberlogin", "/add", "/MemberLogin", "/SqlConf", "/WebMaster"}
var http200 []string

func show() {
	fmt.Println("欢迎使用RealSpeak Urltool 0.1")
	fmt.Println("本产品由ROGER~REALSPEAK开发")
	fmt.Println("目前本产品支持3个命令")
	fmt.Println("使用\"time-\"，可以获得当前日期时间")
	fmt.Println("使用\"read- \" + url，可以获得目标url的源代码")
	fmt.Println("使用\"scan- \" + url，可以使用产品自带字典，来穷举目标url可用的子链接")
	fmt.Printf("RealSpeak Urltool V0.02 © 2019 ALL RIGHTS RESERVED\n")
	fmt.Printf("HELP HAS BEEN ALREADY SHOWN \n")
	catch()
}
func catch() {
	var preCommand string
	var realCommand string
	fmt.Printf("\n")
	fmt.Printf(">>:")
	for preCommand == "" {
		fmt.Scanf("%s %s", &preCommand, &realCommand)
	}
	switch {
	case preCommand == "read-" && realCommand != "":
		httpVisit(realCommand)
	case preCommand == "scan-" && realCommand != "":
		httpGet(realCommand)
	case preCommand == "help-" && realCommand == "":
		show()
	case preCommand == "time-" && realCommand == "":
		now := time.Now()
		fmt.Println(now.Format("2006/01/02 15:04:05"))
		catch()
	case preCommand == "roger":
		fmt.Println("ROGER告诉你，闷声发大财！")
		catch()
	default:
		if preCommand != "" {
			fmt.Println("无效的command，请重新输入！")
		}
		errAlert()
	}
}
func errAlert() {
	catch()
}
func httpGet(httpUrl string) {
	http200 = append(http200[0:0])
	if httpUrl != "" {
		for _, url := range website {
			var httpUrler = httpUrl + url
			httpScan(httpUrler)
		}
		fmt.Println("* * * * * * * * *")
		fmt.Println("已检测的有效地址：")
		if len(http200) == 0 {
			fmt.Println("暂未检测出有效地址")
		} else {
			for _, url := range http200 {
				fmt.Println(url)
			}
		}
		fmt.Println("* * * * * * * * *")
		catch()
	}
}
func httpScan(httpUrl string) {
	if httpUrl != "" {
		resp, err := http.Get(httpUrl)
		if err != nil {
			fmt.Printf("目标url无法访问，请重新输入：\n")
			catch()
		} else {
			fmt.Printf("HTTP %d FOR %s\n", resp.StatusCode, httpUrl)
			if resp.StatusCode == 200 {
				http200 = append(http200, httpUrl)
			}
		}
	}
}
func httpVisit(httpUrl string) {
	resp, err := http.Get(httpUrl)
	if err != nil {
		fmt.Println("目标url无法访问，请重新输入：")
	} else {
		if resp.StatusCode == http.StatusOK {
			//fmt.Println(resp.StatusCode)
			buf := make([]byte, 1024)
			for {
				n, _ := resp.Body.Read(buf)
				if 0 == n {
					break
				}
				fmt.Println(string(buf[:n]))
			}
		}
	}
	catch()
}
func main() {
	fmt.Printf("RealSpeak Urltool V0.02 © 2019 ALL RIGHTS RESERVED\n")
	fmt.Printf("TIPS: 输入\"help-\"查看帮助\n")
	catch()

}
