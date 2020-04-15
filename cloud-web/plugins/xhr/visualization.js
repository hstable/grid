import _pack from './_pack'

export default function({ $axios }) {
  const pack = _pack({ $axios })

  function getLocations() {
    return pack({
      url: apiRoot + `/visualization/locations`,
      method: 'get'
    })
  }
  function getPowerCounts() {
    return pack({
      url: apiRoot + `/visualization/powerCounts`,
      method: 'get'
    })
  }
  function getCompanyDevices() {
    return pack({
      url: apiRoot + `/visualization/companyDevices`,
      method: 'get'
    })
  }
  function getMarks() {
    return pack({
      url: apiRoot + `/visualization/marks`,
      method: 'get'
    })
  }

  return { getLocations, getPowerCounts, getCompanyDevices, getMarks }
}
