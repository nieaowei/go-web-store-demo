function loginClick() {
    if (checkLoginInfo()) {
        var name = $("#entry_name").val();
        var password = $("#entry_password").val();
        $.post("/login",
            {
                username: name,
                password: password
            },
            function (data, textStatus, jqXHR) {
                if (data.status == 200) {
                    alert("登陆成功")
                } else {
                    alert("登陆失败")
                }
            },
            "json"
        );
    }
}

function regesterClick() {
    alert("提交注册信息")
}

function ToLoginClick() {
    var div = $(".containerT");
    div.hide(500, setLogin);
    div.show(1000, AtferTo);
}

function ToRegisterClick() {
    var div = $(".containerT");
    div.hide(500, setRegster);
    div.show(1000, AtferTo);
}

function AtferTo() {
    $("#entry_name").focus()
}

//设置注册所需组件
function setRegster() {
    $(document).unbind('keydown');
    $(document).keydown(function (event) {
        if (event.keyCode == 13) {
            $("#register_btn").click();
        }
    });
    var after = $("#entry_password");
    $("#header1").text("用户注册");
    $("<input type=\"text\" placeholder=\"手机号码\" id=\"entry_pthone\"> ").insertAfter(after);
    $("<input type=\"text\" placeholder=\"邮箱\" id=\"entry_email\"> ").insertAfter(after);
    $("#login_btn").attr('onclick', '').unbind('click');
    $("#login_btn").attr('onclick', 'ToLoginClick()').bind('click');
    $("#register_btn").attr('onclick', '').unbind('click');
    $("#register_btn").attr('onclick', 'regesterClick()').bind('click');
}

//设置登录所需组件
function setLogin() {
    $(document).unbind('keydown');
    $(document).keydown(function (event) {
        if (event.keyCode == 13) {
            $("#login_btn").click();
        }
    });
    $("#header1").text("用户登录");
    $("#entry_pthone").remove();
    $("#entry_email").remove();
    $("#login_btn").attr('onclick', '').unbind('click');
    $("#login_btn").attr('onclick', 'loginClick()').bind('click');
    $("#register_btn").attr('onclick', '').unbind('click');
    $("#register_btn").attr('onclick', 'ToRegisterClick()').bind('click')
}

//检查登录信息
function checkLoginInfo() {
    var name = $("#entry_name").val();
    var password =$("#entry_password").val();
    if(name == ""){
        alert("用户名不能为空");
        return false
    }else if (password==""){
        alert("密码不能为空");
        return false
    }else {
        return true
    }
}

//检查注册信息
function checkRegisterInfo() {

}