let linkRegisterBtn = document.querySelector("#linkRegisterBtn")
linkRegisterBtn.addEventListener("click", function() {
    location.href = "register"
})

document.querySelector("select").onchange = function() {
    let usernameInput = document.querySelector("[name='username']")
    let usernameLabel = usernameInput.nextElementSibling
    
    let value = this.options.selectedIndex
    let replace
    if (value == 0) {
        replace = "手机号"
        
        linkRegisterBtn.style.display = "inline-block";
    } else {
        replace = "工号"

        linkRegisterBtn.style.display = "none";
    }
    usernameInput.setAttribute("placeholder", replace)
    usernameLabel.innerText = replace
}






