import _pack from './_pack'

export default function({ $axios }) {
  const pack = _pack({ $axios })

  function getMarksRisky(page, pageSize) {
    return pack({
      url: apiRoot + `/marks/risky`,
      method: 'get',
      params: { page, pageSize }
    })
  }

  return { getMarksRisky }
}
