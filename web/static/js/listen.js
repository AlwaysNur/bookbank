const audio = document.querySelector('audio'); // <audio>
const descElement = document.querySelector(".description")
const savedPosition = localStorage.getItem(`bookbankPlaybackLocation-${id}`);
const toolbarOptions = document.getElementById("toolbar-options")
const toolBarDiv = document.getElementById("toolbar-options-div")
descElement.innerHTML = descElement.innerHTML.replaceAll(".", ".<br>").replaceAll("!", "!<br>")

function handleMoreClick() {
    if (toolBarDiv.style.display == "none") {
        toolBarDiv.style.display = "block";
        toolBarDiv.style.opacity = "1";
    } else {
        toolBarDiv.style.display = "none";
        toolBarDiv.style.opacity = "0";
    }
}

if (savedPosition) {
    audio.currentTime = parseFloat(savedPosition);
}

window.addEventListener('beforeunload', () => {
    localStorage.setItem(`bookbankPlaybackLocation-${id}`, audio.currentTime);
});

audio.addEventListener('pause', () => {
    localStorage.setItem(`bookbankPlaybackLocation-${id}`, audio.currentTime);
});


