function deleteDriver(id) {
    document.getElementById('delete-success').removeAttribute('hidden')

    const del_url = '/api/drivers/delete?id=' + id + '&ip=' + user_ip
    fetch(del_url)
        .then((response) => {
            return response.json()
        }).then((data) => {
        console.log(data)
    })

    deleteTeamDriver(id)
    setTimeout(() => {location.reload()}, 1000)
}

function deleteTeamDriver(id) {
    for (const i in teams) {
        const team = teams[i]

        const update_url = `/api/teams/update?ip=${user_ip}&id=${team['id']}`
        if (team['driver1'].toString() === id) {
            fetch(update_url + `&driver1=-1`)
                .then((response) => {
                return response.json()
            }).then((data) => {
                console.log(data)
            })
            return
        }

        if (team['driver2'].toString() === id) {
            fetch(update_url + `&driver2=-1`)
                .then((response) => {
                    return response.json()
                }).then((data) => {
                console.log(data)
            })
            return
        }
    }
}