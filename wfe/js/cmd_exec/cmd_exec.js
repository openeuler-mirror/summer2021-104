var cmd_exec = function () {
    var data = jqu.formData('cmd_exec_form');
    console.log(JSON.stringify(data))
    if (data.host == undefined) {
        alert('请选择主机');
        return;
    }
    if (data.cmd == '') {
        alert('请输入命令');
        return;
    }

    jqu.loadJson('/api/cmd/ExecCmd', data, {
        'token': $.cookie('token')
    }, function (result) {
        if (!result.succ) {
            alert(result.stmt);
            return;
        }
        alert('欢迎 ' + result.user_realname + ' 登录 DevOps 系统！');
        $.cookie('token', result.token);
        window.location.href = "/dashboard.html";
    });
    $.ajax({
        url: "/api/cmd/ExecCmd",
        type: "POST",
        data: data,
        headers: {
            'token': $.cookie('token'),
        },
        async: true,
        success: function (result) {
            console.log(result.stdout);
            $("#cmd_exec_stdout").empty();
            $("#cmd_exec_stdout").text(result.stdout);
        },
        fail: function (status) {
            console.log(status);
            $("#cmd_exec_stdout").empty();
        }
    });
}