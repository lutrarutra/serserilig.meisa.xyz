function addDriver() {
    const name_input = document.getElementById('add-driver-name')
    const valid_input = validateInput(name_input)

    const success_alert = document.getElementById('add-success')
    if (valid_input) {
        success_alert.removeAttribute('hidden')

        const add_url = '/api/drivers/add?name=' + name_input.value + '&ip=' + user_ip
        fetch(add_url)
            .then((response) => {
                return response.json()
            }).then((data) => {
            success_alert.innerHTML = data
        })

        setTimeout(() => {location.reload()}, 2000)
    }
}

function deleteModal(driver_name, driver_id) {
    document.getElementById('deleteDriverModalText').innerHTML =
        `Are you sure you want to delete driver ${driver_name}?`

    document.getElementById('deleteDriverModalButtons').innerHTML +=
        `<button type="button" class="btn btn-danger" onclick="deleteDriver('${driver_id}')">Delete</button>`
}

function deleteDriver(id) {
    document.getElementById('delete-success').removeAttribute('hidden')

    const add_url = '/api/drivers/delete?id=' + id + '&ip=' + user_ip
    fetch(add_url)
        .then((response) => {
            return response.json()
        }).then((data) => {
        console.log(data)
        })

    setTimeout(() => {location.reload()}, 2000)
}

function toggleEditButtons() {
    document.querySelectorAll('tr button')
        .forEach(function (button) {
            if (button.hasAttribute('hidden')) {
                button.removeAttribute('hidden')
            } else {
                button.setAttribute('hidden', '')
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
