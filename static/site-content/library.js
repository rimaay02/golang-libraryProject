const form = document.querySelector(".new-book");
const table = document.querySelector("table");
const btn = document.querySelector("button");

let idValue = document.getElementById("id");
let titleValue = document.getElementById("title");
let authorValue = document.getElementById("author");

async function doRequest() {
    let url = 'http://localhost:8080/';
    let res = await fetch(url);

    if (res.ok) {

        let text = await res.text();

        return text;
    } else {
        return `HTTP error: ${res.status}`;
    }
}

doRequest().then(response => {
    var obj = JSON.parse(response);
    obj.forEach(bookObj => {
        let id = bookObj['id'];
        let title = bookObj['title'];
        let author = bookObj['author'];
        table.innerHTML += `<tr>
            <td>${id}</td>
            <td>${title}</td>
            <td>${author}</td>           
        </tr>`;
     });
    
})

btn.onclick = function() {
    let request = {
        id: idValue.value,
        title: titleValue.value,
        author: authorValue.value
    }
    fetch("http://localhost:8080/books", {
        method: "POST",
        body: JSON.stringify(request),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
    });
    location.reload();

}