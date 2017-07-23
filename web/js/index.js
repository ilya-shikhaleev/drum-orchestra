const socketUrl = "ws://192.168.21.59:8083/ws";
const audio = new Audio("audio/hihat.wav");

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
        audio.play();
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