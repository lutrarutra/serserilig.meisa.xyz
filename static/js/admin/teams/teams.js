function buildDriverTable() {
    const driver_list = document.getElementById('driver-list')
    const reserve_driver_cache = []
    for (const index in drivers) {
        const driver = drivers[index]
        if (driver['team-id'] !== -1) { continue }
        driver_list.innerHTML += `
        <tr>
            <td class="driver-name">
                <div class="driver-box reserve-driver">
                    ${driver['name']}
                </div>
            </td>
        </tr>
        `
        reserve_driver_cache.push(driver['name'])
    }
    return reserve_driver_cache
}

function buildTeamTable() {
    const team_list = document.getElementById('team-list')
    for (const index in teams) {
        const team = teams[index]

        team_list.innerHTML += `
                <tr>
                    <td class="team-name align-middle">
                        <div style="border-left: 3px solid ${team['color']}; height: 1.5em; margin: 0 2em">
                            ${team['name']}            
                        </div>
                    </td>
                    <td class="team-points">
                        <input 
                            class="form-control bg-dark border-secondary text-white text-center" 
                            type="number" 
                            value="${team['points']}" 
                            id="points-${team['id']}"
                            disabled
                        >
                    </td>
                    <td class="team-driver1" style="border-left: 2px solid white">
                        <div class="driver-box" id="team-${team['id']}-driver1">
                            ${driver_ids[team['driver1']]}
                        </div>
                    </td>
                    <td class="team-driver2"><div class="driver-box" id="team-${team['id']}-driver2">${driver_ids[team['driver2']]}</div></td>
                </tr>
            `
    }
    greyOutNames()
}

function updateTeamDrivers() {
    for (const i in teams) {
        const team = teams[i]
        const seats = ['1', '2']
        for (const j in seats) {
            const seat = `driver${seats[j]}`
            const new_driver = document.getElementById(`team-${team['id']}-${seat}`).innerHTML.trim()
            const new_driver_id = driver_names[new_driver]
            const prev_driver = driver_ids[team[seat]]
            if (new_driver === prev_driver) { continue }

            console.log(`${team['name']} ${seat}: ${prev_driver} -> ${new_driver}`)

            const driver_update_url = `/api/drivers/update?ip=${user_ip}`
            // Update new driver team
            if (new_driver !== 'No Driver') {
                const new_driver_team = `${driver_update_url}&id=${new_driver_id}&team_id=${team['id']}`
                fetch(new_driver_team).then((response) => {return response.json()}).then((data) => {console.log(data)})
            }
            // Update team driver
            const new_team_driver = `/api/teams/update?ip=${user_ip}&id=${team['id']}&${seat}=${new_driver_id}`
            fetch(new_team_driver).then((response) => {return response.json()}).then((data) => {console.log(data)})
        }
    }
    updateReserveDrivers()
}

function updateTeamPoints() {
    for (const i in teams) {
        const team = teams[i]
        const points_input = document.getElementById(`points-${team['id']}`)

        if (points_input.value !== team['points'].toString()) {
            const update_url = `/api/teams/update?ip=${user_ip}&id=${team['id']}&points=${points_input.value}`
            fetch(update_url).then((response) => {return response.json()}).then((data) => {console.log(data)})
        }
    }
}

function updateReserveDrivers() {
    document.querySelectorAll('.reserve-driver')
        .forEach(function (reserve_slot) {
            const driver_name = reserve_slot.innerHTML.trim()
            if (!(reserve_driver_cache.includes(driver_name)) && driver_name !== 'No Driver') {
                const driver_id = driver_names[driver_name]
                const update_url = `/api/drivers/update?ip=${user_ip}&id=${driver_id}&team_id=-1`
                console.log(update_url)
                fetch(update_url).then((response) => {return response.json()}).then((data) => {console.log(data)})
            }
        })

    setTimeout(() => {location.reload()}, 2000)
}

function showSaveBanner() {
    document.getElementById('saved-alert').removeAttribute('hidden')
}