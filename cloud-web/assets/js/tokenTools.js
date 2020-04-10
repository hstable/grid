const Base64 = require('js-base64').Base64
function getUserInfoFromToken(token = '') {
  if (!token) {
    token = window.localStorage.getItem('token')
  }
  return JSON.parse(decodeURIComponent(Base64.decode(token.split(/\./)[1])))
}

export { getUserInfoFromToken }
