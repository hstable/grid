import _pack from './_pack'

export default function({ $axios }) {
  const pack = _pack({ $axios })

  function getBaseStations(page, pageSize) {
    return pack({
      url: apiRoot + '/baseStations',
      method: 'get',
      params: {
        page,
        pageSize
      }
    })
  }

  function modifyBaseStation(ID, Name, IP, Location) {
    return pack({
      url: apiRoot + '/baseStation/' + ID,
      method: 'put',
      data: {
        Name,
        IP,
        Location
      }
    })
  }

  function newBaseStation(Name, IP, Location) {
    return pack({
      url: apiRoot + '/baseStation',
      method: 'post',
      data: {
        Name,
        IP,
        Location
      }
    })
  }

  function delBaseStation(ID) {
    return pack({
      url: apiRoot + '/baseStation/' + ID,
      method: 'delete'
    })
  }

  return {
    getBaseStations,
    modifyBaseStation,
    newBaseStation,
    delBaseStation
  }
}
