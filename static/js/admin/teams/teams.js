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
                if (driver['team-id'] !== -1) { continue }
                driver_list.innerHTML += `
                <tr>
                    <td class="driver-name">
                        <div class="driver-box">
                            ${driver['name']}
                        </div>
                    </td>
                </tr>
                `
            }
        })
}

async function buildTeamTable() {
    const team_list = document.getElementById('team-list')

    const drivers = await getAllDrivers()
    console.log(drivers)

    const team_resp = await fetch('/api/teams')
    const teams = await team_resp.json()
    for (const index in teams) {
        const team = teams[index]

        const driver1_name = typeof drivers[team['driver1']] === 'undefined' ? 'No Driver' : drivers[team['driver1']]
        const driver2_name = typeof drivers[team['driver2']] === 'undefined' ? 'No Driver' : drivers[team['driver2']]

        team_list.innerHTML += `
                <tr>
                    <td class="team-name">${team['name']}</td>
                    <td class="team-points">${team['points']}</td>
                    <td class="team-driver1"><div class="driver-box" id="driver-box">${driver1_name}</div></td>
                    <td class="team-driver2"><div class="driver-box" id="driver-box">${driver2_name}</div></td>
                </tr>
            `
    }
}

async function getAllDrivers() {
    const response = await fetch('/api/drivers')
    const drivers = await response.json()

    let driver_ids = {}

    for (const index in await drivers) {
        const driver = drivers[index]
        driver_ids[driver['id'].toString()] = driver['name']
    }

    return driver_ids
}