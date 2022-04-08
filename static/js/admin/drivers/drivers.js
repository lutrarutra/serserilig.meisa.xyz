function buildDriverTable() {
    const driver_list = document.getElementById('driver-list')
    for (const index in drivers) {
        const driver = drivers[index]
        driver_list.innerHTML += `
        <tr>
            <td class="driver-name">${driver['name']}</td>
            <td class="driver-points">
                <input 
                    class="form-control bg-dark border-secondary text-white text-center" 
                    type="number" 
                    value="${driver['points']}" 
                    id="points-${driver['id']}"
                    disabled
                >
            </td>
            <td class="driver-penalty-points">
                <input 
                    class="form-control bg-dark border-secondary text-white text-center" 
                    type="number" 
                    value="${driver['penalty-points']}" 
                    id="penalty-points-${driver['id']}"
                    disabled
                >
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
}