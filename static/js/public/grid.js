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

    team_table.innerHTML += `
        <tr>
            <td class="team-logo"></td>
            <td class="team-color"><div id="reserve-border" style="border-left: 3px dotted white; height: 2em"></div></td>
            <td class="team-name">Reserve</td>
            <td class="team-drivers" id="reserve-drivers"></td>
            <td class="team-drivers-pp" id="reserve-drivers-pp"></td>
        </tr>
    `

    const reserve_spot = document.getElementById('reserve-drivers')
    const reserve_pp = document.getElementById('reserve-drivers-pp')
    let reserve_count = 0
    for (const i in drivers) {
        const driver = drivers[i]
        if (driver['team-id'] !== -1) {
            continue
        }

        reserve_spot.innerHTML += `<p class="team-driver">${driver['name']}</p>`
        reserve_pp.innerHTML += `<p class="driver-pp">${driver['penalty-points']}</p>`
        reserve_count++
    }

    document.getElementById('reserve-border').style.height = `${reserve_count * 1.5}em`

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