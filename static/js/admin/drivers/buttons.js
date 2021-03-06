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

        setTimeout(() => {
            name_input.value = ''
            location.reload()
        }, 500)
    }
}

function deleteModal(driver_name, driver_id) {
    document.getElementById('deleteDriverModalText').innerHTML =
        `Are you sure you want to delete driver ${driver_name}?`

    document.getElementById('deleteDriverModalButtons').innerHTML = `
            <button type="button" class="btn btn-primary" data-bs-dismiss="modal">Dismiss</button>
            <button type="button" class="btn btn-danger" onclick="deleteDriver('${driver_id}')">Delete</button>
        `
}

function saveEdits() {
    const url = '/api/drivers'
    fetch(url)
        .then((response) => {
            return response.json()
        })
        .then((data) => {
            const drivers = data
            for (const index in drivers) {
                const driver = drivers[index]
                const driver_new_pp = document.getElementById(`penalty-points-${driver['id']}`)
                const driver_new_points = document.getElementById(`points-${driver['id']}`)

                const base_update_url = `/api/drivers/update?ip=${user_ip}&id=${driver['id']}`
                const pp_update_url = `${base_update_url}&penalty_points=${driver_new_pp.value}`
                const points_update_url = `${base_update_url}&points=${driver_new_points.value}`
                fetch(pp_update_url).then((response) => {console.log(response.text())})
                fetch(points_update_url).then((response) => {console.log(response.text())})
            }
        })

    document.getElementById('saved-alert').removeAttribute('hidden')
    setTimeout(() => {location.reload()}, 1000)
}

function toggleEditTools() {
    document.querySelectorAll('tr button')
        .forEach(function (button) {
            if (button.hasAttribute('hidden')) {
                button.removeAttribute('hidden')
            } else {
                button.setAttribute('hidden', '')
            }
        })
    
    const save_button = document.getElementById('save-button')
    if (save_button.hasAttribute('hidden')) {
        save_button.removeAttribute('hidden')
    } else {
        save_button.setAttribute('hidden', '')
    }

    document.querySelectorAll('input')
        .forEach(function (input){
            if (input.hasAttribute('disabled')) {
                input.removeAttribute('disabled')
            } else {
                input.setAttribute('disabled', '')
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
