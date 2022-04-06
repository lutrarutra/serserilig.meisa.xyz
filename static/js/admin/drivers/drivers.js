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
                    <td class="driver-penalty-points">
                        <div class="btn-group">
                            0
                            <button class="btn btn-sm btn-light ms-4" hidden>-1</button>
                            <button class="btn btn-sm btn-light" hidden>+1</button>
                        </div>
                    </td>
                    <td><button class="btn btn-danger driver-del" hidden>Delete</button></td>
                </tr>
                `
            }
        })
}