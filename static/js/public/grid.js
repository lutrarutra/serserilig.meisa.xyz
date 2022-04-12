function greyOutNames() {
    document.querySelectorAll('.team-driver')
        .forEach(function (value){
            const name = value.innerHTML
            if (name === 'No Driver') {
                value.style.color = '#888'
            }
        })
}

function greyOutPP() {
    document.querySelectorAll('.driver-pp')
        .forEach(function (value){
            const name = value.innerHTML
            if (name === '0') {
                value.style.color = '#888'
            }
        })
}