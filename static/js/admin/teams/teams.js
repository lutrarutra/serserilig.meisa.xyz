function updateTeamDrivers() {
    for (const i in teams) {
        const team = teams[i]
        const seats = ['1', '2']
        for (const j in seats) {
            const seat = `driver${seats[j]}`
            const new_driver = document.getElementById(`team-${team['id']}-${seat}`).innerHTML.trim()
            const new_driver_id = driver_ids[new_driver]
            const prev_driver = driver_names[team[seat]]
            if (new_driver === prev_driver) { continue }

            // console.log(`${team['name']} ${seat}: ${prev_driver} -> ${new_driver}`)

            const driver_update_url = `/api/drivers/update?ip=${user_ip}`
            // Update new driver team
            if (new_driver !== 'No Driver') {
                const new_driver_team = `${driver_update_url}&id=${new_driver_id}&team_id=${team['id']}`
                console.log(new_driver_team)
                fetch(new_driver_team).then((response) => {return response.json()}).then((data) => {console.log(data)})
            }
            // Update team driver
            const new_team_driver = `/api/teams/update?ip=${user_ip}&id=${team['id']}&${seat}=${new_driver_id}`
            console.log(new_team_driver)
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
                const driver_id = driver_ids[driver_name]
                const update_url = `/api/drivers/update?ip=${user_ip}&id=${driver_id}&team_id=-1`
                console.log(update_url)
                fetch(update_url).then((response) => {return response.json()}).then((data) => {console.log(data)})
            }
        })

    // setTimeout(() => {location.reload()}, 2000)
}

function showSaveBanner() {
    document.getElementById('saved-alert').removeAttribute('hidden')
}