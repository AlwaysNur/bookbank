const audio = document.querySelector('audio'); // Get your audio element

const savedPosition = localStorage.getItem(`bookbankPlaybackLocation-${id}`);

if (savedPosition) {
    audio.currentTime = parseFloat(savedPosition);
}

window.addEventListener('beforeunload', () => {
    localStorage.setItem(`bookbankPlaybackLocation-${id}`, audio.currentTime);
});

audio.addEventListener('pause', () => {
    localStorage.setItem(`bookbankPlaybackLocation-${id}`, audio.currentTime);
});

