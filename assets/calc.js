// This function clear all the values
function clearDisplay() {
    document.getElementById("result").value = "";
}

// This function display values
function display(value) {
    document.getElementById("result").value += value;
}
// This function evaluates the expression and return result
function calculate() {
    var p = document.getElementById("result").value;
    var q = eval(p);
    document.getElementById("result").value = q;
}


function saveInfo() {
    var x = document.getElementById('textarea').value
    sessionStorage.setItem("textarea", x)
}

function pasteInfo(){
    var x = sessionStorage.getItem("textarea");
    document.GetElementById("textarea").innerHTML = x;
}