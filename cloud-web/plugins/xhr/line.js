import _pack from './_pack'

export default function({ $axios }) {
  const pack = _pack({ $axios })

  function getLines(page, pageSize) {
    return pack({
      url: apiRoot + '/lines',
      method: 'get',
      params: {
        page,
        pageSize
      }
    })
  }

  function modifyLine(ID, Name, CompanyID, PowerID) {
    return pack({
      url: apiRoot + '/line/' + ID,
      method: 'put',
      data: {
        Name,
        CompanyID,
        PowerID
      }
    })
  }

  function newLine(Name, CompanyID, PowerID) {
    return pack({
      url: apiRoot + '/line',
      method: 'post',
      data: {
        Name,
        CompanyID,
        PowerID
      }
    })
  }

  function delLine(ID) {
    return pack({
      url: apiRoot + '/line/' + ID,
      method: 'delete'
    })
  }

  return { getLines, modifyLine, newLine, delLine }
}
