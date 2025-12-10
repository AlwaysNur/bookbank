const inputEl = document.getElementById('title-input');
const suggBox = document.getElementById('title-suggestions');

function truncateString(str, maxLength) {
    if (!str) {
        return;
    }
    if (str.length > maxLength) {
        return str.slice(0, maxLength - 3) + "..."; // Subtract 3 for the ellipsis
    }
    return str;
}


async function getSuggestions(title) {
    const response = await fetch(`https://www.googleapis.com/books/v1/volumes?q=${encodeURIComponent(title)}`);
    const data = await response.json();
    return data.items || [];
}

function createEl(className, text = "") {
    const el = document.createElement("span");
    el.classList.add(className);
    el.textContent = text;
    return el;
}

function onTitleInput(event) {
    const title = event.target.value.trim().toLowerCase();
    if (!title || title == "") {
        hideSuggestions();
        return;
    }

    getSuggestions(title)
        .then(suggestions => renderSuggestions(suggestions))
        .catch(console.error);
}

function hideSuggestions() {
    suggBox.setAttribute("hidden", "");
}

function showSuggestions() {
    suggBox.removeAttribute("hidden");
}

function renderSuggestions(suggestions) {
    if (!suggestions || suggestions.length === 0) {
        hideSuggestions();
        return;
    }

    showSuggestions();
    suggBox.innerHTML = "";

    suggestions.forEach((item, index) => {
        const info = item.volumeInfo;

        // container
        const suggestionDiv = document.createElement("div");
        suggestionDiv.classList.add("suggestion-item");
        suggestionDiv.dataset.index = index;

        // thumbnail
        const thumbnail = document.createElement("img");
        thumbnail.classList.add("thumbnail");
        thumbnail.src = info.imageLinks?.thumbnail || info.imageLinks?.smallThumbnail || "/static/image/placeholder.png";
        thumbnail.alt = info.title || "Book cover";

        // text block wrapper
        const textBlock = document.createElement("div");
        textBlock.classList.add("text-block");

        // title
        const title = createEl("title", info.title || "Untitled");

        // author
        const author = createEl("author", info.authors ? info.authors.join(", ") : "Unknown author");

        // description
        var description = createEl("description", truncateString(info.description, "65") || "");

        // series (rare)
        const series = createEl("series", info.series || "");

        // ISBN (hidden)
        const isbn = document.createElement("span");
        isbn.classList.add("isbn");
        isbn.hidden = true;
        isbn.textContent = (
            info.industryIdentifiers?.find(id => id.type.includes("ISBN"))?.identifier
        ) || "";

        // assemble DOM
        textBlock.appendChild(title);
        textBlock.appendChild(author);
        textBlock.appendChild(description);
        textBlock.appendChild(series);
        textBlock.appendChild(isbn);

        suggestionDiv.appendChild(thumbnail);
        suggestionDiv.appendChild(textBlock);
        suggestionDiv.addEventListener("click", function (event) {
            var title = this.getElementsByClassName("title")[0].innerHTML;
            var author = this.getElementsByClassName("author")[0].innerHTML;
            var series = this.getElementsByClassName("series")[0].innerHTML;
            var isbn = this.getElementsByClassName("isbn")[0].innerHTML;

            document.getElementById("title-input").value = title;
            document.querySelector('[name="author"]').value = author;
            document.querySelector('[name="series"]').value = series;
            document.querySelector('[name="isbn"]').value = isbn;
            hideSuggestions();

        });
        suggBox.appendChild(suggestionDiv);
    });
}
