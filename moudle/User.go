package moudle

import "time"

//Owner 车主账号信息
type Owner struct {
	ID          uint      `json:"id"`           //数据库内ID
	IDCard      uint      `json:"idcard"`       //车主身份证号码
	Name        string    `json:"name"`         //车主真实姓名
	UserName    string    `json:"username"`     //账号使用昵称
	PassWord    string    `json:"password"`     //账号密码
	CreateTime  time.Time `json:"create_time"`  //创立时间
	PhoneNumber uint      `json:"phone_number"` //车主电话号码
	OpenID      uint      `json:"open_id"`      //用户识别码
}

//Admin 管理员账号信息
type Admin struct {
	ID       uint   `json:"id"`       //数据库内ID
	UserName string `json:"username"` //账号使用昵称
	PassWord string `json:"password"` //账号密码
}

//Auto 车辆信息
type Auto struct {
	ID      uint `json:"id"`       //数据库内ID
	CARRFID uint `json:"car_rfid"` //RFID
	OwnerID
	LicenseNumber string    `json:"lincensenumber"` //车牌号
	AutoType      string    `json:"auto_type"`      //车类型
	AutoColour    string    `json:"auto_colour"`    //车颜色
	AutoTime      time.Time `json:"auto_time"`      //车使用时间
	Flag          bool      `json:"flag"`           //是否在用标志位
}

//RFID RFID基站数据
type RFID struct {
}
