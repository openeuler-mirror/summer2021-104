var initDate = function () {
    $.ajax({
        url: "/api/user/token2user",
        type: "POST",
        data: {},
        headers: {
            'Content-Type': 'application/json',
            'token': $.cookie('token'),
        },
        async: true,
        success: function (data) {
            console.log(data.user_realname);
            $("#real_username").text(data.user_realname);
            if (data.user_realname == undefined) {
                console.log("未登陆")
                alert("请登陆")
                window.location.href = "/"
            }
        },
        fail: function (status) {
            console.log(status)
        }
    });
};

$(initDate);

var logout = function () {
    $.removeCookie('token', {
        path: '/'
    });
    window.location.href = "/"
}
