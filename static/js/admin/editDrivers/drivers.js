function addDriver() {
    const name_input = document.getElementById('add-driver-name')
    const valid_input = validateInput(name_input)

    const success_alert = document.getElementById('add-success')
    if (valid_input) {
        success_alert.removeAttribute('hidden')

        const add_url = '/api/drivers/add?name=' + name_input.value
        fetch(add_url)
            .then((response) => {
            return response.json()
            }).then((data) => {
                console.log(data)
            })
        
        setTimeout(() => {location.reload()}, 2000)
    }
}

function buildDriverTable() {
    const driver_list = document.getElementById('driver-list')
    const url = '/api/drivers'
    fetch(url)
        .then((response) => {
            return response.json();
        })
        .then((data) => {
            const drivers = data
            for (const index in drivers) {
                const driver = drivers[index]
                driver_list.innerHTML += `
                <tr>
                    <td class="driver-name">${driver['name']}</td>
                    <td class="driver-points">${driver['points']}</td>
                    <td class="driver-penalty-points">0</td>
                </tr>
                `
            }
        })
}

function validateInput(input_element) {
    if (input_element.value.length < 1) {
        console.log('invalid')
        input_element.classList.remove('border-success')
        input_element.classList.add('border-danger')
        return false
    }

    console.log('valid')
    input_element.classList.remove('border-danger')
    input_element.classList.add('border-success')
    return true
}