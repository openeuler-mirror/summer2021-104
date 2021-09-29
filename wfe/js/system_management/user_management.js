function GetUserList(KeyWords) {
    var settings = {
        "url": "/api/user/getlist",
        "method": "POST",
        "timeout": 0,
        "data": JSON.stringify({
            "KeyWords": KeyWords
        }),
        "headers": {
            "token": $.cookie("token"),
            "Content-Type": "application/json",
        },
    };
    $.ajax(settings).done(function (response) {
        var UserList = response.UserList;
        $("#user_list_table").empty();
        if (UserList != "" && UserList != null) {
            for (var i = 0; i < UserList.length; i++) {
                $("#user_list_table").append('<tr><td style="vertical-align: middle;">' + UserList[i]["user_id"] + '</td><td style="vertical-align: middle;">' + UserList[i]["user_username"] + '</td><td style="vertical-align: middle;">' + UserList[i]["user_realname"] + "</td><td>" + UserList[i]["role_name"] + "</td><td>" + UserList[i]["user_status"] + "</td><td>" + UserList[i]["user_login_time"] + '</td><td><button class="btn btn-default btn-sm" type="button" onclick=UpdateUserPasswordModal(' + UserList[i]["user_id"] + ",'" + UserList[i]["user_username"] + '\')>重置密码</button> <button class="btn btn-danger btn-sm" type="button" onclick=DeleteUserReconfirm(' + UserList[i]["user_id"] + ",'" + UserList[i]["user_username"] + "')>删除</button></td></tr>")
            }
        }
    })
}

function SearchUser() {
    var KeyWords = $("#SearchKeyWords").val();
    GetUserList(KeyWords)
}

function InitData() {
    var settings = {
        "url": "/api/role/getlist",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "token": $.cookie("token")
        },
    };
    $.ajax(settings).done(function (response) {
        var RoleList = response.RoleList;
        $("#RoleList").empty();
        if (RoleList != "" && RoleList != null) {
            $("#RoleList").append("<option disabled selected hidden>请选择</option>");
            for (var i = 0; i < RoleList.length; i++) {
                $("#RoleList").append('<option value="' + RoleList[i]["role_id"] + '">' + RoleList[i]["role_name"] + "</option>")
            }
        }
    });
    GetUserList("")
}
$(InitData);
var eye = document.getElementById("eye");
var password = document.getElementById("password");

function hideShowPsw() {
    if (password.type == "password") {
        password.type = "text";
        eye.className = "glyphicon glyphicon-eye-open"
    } else {
        password.type = "password";
        eye.className = "glyphicon glyphicon-eye-close"
    }
}

function CreateUser() {
    var Data = jqu.formData("CreateUserForm");
    var settings = {
        "url": "/api/user/create",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "token": $.cookie("token"),
            "Content-Type": "application/json"
        },
        "data": JSON.stringify(Data),
    };
    $.ajax(settings).done(function (response) {
        if (response["message"] == "OK") {
            window.location.reload()
        } else {
            alert("Error.")
        }
    })
}

function DeleteUserReconfirm(UserID, UserName) {
    $("#DeleteModalUserName").text(UserName);
    $("#DeleteModalButton").attr("onclick", "DeleteUserByUserID(" + UserID + ")");
    $("#deleteModal").modal("show")
}

function UpdateUserPasswordModal(UserID, UserName) {
    $("#UpdateModalUserID").text(UserID);
    $("#UpdateModalUserName").attr("value", UserName);
    $("#UpdateModalButton").attr("onclick", "UpdateUserPasswordByUserID(" + UserID + ")");
    $("#passwordModal").modal("show")
}

function UpdateUserPasswordByUserID(UserID) {
    var PasswordText = $("#UpdateModalUserPassword").val();
    var settings = {
        "url": "/api/user/ResetPasswordByUserID",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "token": $.cookie("token"),
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
            "user_id": UserID,
            "user_password": PasswordText
        }),
    };
    $.ajax(settings).done(function (response) {
        console.log(response);
        alert("修改成功！");
        $("#passwordModal").modal("hide")
    })
}

function DeleteUserByUserID(UserID) {
    var settings = {
        "url": "/api/user/delete",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "token": $.cookie("token"),
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
            "user_id": UserID
        }),
    };
    $.ajax(settings).done(function (response) {
        console.log(response);
        $("#deleteModal").modal("hide");
        $(InitData)
    })
};