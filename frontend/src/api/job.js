/* eslint-disable quotes */
import request from '@/utils/request'

export const apiGetJobList = () => {
  let url = `/api/jobs`;
  return request({
    url: url,
    method: 'get',
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

export const apiCreateJobTranscode = (data) => {
  let url = `/api/job/transcode`;
  return request({
    url: url,
    method: 'post',
    headers: {
      'Content-Type': 'application/json'
    },
    data: data
  })
}
