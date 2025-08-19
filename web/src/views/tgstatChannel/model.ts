import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // id
  public title = ''; // 标题
  public subscribers = 0; // 频道人数
  public postReach = 0; // 点赞数
  public citationIndex = 0; // 索引数
  public type = ''; // 类型
  public avatar = ''; // 头像url
  public telegramLink = ''; // telegram地址
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 修改时间
  public deletedAt = ''; // deleted_at

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则
export const rules = {
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入标题',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'id',
    componentProps: {
      placeholder: '请输入id',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '标题',
    key: 'title',
    align: 'left',
    width: -1,
  },
  {
    title: '频道人数',
    key: 'subscribers',
    align: 'left',
    width: -1,
  },
  {
    title: '点赞数',
    key: 'postReach',
    align: 'left',
    width: -1,
  },
  {
    title: '索引数',
    key: 'citationIndex',
    align: 'left',
    width: -1,
  },
  {
    title: '类型',
    key: 'type',
    align: 'left',
    width: -1,
  },
  {
    title: '头像url',
    key: 'avatar',
    align: 'left',
    width: -1,
  },
  {
    title: 'telegram地址',
    key: 'telegramLink',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
  {
    title: '修改时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
];