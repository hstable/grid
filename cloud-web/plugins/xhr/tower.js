import dayjs from 'dayjs'
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

  // example: getTowerMonthMarks(1, "2020-03")
  function getTowerMonthMarks(ID, date) {
    date = dayjs(date).format('YYYY-MM')
    return pack({
      url: apiRoot + '/marks/tower/' + ID,
      method: 'get',
      params: { date }
    })
  }

  return { getTowers, modifyTower, newTower, delTower, getTowerMonthMarks }
}
