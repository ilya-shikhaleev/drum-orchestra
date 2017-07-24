const socketUrl = "ws://" + window.location.hostname + ":" + window.location.port + "/ws";

const socket = new WebSocket(socketUrl);
socket.onopen = function () {
    console.log("Соединение установлено.");
};

socket.onclose = function (event) {
    if (event.wasClean) {
        console.log('Соединение закрыто чисто');
    } else {
        console.log('Обрыв соединения'); // например, "убит" процесс сервера
    }
    console.log('Код: ' + event.code + ' причина: ' + event.reason);
};

socket.onmessage = function (e) {
    console.log("Получены данные " + e.data);
    const data = e.data;
    console.log(data);
    const event = JSON.parse(data);
    console.log(event);
    if (event.event === "play") {
        if (isMobile()) {
            window.speechSynthesis.speak(new SpeechSynthesisUtterance('meow'));
        } else {
            const audio = new Audio("audio/hihat.wav");
            audio.play();
        }
    }
};

socket.onerror = function (error) {
    console.log("Ошибка " + error.message);
};

const button = document.getElementById("playButton");
button.onclick = () => {
    socket.send("play");
    return false;
};


function isMobile() {
    let useragent = navigator.userAgent;

    if (useragent.match(/Android/i)) {
        return true;
    } else if (useragent.match(/webOS/i)) {
        return true;
    } else if (useragent.match(/iPhone/i)) {
        return true;
    } else if (useragent.match(/iPod/i)) {
        return true;
    } else if (useragent.match(/iPad/i)) {
        return true;
    } else if (useragent.match(/Windows Phone/i)) {
        return true;
    } else if (useragent.match(/SymbianOS/i)) {
        return true;
    } else if (useragent.match(/RIM/i) || useragent.match(/BB/i)) {
        return true;
    } else {
        return false;
    }
}