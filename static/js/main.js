var Main=new Vue({
    el:"#maincontainer",
    data:{
        openeds: ['1'],
        uniqueOpened: false,
        tableData:[],
        total:15,
        pagesize:10,
        currentPage:1,
    },
    methods: {
        pageChanged(page){
            Main.currentPage=page
            Main.handleClick()
        },
        handleOpen(key, keyPath) {
            console.log(key, keyPath);
        },
        handleClose(key, keyPath) {
            console.log(key, keyPath);
        },
        handleClick() {
            var that = this;
            $.post("/getTotal",
                {},
                function (data) {
                    if(data.status==200){
                        Main.total=data.data
                        console.log(data.data)
                    }
                },
                "json"
            )
            $.post("/showItem",
                {
                    rows: 10,
                    page: Main.currentPage,
                },
                function (data, textStatus, jqXHR) {
                    if (data.status == 200) {
                        Main.tableData=data.data
                        console.log(Main.tableData)
                    }
                },
                "json"
            );
        },
    }
});

