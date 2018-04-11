/*
* @Author: 01sr
* @Date:   2018-04-07 18:56:35
* @Last Modified by:   01sr
* @Last Modified time: 2018-04-11 17:22:21
 */
package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type OnlinesS struct {
	Device string `json:"device"`
	Type   int    `json:"type"` //设备下线可通过此值判断
	Time   string `json:"time"`
	Code   int    `json:"code"`
	BrasIp string `json:"brasIp"`
	WanIp  string `json:"wanIp"`
}
type WifiOnlinesS struct {
	Onlines []OnlinesS `json:"onlines"`
}
type OnlineDevice struct {
	Status      string       `json:"status"`
	WifiOnlines WifiOnlinesS `json:"wifiOnlines"`
}

type Account_namesS string
type UserS struct {
	Id                 string           `json:"id"`
	Name               string           `json:"name"`
	Mobile             string           `json:"mobile"`
	Vcard              string           `json:"vcard"`
	Avatar             string           `json:"avatar"`
	Savatar            string           `json:"savatar"`
	Time               string           `json:"time"`
	Update             string           `json:"update"`
	Code               string           `json:"code"`
	Country            string           `json:"country"`
	Continue_login_day string           `json:"continue_login_day"`
	Last_login_time    string           `json:"last_login_time"`
	Carrier            string           `json:"carrier"`
	Clear              string           `json:"clear"`
	City_id            int              `json:"city_id"`
	Admin_flag         string           `json:"admin_flag"`
	Did                string           `json:"did"`
	Sign               int              `json:"sign"`
	Account_names      []Account_namesS `json:"account_names"`
}
type LoginResult struct {
	Status string `json:"status"`
	User   UserS  `json:"user"`
}

type TelecomWifiResS struct {
	Password string `json:"password"`
	Code     int    `json:"code"`
	Redirect string `json:"redirect"`
}
type PasswdJson struct {
	Status         string          `json:"status"`
	TelecomWifiRes TelecomWifiResS `json:"telecomWifiRes"`
}

type QTelecomWifiResS struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Code     int    `json:"code"`
}
type QrcodeJson struct {
	Status         string           `json:"status"`
	TelecomWifiRes QTelecomWifiResS `json:"telecomWifiRes"`
}

type OnlineResult struct {
	Status   string `json:"status"`
	Response string `json:"response"`
}

type Mlog struct{}

var (
	mlog  = new(Mlog)
	debug *bool
)

func (*Mlog) toString(msg ...interface{}) string {
	s := fmt.Sprint(msg)
	return s[2 : len(s)-2]
}
func (this *Mlog) i(msg ...interface{}) {
	s := this.toString(msg)
	s = "[Info] " + s
	color.Green(s)
}

func (this *Mlog) d(msg ...interface{}) {
	if *debug {
		s := this.toString(msg)
		s = "[Debug] " + s
		color.Blue(s)
	}
}

func (this *Mlog) w(msg ...interface{}) {
	s := this.toString(msg)
	s = "[Warning] " + s
	color.Yellow(s)
}

func (this *Mlog) e(msg ...interface{}) {
	s := this.toString(msg)
	s = "[Error] " + s
	color.Red(s)
}

func listOnline(devices []OnlinesS) {
	if len(devices) == 0 {
		mlog.i("No device online.")
		exit()
	}
	s := "Devices of online:"
	for _, device := range devices {
		s += "\n" + fmt.Sprintf("%+v", device)
	}
	mlog.i(s)
}

func exit() {
	fmt.Println("Press any key to exit.")
	var e byte
	fmt.Scanf("%b", &e)
	os.Exit(0)
}
func encode(urlstring string) string {
	u, _ := url.Parse(urlstring)
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode
	return u.String()
}

