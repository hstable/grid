import _pack from './_pack'
import { getUserInfoFromToken } from '@/assets/js/tokenTools'

export default function({ $axios }) {
  const pack = _pack({ $axios })

  function getUsers(page, pageSize) {
    return pack({
      url: apiRoot + '/users',
      method: 'get',
      params: {
        page,
        pageSize
      }
    })
  }

  function modifyUser(
    ID,
    Sub,
    Password,
    Name,
    CompanyID,
    Department,
    PhoneNumber,
    WechatID,
    EnterpriseID,
    Admin
  ) {
    return pack({
      url: apiRoot + '/user/' + ID,
      method: 'put',
      data: {
        Sub,
        Password,
        Name,
        CompanyID,
        Department,
        PhoneNumber,
        WechatID,
        EnterpriseID,
        Admin
      }
    })
  }

  function newUser(
    Sub,
    Password,
    Name,
    CompanyID,
    Department,
    PhoneNumber,
    WechatID,
    EnterpriseID,
    Admin
  ) {
    return pack({
      url: apiRoot + '/user',
      method: 'post',
      data: {
        Sub,
        Password,
        Name,
        CompanyID,
        Department,
        PhoneNumber,
        WechatID,
        EnterpriseID,
        Admin
      }
    })
  }

  function delUser(ID) {
    return pack({
      url: apiRoot + '/user/' + ID,
      method: 'delete'
    })
  }

  function login(Username, Password) {
    return pack({
      url: apiRoot + '/user/login',
      method: 'post',
      data: {
        Username,
        Password
      }
    }).then((res) => {
      localStorage.token = res.data.data.token
      const o = getUserInfoFromToken(res.data.data.token)
      localStorage.name = o.name
      localStorage.sub = o.sub
    })
  }

  return { getUsers, modifyUser, newUser, delUser, login }
}
