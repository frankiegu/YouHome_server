package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"crypto/tls"
)

var address string = "https://youhome.xyz"//"https://123.207.55.27"
var tr = &http.Transport{
	TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
}
var client = &http.Client{Transport: tr}
func deleteScene(id string){
	data := struct {
		SceneId string `json:"sceneId"`
	}{id}
	buf , _ := json.Marshal(data)///scene/delete
	//fmt.Println(string(buf))
	//resp, err = http.Post(address + "/scene/delete","application/json",bytes.NewBuffer(tempJson))
	resp, err := client.Post(address+"/v1/scenes/delete","application/json",bytes.NewBuffer(buf))
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func login(code string){
	user := struct{
		Code string `json:"code"`
	}{code}
	//fmt.Println(string(user))
	userJson ,err := json.Marshal(user)
	fmt.Println(string(userJson))
	checkErr(err)
	resp, err := client.Post(address+"/v1/users","application/json",bytes.NewBuffer(userJson))
	checkErr(err)
	body,err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func createScene(userId,sceneName string){
	scene := struct {
		UserId string `json:"userId"`
		SceneName string `json:"sceneName"`
	}{userId,sceneName}
	sceneJson,err := json.Marshal(scene)
	checkErr(err)
	resp, err := client.Post(address+"/v1/scenes","application/json",bytes.NewBuffer(sceneJson))
	checkErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func getAllScene(userId string){
	resp, err := client.Get(address + "/v1/scenes?userId="+userId)
	checkErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func getUserInfo(userId string){
	url := address+"/v1/users?userId="+userId
	//fmt.Println(url)
	resp, err := client.Get(url)
	checkErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func getDeviceOfScene(sceneId string){
	//device/all?sceneId=room
	res,err := client.Get(address + "/v1/devices?sceneId=" + sceneId)
	checkErr(err)

	body,err := ioutil.ReadAll(res.Body)
	checkErr(err)

	fmt.Println(string(body))
}
func updateUserInfo(userId,userName string){
	temp := struct {
		UserId string `json:"userId"`
		UserName string `json:"userName"`
	}{userId,userName}
	buf,err := json.Marshal(temp)
	checkErr(err)
	resp, err := client.Post(address + "/v1/users/userName","application/json",bytes.NewBuffer(buf))
	checkErr(err)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func checkErr(err error){
	if err != nil{
		panic(err)
	}
}
func getDeviceState(deviceId string){
	res,err := client.Get(address + "/v1/devices/states?deviceId=" + deviceId)
	checkErr(err)

	body,err := ioutil.ReadAll(res.Body)
	checkErr(err)

	fmt.Println(string(body))
}
func updateDeviceState(deviceId, operation string){
	temp := struct {
		DeviceId string `json:"deviceId"`
		Operation string `json:"operation"`
	}{deviceId,operation}
	buf,err := json.Marshal(temp)
	checkErr(err)
	resp, err := client.Post(address + "/v1/devices/states","application/json",bytes.NewBuffer(buf))
	checkErr(err)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func getDeviceName(deviceId string){
	//	/v1/devices/devicename?deviceId=abc
	res,err := client.Get(address + "/v1/devices/devicename?deviceId=" + deviceId)
	checkErr(err)

	body,err := ioutil.ReadAll(res.Body)
	checkErr(err)

	fmt.Println(string(body))
}
func updateDeviceName(deviceId,deviceName string){
	temp := struct {
		DeviceId string `json:"deviceId"`
		DeviceName string `json:"deviceName"`
	}{deviceId,deviceName}
	buf,err := json.Marshal(temp)
	checkErr(err)
	resp, err := client.Post(address + "/v1/devices/devicename","application/json",bytes.NewBuffer(buf))
	checkErr(err)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func TestWechatApi(){
	//https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=wx394f9cc0f949d50b&secret=0346967483b301c189254bef576b1091&js_code=JSCODE&grant_type=authorization_code")
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func main() {
	updateDeviceState("4","turn_off")
	getDeviceState("4")
}