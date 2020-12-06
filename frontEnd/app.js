function dis(val)
{
    document.getElementById("result").value+=val
}
function solve()
{
    let x = document.getElementById("result").value
    let y = eval(x)
    document.getElementById("result").value = y
}
function getExpression()
{
    let x = document.getElementById("result").value
    let y = eval(x)
    return x + "=" + y;
}

function clr()
{
    document.getElementById("result").value = ""
}

window.onload = function () {
    var conn;
    var msg = document.getElementById("result");
    var log = document.getElementById("log");

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    document.getElementById("Send").onclick = function () {
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        let expression= getExpression()
        solve()
        conn.send(expression);
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/update");
        conn.onclose = function (evt) {
            let item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            let messages = evt.data.split('\n');
            for (let i = 0; i < messages.length; i++) {
                const item = document.createElement("div");
                item.innerText = messages[i];
                appendLog(item);
            }
        };
    } else {
        let item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
};