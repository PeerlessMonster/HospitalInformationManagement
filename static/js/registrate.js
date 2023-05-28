let registarteBtns = document.querySelectorAll(".btn-primary")
for (let i = 0; i < registarteBtns.length; i++) {
    registarteBtns[i].onclick = function() {
        let registratebtnDiv = this.parentNode
        let doctorNameH5 = registratebtnDiv.previousElementSibling.previousElementSibling
        let doctorName = doctorNameH5.innerText
        
        let result = confirm(`确定挂号【${doctorName}】医生？`)
        if (result == true) {
            let patientName = document.querySelector("#patientName").innerText

            let xhr = new XMLHttpRequest()
            xhr.open("POST", `${patientName}/${doctorName}`)
            xhr.onreadystatechange = function() {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    let order = this.responseText
                    alert(`挂号成功！你是第${order}位`)
                }
            }
            xhr.send()
        }
    }
}