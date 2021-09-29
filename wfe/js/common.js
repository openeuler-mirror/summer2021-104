var initData = function () {
    $("#real_username").text($.cookie("realname"));
    $("#real_username").attr("href", "#")
};
$(initData);
var logout = function () {
    $.removeCookie("token", {
        path: "/"
    });
    window.location.href = "/"
};