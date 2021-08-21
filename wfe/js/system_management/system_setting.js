var eye = document.getElementById("eye");
var password = document.getElementById("password");

function hideShowPsw() {
    if (password.type == "password") {
        password.type = "text";
        eye.className = 'glyphicon glyphicon-eye-open'
    } else {
        password.type = "password";
        eye.className = 'glyphicon glyphicon-eye-close'
    }
}

var eyes = document.getElementById("eyes");
var passwords = document.getElementById("passwords");

function hideShowPsws() {
    if (passwords.type == "password") {
        passwords.type = "text";
        eyes.className = 'glyphicon glyphicon-eye-open'
    } else {
        passwords.type = "password";
        eyes.className = 'glyphicon glyphicon-eye-close'
    }
}

$("[name='test']").bootstrapSwitch();

function raClick() {
    var show = "";
    var apm = document.getElementsByName("identity");

    for (var i = 0; i < apm.length; i++) {
        if (apm[i].checked)
            show = apm[i].value;
    }
    switch (show) {
        case '内置':
            document.getElementById("selfbutton").style.display = "block";
            document.getElementById("selfset").style.display = "none";
            break;
        case '自定义':
            document.getElementById("selfbutton").style.display = "none";
            document.getElementById("selfset").style.display = "block";
            break;
        default:
            document.getElementById("selfbutton").style.display = "none";
            document.getElementById("selfset").style.display = "none";
            break;
    }
}
$(function () {
    $("#checkbox").bootstrapSwitch({
        onText: "ON",
        offText: "OFF",
        onColor: "primary",
        offColor: "info",
        size: "small",
        onSwitchChange: function (event, state) {
            if (state == true) {
                alert("ON");
            } else {
                alert("OFF");
            }
        }
    });
});

function getOne() {
    document.getElementById("one").style.display = "block";
    document.getElementById("two").style.display = "none";
    document.getElementById("four").style.display = "none";
    document.getElementById("five").style.display = "none";
    document.getElementById("seven").style.display = "none";
}

function getTwo() {
    document.getElementById("one").style.display = "none";
    document.getElementById("two").style.display = "block";
    document.getElementById("four").style.display = "none";
    document.getElementById("five").style.display = "none";
    document.getElementById("seven").style.display = "none";
}

function getFour() {
    document.getElementById("one").style.display = "none";
    document.getElementById("two").style.display = "none";
    document.getElementById("four").style.display = "block";
    document.getElementById("five").style.display = "none";
    document.getElementById("seven").style.display = "none";
}

function getFive() {
    document.getElementById("one").style.display = "none";
    document.getElementById("two").style.display = "none";
    document.getElementById("four").style.display = "none";
    document.getElementById("five").style.display = "block";
    document.getElementById("seven").style.display = "none";
}

function getSix() {
    document.getElementById("one").style.display = "none";
    document.getElementById("two").style.display = "none";
    document.getElementById("four").style.display = "none";
    document.getElementById("five").style.display = "none";
    document.getElementById("seven").style.display = "none";
}

function getSeven() {
    document.getElementById("one").style.display = "none";
    document.getElementById("two").style.display = "none";
    document.getElementById("four").style.display = "none";
    document.getElementById("five").style.display = "none";
    document.getElementById("seven").style.display = "block";
}
