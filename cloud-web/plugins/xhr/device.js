import { ToastProgrammatic } from 'buefy'
import _pack from './_pack'
import { toInt } from '@/assets/js/utils'

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
    SerialNumber,
    Network
  ) {
    return pack({
      url: apiRoot + '/device/' + ID,
      method: 'put',
      data: {
        Name,
        TowerID: toInt(TowerID),
        Manufacture,
        ModelName,
        SerialNumber,
        Network
      }
    })
  }

  function newDevice(
    Name,
    TowerID,
    Manufacture,
    ModelName,
    SerialNumber,
    Network
  ) {
    return pack({
      url: apiRoot + `/tower/${TowerID}/device`,
      method: 'post',
      data: {
        Name,
        TowerID: toInt(TowerID),
        Manufacture,
        ModelName,
        SerialNumber,
        Network
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

  function newPhoto(deviceID) {
    return new Promise((resolve, reject) => {
      $axios({
        url: apiRoot + `/device/${deviceID}/newPhoto`,
        method: 'get'
      })
        .then((res) => {
          resolve(res)
        })
        .catch((err) => {
          function notsuretip() {
            ToastProgrammatic.open({
              message: '请求服务器时发生网络错误',
              type: 'is-warning',
              position: 'is-top',
              queue: false,
              duration: 5000
            })
          }
          if (!err.response) {
            notsuretip()
          } else {
            switch (err.response.status) {
              case 401:
                break
              case 400:
                ToastProgrammatic.open({
                  message: err.response.data,
                  type: 'is-warning',
                  position: 'is-top',
                  queue: false,
                  duration: 5000
                })
                break
              default:
                notsuretip()
            }
          }
          reject(err)
        })
    })
  }

  return {
    getDevices,
    modifyDevice,
    newDevice,
    delDevice,
    changeDeviceStatus,
    newPhoto
  }
}
