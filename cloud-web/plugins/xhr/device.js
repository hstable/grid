import _pack from './_pack'

export default function({ $axios }) {
  const pack = _pack({ $axios })

  function getDevices(page, pageSize, towerID) {
    return pack({
      url: apiRoot + `/tower/${towerID}/devices`,
      method: 'get',
      params: {
        page,
        pageSize
      }
    })
  }

  function modifyDevice(
    ID,
    Name,
    TowerID,
    Manufacture,
    ModelName,
    SerialNumber
  ) {
    return pack({
      url: apiRoot + '/device/' + ID,
      method: 'put',
      data: {
        Name,
        TowerID,
        Manufacture,
        ModelName,
        SerialNumber
      }
    })
  }

  function newDevice(Name, TowerID, Manufacture, ModelName, SerialNumber) {
    return pack({
      url: apiRoot + `/tower/${TowerID}/device`,
      method: 'post',
      data: {
        Name,
        TowerID,
        Manufacture,
        ModelName,
        SerialNumber
      }
    })
  }

  function delDevice(ID) {
    return pack({
      url: apiRoot + '/device/' + ID,
      method: 'delete'
    })
  }

  function changeDeviceStatus(deviceID, newStatus) {
    return pack({
      url: apiRoot + `/device/${deviceID}/disabled`,
      method: 'put',
      data: {
        Disabled: !newStatus
      }
    })
  }

  return { getDevices, modifyDevice, newDevice, delDevice, changeDeviceStatus }
}
