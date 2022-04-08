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

    const team_resp = await fetch('/api/teams')
    const teams = await team_resp.json()
    for (const index in teams) {
        const team = teams[index]

        team_list.innerHTML += `
                <tr>
                    <td class="team-name align-middle">
                        <div style="border-left: 3px solid ${team['color']}; height: 1.7em; margin: 0 2em">
                            ${team['name']}            
                        </div>
                    </td>
                    <td class="team-points">${team['points']}</td>
                    <td class="team-driver1"><div class="driver-box" id="team-${team['id']}-driver1">${drivers[team['driver1']]}</div></td>
                    <td class="team-driver2"><div class="driver-box" id="team-${team['id']}-driver2">${drivers[team['driver2']]}</div></td>
                </tr>
            `
    }
}

async function getAllDrivers() {
    const response = await fetch('/api/drivers')
    const drivers = await response.json()

    let driver_ids = {
        '-1': 'No Driver'
    }

    for (const index in await drivers) {
        const driver = drivers[index]
        driver_ids[driver['id'].toString()] = driver['name']
    }

    return driver_ids
}