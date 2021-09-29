function InitData() {
    var settings = {
        "url": "/api/setting/get",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "token": $.cookie("token")
        },
    };
    $.ajax(settings).done(function (response) {
        $("#host_url").attr("value", response["Setting"]["host_url"]);
        $("#ssh_public_key").val(response["Setting"]["ssh_public_key"]);
        $("#ssh_private_key").val(response["Setting"]["ssh_private_key"])
    })
}
$(InitData);

function UpdateSetting() {
    var Data = jqu.formData("UpdateSettingForm");
    console.log();
    var settings = {
        "url": "/api/setting/set",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "token": $.cookie("token"),
            "Content-Type": "application/json"
        },
        "data": JSON.stringify(Data),
    };
    $.ajax(settings).done(function (response) {
        console.log(response);
        alert("Save Success.")
    })
};