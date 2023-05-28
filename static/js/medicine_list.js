document.querySelector("#showList").onclick = function() {
    let showListBtn = this
    let listNo = showListBtn.innerText

    let xhr = new XMLHttpRequest()
    xhr.open("POST", `post/${listNo}`)
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
            let cardBodyParent = showListBtn.parentNode.nextElementSibling
            let cardBodyDiv = cardBodyParent.children[0]
            cardBodyDiv.innerHTML = this.responseText
        }
    }
    xhr.send()
}