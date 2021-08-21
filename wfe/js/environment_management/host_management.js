var initHostDate = function () {
    $.ajax({
        url: "/api/host/ReadHostList",
        type: "POST",
        data: {},
        headers: {
            'Content-Type': 'application/json',
            'token': $.cookie('token'),
        },
        async: true,
        success: function (result) {
            var hosts = result.host_management_list.host_management
            if (hosts != "" && hosts != null) {
                for (var i = 0; i < hosts.length; i++) {
                    var host = hosts[i];
                    console.log(host);
                    $("#host_management_list_table").append("<tr>" +
                        "<td>" + host.host_management_id + "</td>" +
                        "<td>" + host.host_management_ip + "</td>" +
                        "<td>" + host.host_management_name + "</td>" +
                        "<td>" + host.host_management_user + "</td>" +
                        "<td>" + host.host_management_connect_method + "</td>" +
                        "<td>" + host.host_management_create_time + "</td>" +
                        "<td>" +
                        "<button class=\"btn btn-default btn-xs\">\n<i class=\"fa fa-cog\"></i> 配置\n</button>\n" +
                        "<button class=\"btn btn-danger btn-xs disabled\">\n<i class=\"fa fa-trash-o\"></i> 删除\n</button>" +
                        "</td>" +
                        "</tr>"
                    );
                }
            }
        },
        fail: function (status) {
            console.log(status)
        }
    });
};

$(initHostDate);

var CreateHost = function () {
    var data = jqu.formData('create_host_from');
    if (data.HostKey == undefined) {
        data.HostKey = "NULL"
    }
    if (data.HostIP == '') {
        alert('请输入IP');
        return;
    }
    if (data.HostName == '') {
        alert('请输入主机名');
        return;
    }
    if (data.HostUser == '') {
        alert('请输入用户名');
        return;
    }
    if (data.HostPassword == '') {
        alert('请输入密码');
        return;
    }
    $.ajax({
        url: "/api/host/CreateHost",
        type: "POST",
        data: data,
        headers: {
            'token': $.cookie('token'),
        },
        async: true,
        success: function (result) {
            console.log(result);
        },
        fail: function (status) {
            console.log(status);
        }
    });
}