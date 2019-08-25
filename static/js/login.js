function loginClick() {
    checkLoginInfo()
}

function registerClick() {
    $(".containerT").hide(500);
    setTimeout(setRegster,500)
    $(".containerT").show(1000);
}

function setRegster() {
    $("#header1").text("用户注册");
    $("<input type=\"text\" placeholder=\"手机号码\" id=\"entry_pthone\"> ").insertAfter($("#entry_password"));
    $("<input type=\"text\" placeholder=\"邮箱\" id=\"entry_email\"> ").insertAfter($("#entry_password"));
}

function checkLoginInfo() {
    var name = $("#entry_name").val();
    var password =$("#entry_password").val();
    if(name == ""){
        alert("用户名不能为空")
    }else if (password==""){
        alert("密码不能为空")
    }else {
        $.post("/login",
            {
                username:name,
                password:password
            },
            function(data, textStatus, jqXHR){
                if(data.status==200){
                    alert("登陆成功")
                }else{
                    alert("登陆失败")
                }
            },
            "json"
        );
    }
}