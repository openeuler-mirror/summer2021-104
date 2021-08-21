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