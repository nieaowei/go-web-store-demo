function test111() {
    $(".containerT").css("animation", "myfirst 2s forwards")
}

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
                    $(".containerT").hide(500, function () {
                        loadIndexHtml();
                    });
                } else {
                    $("#status").text("账号或密码错误，请重新尝试")
                }
            },
            "json"
        );
    }
}

function regesterClick() {
    if (checkRegisterInfo()) {
        $.post("/register",
            {
                username: $("#entry_name").val(),
                password: $("#entry_password").val(),
                email: $("#entry_email").val(),
                phone: $("#entry_phone").val(),
            },
            function (data, textStatus, jqXHR) {
                if (data.status == 200) {
                    $("#status").text("注册成功")
                } else {
                    $("#status").text("注册失败")
                }
            },
            "json"
        );
    }
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
    $("#status").text("")
    $(".containerT").css("top", "40%");
    $(document).unbind('keydown');
    $(document).keydown(function (event) {
        if (event.keyCode == 13) {
            $("#register_btn").click();
        }
    });
    var after = $("#entry_password");
    $("#header1").text("用户注册");
    $("<input type=\"text\" placeholder=\"手机号码\" id=\"entry_phone\"> ").insertAfter(after);
    $("<input type=\"text\" placeholder=\"邮箱\" id=\"entry_email\"> ").insertAfter(after);
    $("#login_btn").attr('onclick', '').unbind('click');
    $("#login_btn").attr('onclick', 'ToLoginClick()').bind('click');
    $("#register_btn").attr('onclick', '').unbind('click');
    $("#register_btn").attr('onclick', 'regesterClick()').bind('click');
}

//设置登录所需组件
function setLogin() {
    $("#status").text("")
    $(".containerT").css("top", "50%");
    $(document).unbind('keydown');
    $(document).keydown(function (event) {
        if (event.keyCode == 13) {
            $("#login_btn").click();
        }
    });
    $("#header1").text("用户登录");
    $("#entry_phone").remove();
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
        $("#status").text("账号为空，请输入账号")
        return false
    }else if (password==""){
        $("#status").text("密码为空，请输入密码")
        return false
    }else {
        return true
    }
}

//检查注册信息
function checkRegisterInfo() {
    if ($("#entry_name").val() == "") {
        $("#status").text("账号为空，请输入账号");
        return false;
    } else if ($("#entry_password").val() == "") {
        $("#status").text("密码为空，请输入密码");
        return false;
    } else if ($("#entry_email").val() == "") {
        $("#status").text("邮箱为空，请输入邮箱");
        return false;

    } else if ($("#entry_phone").val() == "") {
        $("#status").text("手机号码为空，请输入手机号码");
        return false;
    }
    return true
}

function loadIndexHtml() {
    $(".maincontainer").load("/view/index.html", function (respone) {
        $(".maincontainer").html(respone, function () {
        });
    });
    $(".containerRight").css("animation", "myfirst 3s")
}