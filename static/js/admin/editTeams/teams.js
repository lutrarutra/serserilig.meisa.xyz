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
                </tr>
                `
            }
        })
}