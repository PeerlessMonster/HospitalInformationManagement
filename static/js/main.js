let username = document.querySelector(".dropdown a strong").innerText

let iframe = document.querySelector("iframe")

let liObjs = document.querySelectorAll(".fun")
function sidebarSelectControl() {
    for (let i = 0; i < liObjs.length; i++) {
        aObj = liObjs[i].children[0]
        aObj.classList.remove("active")
        aObj.classList.add("link-dark")
    }
    let thisAObj = this.children[0]
    thisAObj.classList.remove("link-dark")
    thisAObj.classList.add("active")

    let thisFun = thisAObj.querySelector("span").innerText
    let path
    switch (thisFun) {
        case "挂号":
            path = "registrate/" + username
            break;
        case "就诊记录":
            path = "visit_record/" + username
            break
        case "住院记录":
            path = "hospital_record/" + username
            break;
        case "排班":
            path = "work/" + username
            break
        case "坐诊":
            path = "do_visit/" + username
            break
        case "住院治疗":
            path = "do_hospital/" + username
            break
        case "门诊部":
            path = "visit_depart"
            break
        case "住院部":
            path = "hospital_depart"
            break
        case "药房":
            path = "pharmacy"
            break
    }
    iframe.setAttribute("src", path)
}

for (let i = 0; i < liObjs.length; i++) {
    liObjs[i].onclick = sidebarSelectControl
}



let sidebar = document.querySelector("#sidebar")
function iframeSizeControl() {
    iframe.height = document.documentElement.clientHeight
    iframe.width = document.documentElement.clientWidth - sidebar.clientWidth
}
window.onload = iframeSizeControl
window.onresize = iframeSizeControl



let oldPhone
let oldAddress

let phoneInput
let addressInput

document.querySelector("#changInfo").onclick = function() {
    let xhr = new XMLHttpRequest()
    xhr.open("POST", `../info/get/${username}`)
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
            let info = JSON.parse(this.responseText)
            oldPhone = info.phone
            oldAddress = info.address

            phoneInput = document.querySelector("#phone")
            phoneInput.setAttribute("value", oldPhone)
            addressInput = document.querySelector("#address")
            addressInput.setAttribute("value", oldAddress)
        }
    }
    xhr.send()
}

document.querySelector("#postChangeBtn").onclick = function() {
    let newPhone = phoneInput.value
    if (newPhone == "") {
        alert("手机号不能为空！")
        return
    }
    if (newPhone == oldPhone) {
        newPhone = ""
    }

    let newAddress = addressInput.value
    if (newAddress == "") {
        alert("地址不能为空！")
        return
    }
    if (newAddress == oldAddress) {
        newAddress = ""
    }
    if (newPhone == "" && newAddress == "") {
        return
    }

    let xhr = new XMLHttpRequest()
    xhr.open("POST", `../info/change/${oldPhone}`)
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
            alert(this.responseText)
        }
    }
    xhr.send(JSON.stringify({
        "phone": newPhone,
        "address": newAddress
    }))
}

document.querySelector("#changePassword").onclick = function() {
    let role = document.querySelector("#role").innerText
    let password
    for( ; ; ) {
        password = prompt("正在执行敏感操作，需要先确认您的密码：")
        if (password == null) {
            return
        } else if (password == "") {
            alert("请输入密码！")
        } else {
            break
        }
    }
    
    let xhr = new XMLHttpRequest()
    xhr.open("POST", `../password/check/${username}`)
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
            let result = this.responseText
            if (result == "密码错误！") {
                alert(result)
            } else {
                let newPassword
                for ( ; ; ) {
                    newPassword = prompt("请输入新密码：")
                    if (newPassword == null) {
                        return
                    } else if (newPassword == "") {
                        alert("请输入新密码！")
                    } else {
                        break
                    }
                }


                let xhr = new XMLHttpRequest()
                xhr.open("POST", `../password/change/${username}`)
                xhr.onreadystatechange = function() {
                    if (xhr.readyState == 4 && xhr.status == 200) {
                        alert(this.responseText)
                    }
                }
                xhr.send(JSON.stringify({
                    "password": newPassword,
                    "role": role
                }))
            }
        }
    }
    xhr.send(JSON.stringify({
        "password": password,
        "role": role
    }))
}


document.querySelector("#logoutLi").onclick = function() {
    let result = confirm("确认退出登录？")
    if (result == true) {
        location.href = "/"
    }
}