import _pack from './_pack'

export default function({ $axios }) {
  const pack = _pack({ $axios })

  function getTreeLines() {
    return pack({
      url: apiRoot + '/tree/lines',
      method: 'get'
    })
  }
  function getTreeTowers(lineID) {
    return pack({
      url: apiRoot + `/tree/line/${lineID}/towers`,
      method: 'get'
    })
  }

  return { getTreeLines, getTreeTowers }
}
