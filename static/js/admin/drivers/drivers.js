function buildDriverTable() {
    const driver_list = document.getElementById('driver-list')
    for (const index in drivers) {
        const driver = drivers[index]
        driver_list.innerHTML += `
        <tr>
            <td class="driver-name">${driver['name']}</td>
            <td class="driver-points">
                <input 
                    class="form-control bg-dark border-secondary text-white text-center" 
                    type="number" 
                    value="${driver['points']}" 
                    id="points-${driver['id']}"
                    disabled
                >
            </td>
            <td class="driver-penalty-points">
                <input 
                    class="form-control bg-dark border-secondary text-white text-center" 
                    type="number" 
                    value="${driver['penalty-points']}" 
                    id="penalty-points-${driver['id']}"
                    disabled
                >
            </td>
            <td>
                <button 
                class="btn btn-danger driver-del" 
                data-bs-toggle="modal"
                data-bs-target="#deleteDriverModal"
                onclick="deleteModal('${driver['name']}', '${driver['id']}')" 
                hidden>
                    Delete
                </button>
            </td>
        </tr>
        `
    }
}

function deleteDriver(id) {
    document.getElementById('delete-success').removeAttribute('hidden')

    const del_url = '/api/drivers/delete?id=' + id + '&ip=' + user_ip
    fetch(del_url)
        .then((response) => {
            return response.json()
        }).then((data) => {
        console.log(data)
    })

    deleteTeamDriver(id)
    setTimeout(() => {location.reload()}, 1000)
}

function deleteTeamDriver(id) {
    for (const i in teams) {
        const team = teams[i]

        const update_url = `/api/teams/update?ip=${user_ip}&id=${team['id']}`
        if (team['driver1'].toString() === id) {
            fetch(update_url + `&driver1=-1`)
                .then((response) => {
                return response.json()
            }).then((data) => {
                console.log(data)
            })
            return
        }

        if (team['driver2'].toString() === id) {
            fetch(update_url + `&driver2=-1`)
                .then((response) => {
                    return response.json()
                }).then((data) => {
                console.log(data)
            })
            return
        }
    }
}