function GetHostList(KeyWords) {
    $.ajax({
        url: "/api/host/ReadHostList",
        type: "POST",
        data: JSON.stringify({
            "KeyWords": KeyWords
        }),
        headers: {
            "Content-Type": "application/json",
            "token": $.cookie("token"),
        },
        async: true,
        success: function (result) {
            $("#host_management_list_table").empty();
            var hosts = result.host_management_list.host_management;
            if (hosts != "" && hosts != null) {
                for (var i = 0; i < hosts.length; i++) {
                    var host = hosts[i];
                    console.log(host);
                    $("#host_management_list_table").append("<tr>" + "<td>" + host.host_management_id + "</td>" + "<td>" + host.host_management_ip + "</td>" + "<td>" + host.host_management_name + "</td>" + "<td>" + host.host_management_user + "</td>" + "<td>" + host.host_management_create_time + "</td>" + "<td>" + '<button class="btn btn-danger" onclick="DeleteHost(\'' + host.host_management_id + "')\">删除</button>" + "</td>" + "</tr>")
                }
            }
        },
        fail: function (status) {
            console.log(status)
        }
    })
}

function initHostData() {
    GetHostList("")
}
$(initHostData);
var CreateHost = function () {
    var data = jqu.formData("create_host_from");
    if (data.HostIP == "") {
        alert("请输入IP");
        return
    }
    if (data.HostPort == "") {
        alert("请输入端口");
        return
    }
    if (data.HostName == "") {
        alert("请输入主机名");
        return
    }
    if (data.HostUser == "") {
        alert("请输入用户名");
        return
    }
    if (data.HostPassword == "") {
        alert("请输入密码");
        return
    }
    var settings = {
        url: "/api/host/CreateHost",
        type: "POST",
        data: JSON.stringify(data),
        headers: {
            "token": $.cookie("token"),
            "Content-Type": "application/json",
        },
        async: true
    };
    $.ajax(settings).success(function (response) {
        if (response["message"] == "OK") {
            $("#CreateHost").modal("hide");
            $("#HostIP").val("");
            $("#HostName").val("");
            $("#HostPassword").val("");
            $(initHostData)
        } else {
            alert(response["message"])
        }
    })
};

function DeleteHost(HostID) {
    var message = "确定要删除主机编号为 " + HostID + " 的主机吗？";
    if (confirm(message)) {
        var data = {
            "HostID": HostID
        };
        var settings = {
            "url": "/api/host/DeleteHost",
            "method": "POST",
            "timeout": "0",
            "data": data,
            "headers": {
                "token": $.cookie("token"),
            },
        };
        $.ajax(settings).done(function (response) {
            $(initHostData)
        })
    } else {
        return
    }
}

function SearchHost() {
    var KeyWords = $("#SearchHostKeyWords").val();
    GetHostList(KeyWords)
};