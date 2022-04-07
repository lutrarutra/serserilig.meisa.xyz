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
                        <div class="input-group">
                            <input 
                                class="form-control bg-dark border-secondary text-white text-center" 
                                type="number" 
                                value="${driver['penalty-points']}" 
                                id="penalty-points-${driver['id']}"
                                disabled
                            >
                        </div>
                    </td>
                    <td>
                        <button 
                        class="btn btn-danger driver-del" 
                        data-bs-toggle="modal"
                        data-bs-target="#deleteDriverModal"
                        onclick="deleteModal('${driver['name']}', '${driver['id']}')" 
                        hidden>
                            Delete
                        </button>
                    </td>
                </tr>
                `
            }
        })
}