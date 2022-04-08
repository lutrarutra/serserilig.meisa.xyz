function buildGridTable() {
    const team_table = document.getElementById('team-list')
    for (const i in teams) {
        const team = teams[i]
        const driver1 = driver_ids[team['driver1']]
        const driver2 = driver_ids[team['driver2']]

        team_table.innerHTML += `
            <tr>
                <td class="team-logo">
                    <img src="/static/images/team_logos/${team['abbr']}.png" alt="${team['abbr']}">
                </td>
                <td class="team-color"><div style="border-left: 3px solid ${team['color']}; height: 2em"></div></td>
                <td class="team-name">${team['name']}</td>
                <td class="team-drivers">
                    ${driver1['name']}
                    <br>
                    ${driver2['name']}
                </td>
                <td class="team-driver-pp">
                    ${driver1['penalty-points']}
                    <br>
                    ${driver2['penalty-points']}
                </td>
            </tr>
        `
    }
}