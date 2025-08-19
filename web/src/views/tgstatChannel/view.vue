<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="tgstat频道详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                标题
              </template>
              {{ formValue.title }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                频道人数
              </template>
              {{ formValue.subscribers }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                点赞数
              </template>
              {{ formValue.postReach }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                索引数
              </template>
              {{ formValue.citationIndex }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                类型
              </template>
              {{ formValue.type }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                头像url
              </template>
              {{ formValue.avatar }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                telegram地址
              </template>
              {{ formValue.telegramLink }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/tgstatChannel';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';
  import { getFileExt } from '@/utils/urlUtils';

  const message = useMessage();

  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const dialogWidth = computed(() => {
    return adaModalWidth(580);
  });
  const fileAvatarCSS = computed(() => {
    return {
      '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
      '--n-font-size': `18px`,
    };
  });

  // 下载
  function download(url: string) {
    window.open(url);
  }

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less" scoped></style>