func main() {
	account := flag.String("a", "", "The `account(phone number)` of ChinaTelecom(required!).")
	passwd := flag.String("p", "", "The `password` of '掌上大学'(required!).")
	ttype := flag.String("t", "1", "If your account support multiply devices, you can set it `0 or 1` to distinguish different devices.")
	behavior := flag.Int("b", 1, "Set `1 or 0` to login or log out.")
	list := flag.Bool("l", false, "List devices of online, can't use with -b together.")
	force := flag.Bool("f", false, "If your account is using by another device, make it offline forcedly.")
	debug = flag.Bool("d", false, "Enable debug mode.")
	hostname, err := os.Hostname()

	if err != nil {
		mlog.e(err.Error() + " set default name:\"default\"")
		hostname = "default"
	}
	name := flag.String("n", hostname, "The `device name`.")
	flag.Parse() //解析输入的参数
	if *account == "" || *passwd == "" {
		mlog.e("The -a [account] and the -p [password] must be set!\nUsing -h to see more.")
		exit()
	}
	mlog.i("accout: ", *account, ", password: ", *passwd, ", device name: ", *name)
	wanIp, brasIp, err := initial()
	if err != nil {
		mlog.e(err.Error())
		exit()
	}

	user, err := login(*account, *passwd)
	if err != nil {
		mlog.e(err.Error())
		exit()
	}
	if *list {
		devices, err := getOnlineDeviceList(user.Id, *account, *passwd)
		if err != nil {
			mlog.e(err.Error())
			exit()
		}
		listOnline(devices)
		exit()
	}

	if *behavior == 0 {
		// offline
		if wanIp != "0" {
			mlog.w("Already offline.")
			exit()
		}
		devices, err := getOnlineDeviceList(user.Id, *account, *passwd)

		if err != nil {
			mlog.e(err.Error())
			exit()
		}
		dd := fmt.Sprintf("%+v", devices)
		mlog.i(dd)
		if len(devices) == 0 {
			mlog.e("The current account does not match the login account.")
			exit()
		}
		for _, device := range devices {
			if strconv.Itoa(device.Type) == *ttype {
				err = kickOffDevice(user.Id, *account, *passwd, device.WanIp, device.BrasIp)
				if err != nil {
					mlog.e(err.Error())
					exit()
				}
				mlog.i("Log out successfully.")
			}
		}
	} else {
		// online
		if wanIp == "0" {
			mlog.w("Already online.")
			exit()
		}
		code, err := getPasswd(user.Id, *account, *passwd)
		if err != nil {
			mlog.e(err.Error())
			exit()
		}
		mlog.i(code)
		// 密码获取成功
		qrcode, err := getQrCode(wanIp, brasIp, *name)
		if err != nil {
			mlog.e(err.Error())
			exit()
		}
		mlog.i(qrcode)
		//qrcode获取成功
		err = online(user.Id, *account, *passwd, code, qrcode, *ttype)
		if err != nil && strings.Contains(err.Error(), "检测到你的帐号在其他设备登录") && *force {
			var devices []OnlinesS
			devices, err = getOnlineDeviceList(user.Id, *account, *passwd)
			if err != nil {
				mlog.e(err.Error())
				exit()
			}
			for _, device := range devices {
				if strconv.Itoa(device.Type) == *ttype {
					mlog.i("The account(type:" + *ttype + ") is using by \"" + device.Device + "\".")
					err = kickOffDevice(user.Id, *account, *passwd, device.WanIp, device.BrasIp)
					if err != nil {
						mlog.e(err.Error())
						exit()
					}
					mlog.i("Force \"" + device.Device + "\" offline successfully.")
					break
				}
			}
			time.Sleep(time.Second)
			err = online(user.Id, *account, *passwd, code, qrcode, *ttype)
			if err != nil {
				mlog.e("test " + err.Error())
			}
		}

		if err != nil {
			mlog.e(err.Error())
		} else {
			mlog.i("Login successfully.")
		}
	}
	exit()
}

func newClient(timeoutSecond time.Duration) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*timeoutSecond)
			return c, err
		},
		// MaxIdleConnsPerHost:   10,
		// ResponseHeaderTimeout: time.Second * 2,
	}
	client := &http.Client{Transport: tr}
	return client
}

// wanIp=="0"表示已登录校园网
func initial() (wanIp, brasIp string, err error) {
	defer func() {
		if r := recover(); r != nil {
			mlog.e(r)
		}
	}()
	u := encode("http://pre.f-young.cn/")
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", "", err
	}
	rep, err := newClient(10).Do(req)
	if err != nil {
		return "", "", err
	}
	if rep.StatusCode != 200 {
		return "", "", errors.New("Not the Telecom campus network.")
	}

	u = encode("HTTP://test.f-young.cn/")
	mlog.d("Access: " + u)
	req, err = http.NewRequest("GET", u, nil)
	if err != nil {
		return "", "", err
	}
	defer rep.Body.Close()
	rep, err = http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return "", "", err
	}
	defer rep.Body.Close()
	mlog.d("Response code: ", rep.StatusCode, "\nResponse header: ", rep.Header)
	if rep.StatusCode == 302 {
		content := rep.Header.Get("Location")
		mlog.d(content)
		argString := strings.Split(content, "?")
		args := strings.SplitN(argString[1], "&", -1)
		for _, param := range args {
			if strings.Contains(param, "wlanuserip") {
				wanIp = strings.Split(param, "=")[1]
			}
			if strings.Contains(param, "mscgip") {
				brasIp = strings.Split(param, "=")[1]
			}
		}
		return wanIp, brasIp, nil
	}
	if rep.StatusCode == 200 {
		return "0", "", nil
	}
	return "", "", errors.New("Failed to detect net state!")
}

