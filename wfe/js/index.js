var devops_login = function () {
    var data = jqu.formData('devops_login_form');
    if (data.username == '') {
        alert('请输入用户名');
        return;
    }
    if (data.password == '') {
        alert('请输入密码');
        return;
    }
    jqu.loadJson('/api/user/login', data, function (result) {
        if (!result.succ) {
            alert(result.stmt);
            return;
        }
        alert('欢迎 ' + result.user_realname + ' 登录 DevOps 系统！');
        $.cookie('token', result.token);
        $.cookie('realname', result.user_realname);
        $.cookie('user_id', result.user_id);
        $.cookie('username', data.username);
        window.location.href = "/dashboard.html";
    });
};
