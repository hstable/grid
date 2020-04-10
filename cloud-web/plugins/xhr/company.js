import _pack from './_pack'

export default function({ $axios }) {
  const pack = _pack({ $axios })

  function getCompanies(page, pageSize) {
    return pack({
      url: apiRoot + '/companies',
      method: 'get',
      params: {
        page,
        pageSize
      }
    })
  }

  function modifyCompany(ID, Name) {
    return pack({
      url: apiRoot + '/company/' + ID,
      method: 'put',
      data: {
        Name
      }
    })
  }

  function newCompany(Name) {
    return pack({
      url: apiRoot + '/company',
      method: 'post',
      data: {
        Name
      }
    })
  }

  function delCompany(ID) {
    return pack({
      url: apiRoot + '/company/' + ID,
      method: 'delete'
    })
  }

  return { getCompanies, modifyCompany, newCompany, delCompany }
}
