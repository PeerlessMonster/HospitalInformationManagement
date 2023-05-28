let passwordConfirmInput = document.querySelector("[placeholder='请再次输入密码确认']")
passwordConfirmInput.onchange = function() {
    let registerBtn = document.querySelector("#registerBtn")
    let first = passwordConfirmInput.previousElementSibling.value
    let second = passwordConfirmInput.value

    let replace
    if (first != second) {
        replace = "block"
        registerBtn.setAttribute("disabled", "disabled")
    } else {
        replace = "none"
        registerBtn.removeAttribute("disabled")
    }
    document.querySelector("#passwordNEqual").style.display = replace
}

document.querySelector("#linkLoginBtn").addEventListener("click", function() {
    location.href = "/"
})