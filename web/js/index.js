const socketUrl = "ws://" + window.location.hostname + ":" + window.location.port + "/ws";
const audioPath = "audio/hihat.wav";
const audio = new Audio(audioPath); //Load into cache

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
        const audio = new Audio(audioPath);
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