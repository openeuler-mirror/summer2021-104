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
        $("#RoleListTable").empty();
        if (RoleList != "" && RoleList != null) {
            for (var i = 0; i < RoleList.length; i++) {
                var TableTd = '<td style="vertical-align: middle;">' + RoleList[i]["role_id"] + "</td>";
                TableTd += '<td style="vertical-align: middle;">' + RoleList[i]["role_name"] + "</td>";
                TableTd += '<td style="vertical-align: middle;">' + RoleList[i]["user_num"] + "</td>";
                TableTd += '<td style="vertical-align: middle;">' + RoleList[i]["role_remark"] + "</td>";
                TableTd += '<td style="vertical-align: middle;"><button class="btn btn-danger btn-sm" type="button" onclick=DeleteRoleReconfirm(' + RoleList[i]["role_id"] + ",'" + RoleList[i]["role_name"] + "'," + RoleList[i]["user_num"] + ")>删除</button></td>";
                $("#RoleListTable").append("<tr>" + TableTd + "</tr>")
            }
        }
    })
}
$(InitData);

function DeleteRoleReconfirm(role_id, role_name, user_num) {
    $("#DeleteRoleName").text(role_name);
    $("#DeleteRoleModalButton").attr("onclick", "DeleteRole(" + role_id + "," + user_num + ")");
    $("#DeleteRoleModal").modal("show")
}

function DeleteRole(role_id, user_num) {
    if (user_num != 0) {
        alert("用户数量不为0，禁止删除。");
        $("#DeleteRoleModal").modal("hide");
        return
    }
    var settings = {
        "url": "/api/role/delete",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "token": $.cookie("token"),
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
            "role_id": role_id
        }),
    };
    $.ajax(settings).done(function (response) {
        console.log(response);
        $("#DeleteRoleModal").modal("hide");
        $(InitData)
    })
}

function CreateRoleFunc() {
    var Data = jqu.formData("CreateRoleForm");
    console.log(JSON.stringify(Data));
    var settings = {
        "url": "/api/role/create",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "token": $.cookie("token"),
            "Content-Type": "application/json"
        },
        "data": JSON.stringify(Data),
    };
    $.ajax(settings).done(function (response) {
        console.log(response)
    });
    $("#addRole").modal("hide");
    $(InitData)
};