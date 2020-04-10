import _pack from './_pack'

export default function({ $axios }) {
  const pack = _pack({ $axios })

  function getTowers(page, pageSize, lineID) {
    return pack({
      url: apiRoot + `/line/${lineID}/towers`,
      method: 'get',
      params: {
        page,
        pageSize
      }
    })
  }

  function modifyTower(ID, Name, LineID, BaseStationID, Location) {
    return pack({
      url: apiRoot + '/tower/' + ID,
      method: 'put',
      data: {
        Name,
        LineID,
        BaseStationID,
        Location
      }
    })
  }

  function newTower(Name, lineID, BaseStationID, Location) {
    return pack({
      url: apiRoot + `/line/${lineID}/tower`,
      method: 'post',
      data: {
        Name,
        BaseStationID,
        Location
      }
    })
  }

  function delTower(ID) {
    return pack({
      url: apiRoot + '/tower/' + ID,
      method: 'delete'
    })
  }

  return { getTowers, modifyTower, newTower, delTower }
}
