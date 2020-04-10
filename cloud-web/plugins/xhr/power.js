import _pack from './_pack'

export default function({ $axios }) {
  const pack = _pack({ $axios })

  function getPowers(page, pageSize) {
    return pack({
      url: apiRoot + '/powers',
      method: 'get',
      params: {
        page,
        pageSize
      }
    })
  }

  function modifyPower(ID, Name) {
    return pack({
      url: apiRoot + '/power/' + ID,
      method: 'put',
      data: {
        Name
      }
    })
  }

  function newPower(Name) {
    return pack({
      url: apiRoot + '/power',
      method: 'post',
      data: {
        Name
      }
    })
  }

  function delPower(ID) {
    return pack({
      url: apiRoot + '/power/' + ID,
      method: 'delete'
    })
  }

  return { getPowers, modifyPower, newPower, delPower }
}
