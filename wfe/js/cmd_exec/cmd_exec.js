function InitData() {
    var settings = {
        "url": "/api/host/ReadHostList",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "token": $.cookie("token")
        },
        "data": JSON.stringify({
            "KeyWords": ""
        })
    };
    $.ajax(settings).done(function (response) {
        var HostList = response["host_management_list"]["host_management"];
        if (HostList != null) {
            $("#host_list").empty();
            $("#host_list").append("<option disabled selected hidden>请选择</option>");
            for (var i = 0; i < HostList.length; i++) {
                $("#host_list").append('<option value="' + HostList[i]["host_management_id"] + '">' + HostList[i]["host_management_user"] + "@" + HostList[i]["host_management_ip"] + "</option>")
            }
        }
    })
}
$(InitData);
var cmd_exec = function () {
    var data = jqu.formData("cmd_exec_form");
    if (data.host == undefined) {
        alert("请选择主机");
        return
    }
    if (data.cmd == "") {
        alert("请输入命令");
        return
    }
    var settings = {
        "url": "/api/cmd/ExecCmd",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "token": $.cookie("token"),
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
            "host_id": data.host,
            "cmd": data.cmd
        }),
    };
    $.ajax(settings).done(function (response) {
        $("#cmd_exec_stdout").empty();
        $("#cmd_exec_stdout").text(response["stdout"])
    })
};