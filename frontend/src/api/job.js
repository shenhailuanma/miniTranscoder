/* eslint-disable quotes */
import request from '@/utils/request'

export const apiGetJobList = (size, page) => {
  let url = `/api/jobs?size=${size}&page=${page}`;
  return request({
    url: url,
    method: 'get',
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

export const apiGetJobsCount = () => {
  let url = `/api/jobs/count`;
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
