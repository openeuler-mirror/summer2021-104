function InitData() {
    var settings = {
        url: "/api/audit/log",
        type: "POST",
        headers: {
            "token": $.cookie("token"),
            "Content-Type": "application/json",
        },
        async: true
    };
    $.ajax(settings).success(function (response) {
        console.log(response);
        $("#event_list_table").empty();
        var events = response["event_list"];
        if (events != "" && events != null) {
            for (var i = 0; i < events.length; i++) {
                $("#event_list_table").append("<tr><td>" + events[i]["username"] + "</td><td>" + events[i]["event"] + "</td><td>" + events[i]["event_time"] + "</td></tr>")
            }
        }
    })
}
$(InitData);