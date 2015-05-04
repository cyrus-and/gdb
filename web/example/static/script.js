function Client(host, onNotification, onOutput) {
    var notificationWs = new WebSocket('ws://' + host + '/notifications');
    var terminalWs = new WebSocket('ws://' + host + '/terminal');

    notificationWs.onmessage = function (event) {
        var object = JSON.parse(event.data);
        onNotification(object);
    };

    terminalWs.onmessage = function (event) {
        onOutput(event.data);
    };

    this.input = function (data) {
        terminalWs.send(data);
    };

    this.command = function (onReply) {
        var command = Array.prototype.slice.call(arguments).slice(1);
        var ajax = new XMLHttpRequest();
        ajax.open('POST', 'http://' + host + '/send');
        ajax.onreadystatechange = function () {
            if (ajax.readyState === 4 && ajax.status === 200){
                var object = JSON.parse(ajax.responseText);
                onReply(object);
            }
        };
        ajax.send(JSON.stringify(command));
    };

    this.interrupt = function (onReply) {
        var ajax = new XMLHttpRequest();
        ajax.open('POST', 'http://' + host + '/interrupt');
        ajax.onreadystatechange = function () {
            if (ajax.readyState === 4 && ajax.status === 200){
                onReply && onReply();
            }
        };
        ajax.send(null);
    };
}

///

var host = 'localhost:8080';

function bind(id, action) {
    var elem = document.getElementById(id);
    elem.addEventListener('click', action);
}

function onNotification(object) {
    console.log(JSON.stringify(object, null, 4));
    if (object.type === 'exec') {
        var status = object.class;
        if (status === 'stopped') {
            status += ', ' + object.payload.reason;
            if (typeof object.payload.frame !== 'undefined') {
                if (typeof object.payload.frame.func !== 'undefined') {
                    status += ', ' + object.payload.frame.func;
                }
                if (typeof object.payload.frame.file !== 'undefined') {
                    status += ' ' + object.payload.frame.file;
                }
                if (typeof object.payload.frame.line !== 'undefined') {
                    status += ':' + object.payload.frame.line;
                }
                if (typeof object.payload.frame.fullname !== 'undefined') {
                    var filePath = object.payload.frame.fullname;
                    var ajax = new XMLHttpRequest();
                    ajax.open('POST', 'http://' + host + '/source' + filePath);
                    ajax.onreadystatechange = function () {
                        if (ajax.readyState === 4 && ajax.status == 200){
                            var container = document.getElementById('source');
                            container.innerHTML = '';
                            var source = ajax.responseText;
                            var lines = source.split('\n');
                            for (var i = 0; i < lines.length; i++) {
                                var line = document.createElement('span');
                                line.id = i + 1;
                                line.appendChild(document.createTextNode(lines[i] || ' '));
                                container.appendChild(line);
                            }
                            var lineNumber = object.payload.frame.line;
                            document.getElementById(lineNumber).className = 'selected';
                        }
                    };
                    ajax.send(null);
                }
            }
        }
        document.getElementById('status').innerHTML = status;
    }
}

function onOutput(data) {
    document.getElementById('output').innerHTML += data;
}

var client = new Client(host, onNotification, onOutput);

function dumpResult(object) {
    console.log(JSON.stringify(object, null, 4));
    var result = document.getElementById('result');
    if (object.class === 'error') {
        result.innerHTML = object.payload.msg;
    } else {
        result.innerHTML = 'Done';
    }
}

bind('load', function () {
    var fileName = document.getElementById('file').value;
    client.command(dumpResult, 'file-exec-and-symbols', fileName);
});

bind('run', function () {
    client.interrupt();
    document.getElementById('output').innerHTML = '';
    var argText = document.getElementById('arguments').value.trim();
    if (argText !== '') {
        var args = argText.split(/ +/);
        client.command.apply(null, [dumpResult, 'exec-arguments'].concat(args));
    }
    client.command(dumpResult, 'exec-run', '--start');
});

bind('continue', function () {
    client.command(dumpResult, 'exec-continue');
});

bind('step', function () {
    client.command(dumpResult, 'exec-next');
});

bind('interrupt', function () {
    client.interrupt();
});

bind('send', function () {
    var line = document.getElementById('line').value;
    client.input(line + '\n');
});

bind('eot', function () {
    client.input('\x04');
});
