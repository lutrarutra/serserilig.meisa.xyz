function buildDriverStandings() {
    let driver_standings = document.getElementById('driver-standings')
    let pos_count = 1
    for (const driver_id in drivers) {
        const driver = drivers[driver_id]
        const team = driver['team-id'] === -1 ? {'abbr': 'res', 'color': '#fff'} : teams_by_id[driver['team-id']]
        const border_style = driver['team-id'] === -1 ? 'dotted' : 'solid'
        driver_standings.innerHTML += `
                <tr>
                    <th scope="row" class="driver-pos">${pos_count}</th>
                    <td class="driver-team-color"><div style="border-left: 3px ${border_style} ${team['color']}; height: 2em"></div></td>
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
                        <img src="/static/images/team_logos/${team['abbr']}.png" alt="${team['abbr']}">
                    </td>
                    <td class="team-color"><div style="border-left: 3px solid ${team['color']}; height: 2em"></div></td>
                    <td class="team-name">${team['name']}</td>
                    <td class="team-points">${team['points']}</td>
                </tr>
            `
        pos_count++
    }
}