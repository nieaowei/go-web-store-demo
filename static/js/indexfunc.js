function showItem() {
    $.post(
        "showItem",
        {
          page:1,
          rows:10,
        },
        function (data) {
            var a = data.data;
            $(".containerRight").text(a)
            for(var i=0;i<10;i++){
                $(".containerRight").text("标题"+a[i].title+"描述"+ a[i].sell_point + "价格"+a[i].price +"库存"+ a[i].num +"更新时间"+ a[i].updated+"\r\n")
            }
        },
        "json",
    )
}