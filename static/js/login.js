function loginClick() {
    checkLoginInfo()
}

function checkLoginInfo() {
    var name = document.getElementById("entry_name").value
    var password = document.getElementById("entry_password").value
    if(name == ""){
        alert("用户名不能为空")
    }else if (password==""){
        alert("密码不能为空")
    }
    if(name =="nieaowei" && password=="nieaowei"){

    }else {
        alert("密码错误")
        document.getElementById("name")
    }
}