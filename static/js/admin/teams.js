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

function buildTeamTable() {
    const team_list = document.getElementById('team-list')
    const url = '/api/teams'
    fetch(url)
        .then((response) => {
            return response.json();
        })
        .then((data) => {
            const teams = data
            for (const index in teams) {
                const team = teams[index]
                team_list.innerHTML += `
                    <tr>
                        <td class="team-name">${team['name']}</td>
                        <td class="team-points">${team['points']}</td>
                        <td class="team-driver1"><div class="driver-box" id="driver-box">NULL</div></td>
                        <td class="team-driver2"><div class="driver-box" id="driver-box">NULL</div></td>
                    </tr>
                `
            }
        })
}