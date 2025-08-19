import { http, jumpExport } from '@/utils/http/axios';

// 获取tgstat频道列表
export function List(params) {
  return http.request({
    url: '/tgstatChannel/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除tgstat频道
export function Delete(params) {
  return http.request({
    url: '/tgstatChannel/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑tgstat频道
export function Edit(params) {
  return http.request({
    url: '/tgstatChannel/edit',
    method: 'POST',
    params,
  });
}

// 获取tgstat频道指定详情
export function View(params) {
  return http.request({
    url: '/tgstatChannel/view',
    method: 'GET',
    params,
  });
}

// 导出tgstat频道
export function Export(params) {
  jumpExport('/tgstatChannel/export', params);
}