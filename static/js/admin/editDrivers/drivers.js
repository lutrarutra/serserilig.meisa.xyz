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
                    <td class="driver-name">${driver['name']}</td>
                    <td class="driver-points">${driver['points']}</td>
                    <td class="driver-penalty-points">0</td>
                </tr>
                `
            }
        })
}