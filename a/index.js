const btnUpload = document.getElementById("option-upload");
const btnLoad = document.getElementById("option-load");

btnUpload.addEventListener("click", goToUpload);

/* let btnUpload = document.getElementById("upload");
btnUpload.addEventListener("click", uploadFile);

function uploadFile() {
    let inputFile = document.getElementById("input-file");
    let file = inputFile.files[0];

    let dataForm = new FormData();
    dataForm.append("file", file);

    fetch("/api/v1/upload", {
        method: "POST",
        body: dataForm,
    })
        .then((responsive) => {
            if (responsive.status >= 400) {
                throw true;
            }
        })
        .then(() => {
            let title = document.getElementById("title");
            title.textContent = "Uploaded File";
        })
        .catch(() => {
            let title = document.getElementById("title");
            title.textContent = "Failure to Upload File";
        });
} */
