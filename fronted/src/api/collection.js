import request from '@/utils/request'

export function fetchCollection() {
  return request({
    url: '/collection',
    method: 'get'
  })
}

export function cudCollection(data) {
  return request({
    url: `/collection`,
    method: 'post',
    data
  })
}
