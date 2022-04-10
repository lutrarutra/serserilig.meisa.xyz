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
                    <p class="team-driver">${driver1['name']}</p>
                    <p class="team-driver">${driver2['name']}</p>
                </td>
                <td class="team-driver-pp">
                    <p class="driver-pp">${driver1['penalty-points']}</p>
                    <p class="driver-pp">${driver2['penalty-points']}</p>
                </td>
            </tr>
        `
    }
    greyOutNames()
    greyOutPP()
}

function greyOutNames() {
    document.querySelectorAll('.team-driver')
        .forEach(function (value){
            const name = value.innerHTML
            if (name === 'No Driver') {
                value.style.color = '#888'
            }
        })
}

function greyOutPP() {
    document.querySelectorAll('.driver-pp')
        .forEach(function (value){
            const name = value.innerHTML
            if (name === '0') {
                value.style.color = '#888'
            }
        })
}