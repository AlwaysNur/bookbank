const audio = document.querySelector('audio'); // <audio>
const descElement = document.querySelector(".description")
const savedPosition = localStorage.getItem(`bookbankPlaybackLocation-${id}`);

descElement.innerHTML = descElement.innerHTML.replaceAll(".", ".<br>").replaceAll("!", "!<br>")

if (savedPosition) {
    audio.currentTime = parseFloat(savedPosition);
}

window.addEventListener('beforeunload', () => {
    localStorage.setItem(`bookbankPlaybackLocation-${id}`, audio.currentTime);
});

audio.addEventListener('pause', () => {
    localStorage.setItem(`bookbankPlaybackLocation-${id}`, audio.currentTime);
});

