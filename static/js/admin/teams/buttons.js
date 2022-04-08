function enableBoxes() {
    let items = document.querySelectorAll('.driver-box')
    for (const index in items) {
        const item = items[index]

        if (typeof item === 'object') {
            item.classList.add('box')
            item.setAttribute('draggable', 'true')
        }
    }
}

function enableInputs() {
    document.querySelectorAll('table input')
        .forEach(function (input) {
            input.removeAttribute('disabled')
        })
}