function InitData() {
    var settings = {
        "url": "/api/host/ReadHostList",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "token": $.cookie("token")
        },
        "data": JSON.stringify({
            "KeyWords": ""
        }),
        "async": false
    };
    $.ajax(settings).done(function (response) {
        var HostList = response["host_management_list"]["host_management"];
        if (HostList != null) {
            $("#MonitorList").empty();
            for (var i = 0; i < HostList.length; i++) {
                console.log(HostList[i]);
                var DIV = '<div class="monitor" role="group">';
                DIV += '<div class="monitor-view" role="group">';
                DIV += '<button type="button" class="btn btn-default dropdown-toggle">';
                DIV += HostList[i]["host_management_name"];
                DIV += '<span class="glyphicon glyphicon-chevron-down"></span>';
                DIV += "</button>";
                DIV += '<div class="allview" style="display: block;">';
                DIV += '<div class="Line-chart" id="CPU' + HostList[i]["host_management_id"] + '"></div>';
                DIV += '<div class="Histogram" id="Memory' + HostList[i]["host_management_id"] + '" style="margin-left: 50px;"></div>';
                DIV += "</div></div></div>";
                $("#MonitorList").append(DIV);
                var CPU = echarts.init(document.getElementById("CPU" + HostList[i]["host_management_id"]));
                var Memory = echarts.init(document.getElementById("Memory" + HostList[i]["host_management_id"]));
                var CPUsettings = {
                    "url": "/api/monitor/getcpu",
                    "method": "POST",
                    "timeout": 0,
                    "headers": {
                        "token": $.cookie("token"),
                        "Content-Type": "application/json"
                    },
                    "data": JSON.stringify({
                        "host_id": HostList[i]["host_management_id"]
                    }),
                    "async": false
                };
                var Memorysettings = {
                    "url": "/api/monitor/getmemory",
                    "method": "POST",
                    "timeout": 0,
                    "headers": {
                        "token": $.cookie("token"),
                        "Content-Type": "application/json"
                    },
                    "data": JSON.stringify({
                        "host_id": HostList[i]["host_management_id"]
                    }),
                    "async": false
                };
                $.ajax(CPUsettings).done(function (response) {
                    var CPUOption = {
                        title: {
                            text: "CPU 使用率"
                        },
                        xAxis: {
                            type: "category",
                            data: response["Time"]
                        },
                        yAxis: {
                            type: "value",
                            axisLabel: {
                                formatter: "{value} %"
                            }
                        },
                        series: [{
                            data: response["Value"],
                            type: "line",
                            smooth: true
                        }]
                    };
                    CPU.setOption(CPUOption)
                });
                $.ajax(Memorysettings).done(function (response) {
                    var MemoryOption = {
                        title: {
                            text: "Memory 使用率"
                        },
                        xAxis: {
                            type: "category",
                            data: response["Time"]
                        },
                        yAxis: {
                            type: "value",
                            axisLabel: {
                                formatter: "{value} %"
                            }
                        },
                        series: [{
                            data: response["Value"],
                            type: "line",
                            smooth: true
                        }]
                    };
                    Memory.setOption(MemoryOption)
                })
            }
        }
    })
}
$(InitData);