package models

// Device 设备信息
type Device struct {
	ID int64 `json:"id" gorm:"id" mapstructure:"id"`
	//Deleted           int64     `json:"deleted" gorm:"deleted" mapstructure:"deleted"`                               //  是否删除  0-否 1-是
	CreateTime string `json:"createTime" gorm:"create_time" mapstructure:"create_time"` //UpdateTime time.Time `json:"updateTime" gorm:"update_time" mapstructure:"update_time" time_format:"2025-06-20 05:10:40"` // 更新时间
	//Version           int64     `json:"version" gorm:"version" mapstructure:"version"`                               // 版本号
	SerialNo   string `json:"serialNo" gorm:"serial_no" mapstructure:"serial_no"`       // 序列号
	AdminId    int64  `json:"adminId" gorm:"admin_id" mapstructure:"admin_id"`          // 添加人ID
	AdminName  string `json:"adminName" gorm:"admin_name" mapstructure:"admin_name"`    // 添加人名称
	AdminPid   int64  `json:"adminPid" gorm:"admin_pid" mapstructure:"admin_pid"`       // 添加人上级ID
	AgencyId   int64  `json:"agencyId" gorm:"agency_id" mapstructure:"agency_id"`       // 所属代理ID
	MerchantId int64  `json:"merchantId" gorm:"merchant_id" mapstructure:"merchant_id"` // 所属商户ID
	DeviceName string `json:"deviceName" gorm:"device_name" mapstructure:"device_name"` // 设备型号
	UserName   string `json:"userName" gorm:"user_name" mapstructure:"user_name"`       // 用户姓名
	UserPhone  string `json:"UserPhone" gorm:"user_phone" mapstructure:"user_phone"`    // 用户手机号
	Udid       string `json:"Udid" gorm:"udid" mapstructure:"udid"`                     // 设备uuid
	// LastConnectTime time.Time `json:"LastConnectTime" gorm:"last_connect_time" mapstructure:"last_connect_time" time_format:"2006-01-02T15:04:05"` // 上次通信时间
	MdmStatus     int64 `json:"MdmStatus" gorm:"mdm_status" mapstructure:"mdm_status"`             // mdm服务器注册状态:0未注册，1已注册
	DeviceStatus  int64 `json:"DeviceStatus" gorm:"device_status" mapstructure:"device_status"`    // 设备状态:0正常
	FactoryStatus int64 `json:"FactoryStatus" gorm:"factory_status" mapstructure:"factory_status"` // 恢复出厂限制:0未限制，1已限制
	UsbStatus     int64 `json:"UsbStatus" gorm:"usb_status" mapstructure:"usb_status"`             // USB限制:0未限制，1已限制
	ActiveStatus  int64 `json:"ActiveStatus" gorm:"active_status" mapstructure:"active_status"`    // 激活锁状态:0未上锁，1已上锁
	LockStatus    int64 `json:"LockStatus" gorm:"lock_status" mapstructure:"lock_status"`          // 锁定状态:0未丢失，1已丢失
	AbmStatus     int64 `json:"AbmStatus" gorm:"abm_status" mapstructure:"abm_status"`             // ABM状态:0不在库，1在库
	//Longitude         string    `json:"longitude" gorm:"longitude" mapstructure:"longitude"`                         // 经度
	//Latitude          string    `json:"latitude" gorm:"latitude" mapstructure:"latitude"`                            // 纬度
	ProductName       string `json:"ProductName" gorm:"product_name" mapstructure:"product_name"`
	Imei              string `json:"imei" gorm:"imei" mapstructure:"imei"`
	Imei2             string `json:"imei2" gorm:"imei2" mapstructure:"imei2"`
	Meid              string `json:"meid" gorm:"meid" mapstructure:"meid"`
	PhoneNumberRecord string `json:"PhoneNumberRecord" gorm:"phone_number_record" mapstructure:"phone_number_record"` // 插卡记录
	OsVersion         string `json:"OsVersion" gorm:"os_version" mapstructure:"os_version"`
	//RestrictItems     string    `json:"RestrictItems" gorm:"restrict_items" mapstructure:"restrict_items"` // 功能限制条目
	ActiveBypass string `json:"ActiveBypass" gorm:"active_bypass" mapstructure:"active_bypass"` // 激活锁绕过码
	//UnlockToken       string    `json:"unlock_token" gorm:"unlock_token" mapstructure:"unlock_token"`       // 解锁token
	IpLocation  string `json:"IpLocation" gorm:"ip_location" mapstructure:"ip_location"`    // Ip归属地
	ProfileUuid string `json:"ProfileUuid" gorm:"profile_uuid" mapstructure:"profile_uuid"` // DEP自动分配文件UUID
	//BatteryLevel int64  `json:"battery_level" gorm:"battery_level" mapstructure:"battery_level"` // 电量
	//NetType      int64  `json:"net_type" gorm:"net_type" mapstructure:"net_type"`
	//AssistCode   string `json:"assist_code" gorm:"assist_code" mapstructure:"assist_code"` // 监管助手激活码
}

// TableName 表名称
func (*Device) TableName() string {
	return "device"
}
