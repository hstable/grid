import { ToastProgrammatic } from 'buefy'

function handleResponse(res, that, suc, err) {
  if (!res.data) {
    if (err && err instanceof Function) {
      err.apply(that)
    }
    return
  }
  if (res.data.code === 'SUCCESS') {
    suc.apply(that)
  } else if (err && err instanceof Function) {
    err.apply(that)
  } else {
    ToastProgrammatic.open({
      message: res.data.message,
      type: 'is-warning',
      position: 'is-top',
      queue: false,
      duration: 5000
    })
  }
}

/*
var myURL = parseURL('http://abc.com:8080/dir/index.html?id=255&m=hello#top');
myURL.file;     // = 'index.html'
myURL.hash;     // = 'top'
myURL.host;     // = 'abc.com'
myURL.query;    // = '?id=255&m=hello'
myURL.params;   // = Object = { id: 255, m: hello }
myURL.path;     // = '/dir/index.html'
myURL.segments; // = Array = ['dir', 'index.html']
myURL.port;     // = '8080'
myURL.protocol; // = 'http'
myURL.source;   // = 'http://abc.com:8080/dir/index.html?id=255&m=hello#top'
*/
function parseURL(url) {
  if (!url.includes('://')) {
    url = 'http://' + url
  }
  const a = document.createElement('a')
  a.href = url
  return {
    source: url,
    protocol: a.protocol.replace(':', ''),
    host: a.hostname,
    port: a.port,
    query: a.search,
    params: (function() {
      const ret = {}
      const seg = a.search.replace(/^\?/, '').split('&')
      const len = seg.length
      let i = 0
      let s
      for (; i < len; i++) {
        if (!seg[i]) {
          continue
        }
        s = seg[i].split('=')
        ret[s[0]] = s[1]
      }
      return ret
    })(),
    file: (a.pathname.match(/\/([^/?#]+)$/i) || [undefined, ''])[1],
    hash: a.hash.replace('#', ''),
    path: a.pathname.replace(/^([^/])/, '/$1'),
    relative: (a.href.match(/tps?:\/\/[^/]+(.+)/) || [undefined, ''])[1],
    segments: a.pathname.replace(/^\//, '').split('/')
  }
}

/* 判断一个IPv4的地址是否是内网地址 */
function isIntranet(url) {
  const host = parseURL(url).host
  const arr = host.split('.')
  if (arr.length !== 4) {
    return host === 'localhost' || host === 'local'
  }
  if (arr.some((p) => parseInt(p) < 0 || parseInt(p) > 255)) {
    // 每一位必须是[0,255]
    return false
  }
  let bin = '' // 传入的IP的二进制表示
  arr.forEach((p) => {
    const t = parseInt(p).toString(2)
    bin += '0'.repeat(8 - t.length) + t
  })
  const list = [
    '0.0.0.0/32',
    '10.0.0.0/8',
    '127.0.0.0/8',
    '169.254.0.0/16',
    '172.16.0.0/12',
    '192.168.0.0/16',
    '224.0.0.0/4',
    '240.0.0.0/4',
    '255.255.255.255/32'
  ]
  return list.some((mask) => {
    let arr = mask.split('/')
    const prefix = arr[0]
    let suffix = arr[1]
    arr = prefix.split('.')
    let b = ''
    arr.forEach((p) => {
      const t = parseInt(p).toString(2)
      b += '0'.repeat(8 - t.length) + t
    })
    suffix = parseInt(suffix)
    let is = true
    for (let i = 0; i < suffix; i++) {
      if (b[i] !== bin[i]) {
        is = false
        break
      }
    }
    return is
  })
}

function isVersionGreaterEqual(va, vb) {
  if (va.toLowerCase() === 'debug') {
    return true
  }
  va = va.trim()
  vb = vb.trim()
  if (va.length > 0 && va[0] === 'v') {
    va = va.substr(1)
  }
  if (vb.length > 0 && vb[0] === 'v') {
    vb = vb.substr(1)
  }
  va.replace('-', '.')
  vb.replace('-', '.')
  const a = va.split('.')
  const b = vb.split('.')
  const minlen = Math.min(a.length, b.length)
  for (let i = 0; i < minlen; i++) {
    if (parseInt(a[i]) < parseInt(b[i])) {
      return false
    }
    if (parseInt(a[i]) > parseInt(b[i])) {
      return true
    }
  }
  return a.length >= b.length
}

function toInt(s) {
  if (typeof s === 'string') {
    return parseInt(s)
  } else if (typeof s === 'number') {
    return parseInt(s)
  } else if (typeof s === 'boolean') {
    return s ? 1 : 0
  }
  return s
}

export { handleResponse, parseURL, isIntranet, isVersionGreaterEqual, toInt }
