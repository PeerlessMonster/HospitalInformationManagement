package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	DBUtil "hospital/conn"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/*")
	router.StaticFS("/static", http.Dir("static"))

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	router.GET("/register", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "register.html", nil)
	})
	router.POST("/register", checkRegister)
	router.POST("/info/get/:username", sendPersonalInfo)
	router.POST("/info/change/:username", changePersonalInfo)
	router.POST("/password/check/:name", beforeChangePassword)
	router.POST("/password/change/:name", changePassword)
	router.POST("/index", checkLogin)
	router.GET("/registrate/:username", registrate)
	router.POST("/registrate/:patientName/:doctorName", procRegistrate)
	router.GET("/visit_record/:patientName", procVisit)
	router.POST("/visit_record/post/:listNo", procMedicinelist)
	router.GET("/hospital_record/:patientName", procHospital)
	router.GET("/empty", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "empty.html", nil)
	})

	router.Run(":6660")
}

func judgeRole(role int) (tableName, colName string) {
	if role == 0 {
		tableName = "patient"
		colName = "patientPhone"
	} else if role == 1 {
		tableName = "doctor"
		colName = "employeeNo"
	} else if role == 2 {
		tableName = "admin"
		colName = "employeeNo"
	}
	return
}

func checkRegister(ctx *gin.Context) {
	phone := ctx.PostForm("username")
	db := DBUtil.GetInstance()
	rows, err := db.Raw("SELECT * FROM patient WHERE patientPhone=?", phone).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		back(ctx, "该用户已注册！", "register.html")
		return
	} else {
		var lastNo string
		db.Raw("SELECT patientNo FROM patient ORDER BY patientNo DESC LIMIT 1").Scan(&lastNo)

		noInfo, err := strconv.Atoi(lastNo[2:])
		if err != nil {
			panic(err)
		}
		noInfo = noInfo + 1
		no := fmt.Sprintf("%07d", noInfo)
		no = "PA" + no
		fmt.Println(no)
		password := ctx.PostForm("password")
		name := ctx.PostForm("name")
		sexInfo := ctx.PostForm("sex")
		sexSign, err := strconv.Atoi(sexInfo)
		if err != nil {
			panic(err)
		}
		var sex string
		if sexSign == 0 {
			sex = "男"
		} else if sexSign == 1 {
			sex = "女"
		}
		address := ctx.PostForm("address")
		if address == "" {
			db.Exec("INSERT INTO patient VALUES(?, ?, ?, ?, ?, NULL)", no, password, name, sex, phone)
		} else {
			db.Exec("INSERT INTO patient VALUES(?, ?, ?, ?, ?, ?)", no, password, name, sex, phone, address)
		}

		back(ctx, "注册成功！点击确定去登录", "login.html")
	}
}

func checkLogin(ctx *gin.Context) {
	roleInfo := ctx.PostForm("role")
	role, err := strconv.Atoi(roleInfo)
	if err != nil {
		panic(err)
	}
	username := ctx.PostForm("username")

	rightPassword := checkPassword(username, role)
	if rightPassword == "" {
		back(ctx, "用户名不存在！", "login.html")
		return
	}
	password := ctx.PostForm("password")
	if password != rightPassword {
		back(ctx, "密码错误！", "login.html")
		return
	}

	var colname2 string
	if role == 0 {
		colname2 = "patientName"
	} else if role == 1 {
		colname2 = "doctorName"
	}
	var name string
	if role == 2 {
		name = "admin"
	} else if role == 0 || role == 1 {
		tableName, colName := judgeRole(role)
		sql := "SELECT " + colname2 + " FROM " + tableName + " WHERE " + colName + " = ?"
		db := DBUtil.GetInstance()
		db.Raw(sql, username).Scan(&name)
	}

	var funs []string
	if role == 0 {
		funs = []string{"挂号", "就诊记录", "住院记录"}
	} else if role == 1 {
		funs = []string{"排班", "坐诊", "住院治疗"}
	} else if role == 2 {
		funs = []string{"门诊部", "住院部", "药房"}
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"username": name,
		"role":     role,
		"funs":     funs,
	})
}

func checkPassword(username string, role int) (rightPassword string) {
	tableName, colName := judgeRole(role)

	db := DBUtil.GetInstance()
	sql := "SELECT password FROM " + tableName + " WHERE " + colName + " = ?"
	db.Raw(sql, username).Scan(&rightPassword)

	return
}

type UserInfo struct {
	Password string `json:"password"`
	RoleInfo string `json:"role"`
}

func beforeChangePassword(ctx *gin.Context) {
	name := ctx.Param("name")
	var userInfo UserInfo
	err := ctx.BindJSON(&userInfo)
	if err != nil {
		panic(err)
	}
	role, err := strconv.Atoi(userInfo.RoleInfo)
	if err != nil {
		panic(err)
	}

	var username string
	if role == 0 {
		username = DBUtil.Select_patientPhone_by_patientName(name)
	} else if role == 1 {
		username = DBUtil.Select_doctorNo_by_doctorName(name)
	}

	if checkPassword(username, role) != userInfo.Password {
		ctx.Writer.Write([]byte("密码错误！"))
	}
}

func changePassword(ctx *gin.Context) {
	name := ctx.Param("name")
	var userInfo UserInfo
	err := ctx.BindJSON(&userInfo)
	if err != nil {
		panic(err)
	}
	role, err := strconv.Atoi(userInfo.RoleInfo)
	if err != nil {
		panic(err)
	}

	var username string
	if role == 0 {
		username = DBUtil.Select_patientPhone_by_patientName(name)
	} else if role == 1 {
		username = DBUtil.Select_doctorNo_by_doctorName(name)
	}

	tableName, colName := judgeRole(role)
	sql := "UPDATE " + tableName + " SET password = ? WHERE " + colName + " = ?"
	db := DBUtil.GetInstance()
	db.Exec(sql, userInfo.Password, username)

	ctx.Writer.Write([]byte("修改成功！"))
}

func sendPersonalInfo(ctx *gin.Context) {
	username := ctx.Param("username")
	patientNo := DBUtil.Select_patientNo_by_patientName(username)

	db := DBUtil.GetInstance()
	sql := "SELECT patientPhone, patientAddress FROM patient WHERE patientNo = ?"
	row := db.Raw(sql, patientNo).Row()

	var phone, address string
	row.Scan(&phone, &address)
	ctx.JSON(http.StatusOK, gin.H{
		"phone":   phone,
		"address": address,
	})
}

func changePersonalInfo(ctx *gin.Context) {
	username := ctx.Param("username")
	var patientInfo struct {
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}
	err := ctx.BindJSON(&patientInfo)
	if err != nil {
		panic(err)
	}

	db := DBUtil.GetInstance()
	if patientInfo.Phone != "" {
		sql := "SELECT * FROM patient WHERE patientPhone = ?"
		row := db.Raw(sql, username).Row()
		if row != nil {
			ctx.Writer.Write([]byte("该手机号已被注册！"))
			return
		}

		sql2 := "UPDATE patient SET patientPhone = ? WHERE patientPhone = ?"
		db.Exec(sql2, patientInfo.Phone, username)
	}
	if patientInfo.Address != "" {
		sql := "UPDATE patient SET patientAddress = ? WHERE patientPhone = ?"
		db.Exec(sql, patientInfo.Address, username)
	}
	ctx.Writer.Write([]byte("保存成功！"))
}

func registrate(ctx *gin.Context) {
	db := DBUtil.GetInstance()
	sql := "SELECT doctorName, sex, department, staff FROM doctor WHERE employeeNo IN (SELECT employeeNo FROM work WHERE section='门诊部')"
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var doctors [][]string
	for rows.Next() {
		var doctor [4]string
		rows.Scan(&doctor[0], &doctor[1], &doctor[2], &doctor[3])
		doctors = append(doctors, doctor[:])
	}

	patientName := ctx.Param("username")
	ctx.HTML(http.StatusOK, "registrate.html", gin.H{
		"doctors":     doctors,
		"patientName": patientName,
	})
}

func procRegistrate(ctx *gin.Context) {
	doctorName := ctx.Param("doctorName")
	db := DBUtil.GetInstance()
	doctorNo := DBUtil.Select_doctorNo_by_doctorName(doctorName)

	patientName := ctx.Param("patientName")
	patientNo := DBUtil.Select_patientNo_by_patientName(patientName)

	sql3 := "INSERT INTO REGISTRATE VALUE (NULL, ?, ?)"
	db.Exec(sql3, patientNo, doctorNo)

	sql4 := "SELECT theOrder FROM registrate ORDER BY theOrder DESC LIMIT 1"
	row3 := db.Raw(sql4).Row()
	var order int
	row3.Scan(&order)

	orderInfo := strconv.Itoa(order)
	ctx.Writer.Write([]byte(orderInfo))
}

func procVisit(ctx *gin.Context) {
	patientName := ctx.Param("patientName")
	patientNo := DBUtil.Select_patientNo_by_patientName(patientName)
	db := DBUtil.GetInstance()
	sql := "SELECT date, doctorName, department, staff, symptom, listNo FROM visit_record v, doctor d WHERE v.employeeNo=d.employeeNo AND patientNo = ?"
	rows, err := db.Raw(sql, patientNo).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var records [][]string
	for rows.Next() {
		var record [6]string
		rows.Scan(&record[0], &record[1], &record[2], &record[3], &record[4], &record[5])
		records = append(records, record[:])
	}

	ctx.HTML(http.StatusOK, "visit_record.html", gin.H{
		"patientName": patientName,
		"records":     records,
	})
}

func procMedicinelist(ctx *gin.Context) {
	listNo := ctx.Param("listNo")
	fmt.Println(listNo)

	db := DBUtil.GetInstance()
	sql := "SELECT medicineName, medicinePrice, count, method FROM list l, medicine m WHERE l.medicineNo=m.medicineNo AND listNo = ?"
	rows, err := db.Raw(sql, listNo).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	listTable := "<table><tr><th>药品名</th><th>单价</th><th>数量</th><th>用法</th></tr>"
	var sum int = 0
	var totalPrice float32 = 0
	for rows.Next() {
		var medicineName, method string
		var medicinePrice float32
		var count int
		rows.Scan(&medicineName, &medicinePrice, &count, &method)
		priceInfo := fmt.Sprintf("%.2f", medicinePrice)
		countInfo := strconv.Itoa(count)

		sum += count
		totalPrice += float32(count) * medicinePrice

		oneLine := "<tr><td>" + medicineName + "</td><td>" + priceInfo + "</td><td class='count'>" + countInfo + "</td><td>" + method + "</td></tr>"
		listTable += oneLine
	}

	sumInfo := strconv.Itoa(sum)
	totalPriceInfo := fmt.Sprintf("%.2f", totalPrice)
	listTable += "<tr><td colspan='4'>&nbsp;</td></tr><tr><th>总数量</th><td>" + sumInfo + "</td><th>总价</th><td>" + totalPriceInfo + "</td></tr></table>"
	fmt.Println(listTable)
	ctx.Writer.Write([]byte(listTable))
}

func procHospital(ctx *gin.Context) {
	patientName := ctx.Param("patientName")
	patientNo := DBUtil.Select_patientNo_by_patientName(patientName)
	db := DBUtil.GetInstance()
	sql := "SELECT balance, inDate, h.roomNo, location, roomPrice, r.department, bedNo, doctorName, sex, d.department, staff, fileNo, outDate FROM hospital_file h, room r, doctor d WHERE h.roomNo=r.roomNo AND h.employeeNo=d.employeeNo AND patientNo = ?"
	row := db.Raw(sql, patientNo).Row()
	if row == nil {
		ctx.HTML(http.StatusOK, "hospital_record.html", gin.H{
			"patientName": patientName,
			"records":     nil,
		})
		return
	}

	var balance float32
	var inDate string
	var outDate string // 如果是NULL，后面的都读不到了，疑似源码bug
	var roomNo string
	var location string
	var roomPrice float32
	var roomDepart string
	var bedNo string
	var doctorName string
	var sex string
	var doctorDepart string
	var staff string
	var fileNo string
	row.Scan(&balance, &inDate, &roomNo, &location, &roomPrice, &roomDepart, &bedNo, &doctorName, &sex, &doctorDepart, &staff, &fileNo, &outDate)

	sql2 := "SELECT date, symptom, listNo FROM hospital_record WHERE fileNo = ?"
	rows, err := db.Raw(sql2, fileNo).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var records [][]string
	for rows.Next() {
		var record [3]string
		rows.Scan(&record[0], &record[1], &record[2])
		records = append(records, record[:])
	}

	ctx.HTML(http.StatusOK, "hospital_record.html", gin.H{
		"patientName":  patientName,
		"balance":      balance,
		"inDate":       inDate,
		"outDate":      outDate,
		"roomNo":       roomNo,
		"location":     location,
		"roomPrice":    roomPrice,
		"roomDepart":   roomDepart,
		"bedNo":        bedNo,
		"doctorName":   doctorName,
		"sex":          sex,
		"doctorDepart": doctorDepart,
		"staff":        staff,
		"records":      records,
	})
}

func back(ctx *gin.Context, message string, path string) {
	js := "<script>alert('" + message + "')</script>"
	ctx.Writer.Write([]byte(js))
	ctx.HTML(http.StatusOK, path, nil)
}
