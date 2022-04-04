function buildDriverStandings() {
    let driver_standings = document.getElementById('driver-standings')
    let pos_count = 1
    for (const driver_id in drivers) {
        const driver = drivers[driver_id]
        driver_standings.innerHTML += `
                <tr>
                    <th scope="row" class="driver-pos">${pos_count}</th>
                    <td class="driver-team-logo">
                        <img src="https://www.formula1.com/content/dam/fom-website/teams/2022/${driver['team']}-logo.png.transform/2col/image.png" 
                        alt="${driver['team_abbr']}">
                    </td>
                    <td class="driver-team-color"><div class="color-${driver['team']}"></div></td>
                    <td class="driver-name">${driver['name']}</td>
                    <td class="driver-points">${driver['points']}</td>
                </tr>
            `
        pos_count++
    }
}

function buildTeamStandings() {
    let team_standings = document.getElementById('team-standings')
    let pos_count = 1
    for (const team_name in teams) {
        const team = teams[team_name]
        team_standings.innerHTML += `
                    <tr>
                        <th scope="row" class="team-pos">${pos_count}</th>
                        <td class="team-logo">
                            <img src="https://www.formula1.com/content/dam/fom-website/teams/2022/${team_name}-logo.png.transform/2col/image.png" 
                            alt="${team['abbr']}">
                        </td>
                        <td class="team-color"><div class="color-${team_name}"></div></td>
                        <td class="team-name">${team['name']}</td>
                        <td class="team-points">${team['points']}</td>
                    </tr>
        `
        pos_count++
    }
}