func login(account, passwd string) (*UserS, error) {
	defer func() {
		if r := recover(); r != nil {
			mlog.e(r)
		}
	}()

	u := encode("https://www.loocha.com.cn:8443/login?1=Android_college_100.100.100")
	mlog.d("Access: " + u)
	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(account, passwd)
	response, err := newClient(0).Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	mlog.d("Response code: ", response.StatusCode)
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		mlog.d("Response body: ", string(body))
		loginResult := &LoginResult{}
		err = json.Unmarshal(body, loginResult)
		if err != nil {
			return nil, err
		}
		if loginResult.Status != "0" {
			return nil, errors.New("Failed to resolve user info![0].")
		}
		return &loginResult.User, nil
	}
	return nil, errors.New("Failed to resolve user info![1].")
}

func getPasswd(id, account, passwd string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			mlog.e(r)
		}
	}()
	u := encode("https://wifi.loocha.cn/" + id + "/wifi/telecom/pwd?type=4&1=Android_college_100.100.100")
	mlog.d("Access: " + u)
	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}
	request.SetBasicAuth(account, passwd)
	response, err := newClient(0).Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	mlog.d("Response code: ", response.StatusCode)
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		mlog.d("Response body: ", string(body))
		passwdJson := &PasswdJson{}
		err = json.Unmarshal(body, passwdJson)
		if err != nil {
			return "", err
		}
		code := passwdJson.TelecomWifiRes.Password
		if passwdJson.Status != "0" {
			if code == "0" || code == "" {
				return "", errors.New("Failed to get password![2]")
			}
			return "", errors.New(code)
		}
		return code, nil
	}
	return "", errors.New("Failed to get password![3]")
}

func getQrCode(ip, brasIp, name string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			mlog.e(r)
		}
	}()
	u := encode("https://wifi.loocha.cn/0/wifi/qrcode" + "?1=Android_college_100.100.100&brasip=" + brasIp + "&ulanip=" + ip + "&wlanip=" + ip + "&mm=" + name)
	mlog.d("Access: " + u)
	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}
	response, err := newClient(0).Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	mlog.d("Response code: ", response.StatusCode)
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		mlog.d("Response body: ", string(body))
		qrcodeJson := &QrcodeJson{}
		err = json.Unmarshal(body, qrcodeJson)
		if err != nil {
			return "", err
		}
		if qrcodeJson.Status != "0" {
			return "", errors.New("Failed to get qrcode![4]")
		}
		qrcode := qrcodeJson.TelecomWifiRes.Password
		return qrcode, nil
	}
	mlog.d(response.StatusCode)

	return "", errors.New("Failed to get qrcode![5]")
}

func online(id, account, passwd, code, qrcode, ttype string) error {
	defer func() {
		if r := recover(); r != nil {
			mlog.e(r)
		}
	}()
	param := "1=Android_college_100.100.100&qrcode=" + qrcode + "&code=" + code + "&type="
	param += ttype
	u := encode("https://wifi.loocha.cn/" + id + "/wifi/telecom/auto/login?" + param)
	mlog.d("Access: " + u)
	request, err := http.NewRequest("POST", u, nil)
	if err != nil {
		return err
	}
	request.SetBasicAuth(account, passwd)
	response, err := newClient(0).Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	mlog.d("Response code: ", response.StatusCode)
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		mlog.d("Response body: ", string(body))
		onlineResult := &OnlineResult{}
		err = json.Unmarshal(body, onlineResult)
		if err != nil {
			return err
		}
		status := onlineResult.Status
		if status != "0" {
			return errors.New(onlineResult.Response)
		}
		return nil
	}
	return errors.New("Failed to log out!")
}

func getOnlineDeviceList(id, account, passwd string) ([]OnlinesS, error) {
	defer func() {
		if r := recover(); r != nil {
			mlog.e(r)
		}
	}()
	u := encode("https://wifi.loocha.cn/" + id + "/wifi/status?1=Android_college_100.100.100")
	mlog.d("Access: " + u)
	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(account, passwd)
	if err != nil {
		return nil, err
	}
	response, err := newClient(0).Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	mlog.d("Response code: ", response.StatusCode)
	if response.StatusCode == http.StatusOK {
		onlineDevice := &OnlineDevice{}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		mlog.d("Response body: ", string(body))
		err = json.Unmarshal(body, onlineDevice)
		if err != nil {
			return nil, err
		}
		if onlineDevice.Status != "0" {
			return nil, errors.New("Error: " + onlineDevice.Status)
		}
		return onlineDevice.WifiOnlines.Onlines, nil
	}
	return nil, errors.New("Failed to get devices of online!")
}

func kickOffDevice(id, account, passwd, ip, brasIp string) error {
	defer func() {
		if r := recover(); r != nil {
			mlog.e(r)
		}
	}()
	request, err := http.NewRequest("DELETE", "https://wifi.loocha.cn/"+id+"/wifi/kickoff?1=Android_college_100.100.100&wanip="+ip+"&brasip="+brasIp, nil)
	if err != nil {
		return err
	}
	request.SetBasicAuth(account, passwd)
	response, err := newClient(0).Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if responseCode := response.StatusCode; responseCode == 200 {
		return nil
	}
	return errors.New("Failed to log out!")
}
