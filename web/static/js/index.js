var navbar = document.getElementById("navbar");
var body = document.getElementsByTagName("body")[0];
var main = document.getElementsByTagName("main")[0];
var menuBtn = document.getElementById("menu-open-btn");
var maskOverlay = document.getElementById("mask-overlay");

async function doNavbar() {
    if (navbar.style.display == "flex") {
        maskOverlay.style.display = "none";
        menuBtn.style.display = "flex";
        navbar.classList.remove("animation-open");
        navbar.classList.add("animation-close");
        await new Promise(r => setTimeout(r, 500));
        navbar.style.display = "none";
        maskOverlay.removeEventListener("click", doNavbar);
    } else {
        menuBtn.style.display = "none"
        navbar.style.display = "flex";
        navbar.classList.remove("animation-close");
        navbar.classList.add("floating-navbar");
        navbar.classList.add("animation-open");
        maskOverlay.style.display = "block";
        await new Promise(r => setTimeout(r, 200));
        maskOverlay.addEventListener("click", doNavbar);
    }
}

// Prevent event bubbling from navbar and menu button to overlay
// navbar.addEventListener("click", function (e) { e.stopPropagation(); });
// if (menuBtn) menuBtn.addEventListener("click", function (e) { e.stopPropagation(); });

async function version() {
    const response = await fetch("/api/version");
    const data = await response.text();
    document.getElementById("version").innerHTML = "V" + data;
    document.getElementById("version").innerHTML = document.getElementById("version").innerHTML.replace(" ", "").replace("\n", "");
}
version()