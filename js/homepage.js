const navbar = document.getElementById("navbar");
navbar.addEventListener("click", (event) => {
    if (event.target.id !== "navbar") {
        let data = {Content: event.target.id,};
        fetch("/get_content", {
            method: "POST", body: JSON.stringify(data)
        }).then((response) => {
            response.text().then(function (data) {
                let result = JSON.parse(data);
                document.getElementById("content").innerHTML = result["Content"]
                addInteraction(event.target.id)
            });

        }).catch((error) => {
            console.log(error)
        });
    }
})

function addInteraction(content) {
    if (content === "first_tab") {
        firstTabInteraction();
    } else if (content === "second_tab") {
        secondTabInteraction();
    }
}

function firstTabInteraction() {
    const firstButton = document.getElementById("first_button")
    firstButton.addEventListener("click", (event) => {
        let actualData = document.getElementById("first_data")
        actualData.textContent = +actualData.textContent + 1
    })
}

function secondTabInteraction() {
    const secondButton = document.getElementById("second_button")
    secondButton.addEventListener("click", (event) => {
        let actualData = document.getElementById("second_data")
        actualData.textContent = +actualData.textContent - 1
    })
}

