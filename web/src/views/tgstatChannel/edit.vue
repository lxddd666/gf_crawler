<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑tgstat频道 #' + formValue.id : '添加tgstat频道'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="标题" path="title">
                  <n-input placeholder="请输入标题" v-model:value="formValue.title" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="频道人数" path="subscribers">
                  <n-input-number placeholder="请输入频道人数" v-model:value="formValue.subscribers" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="点赞数" path="postReach">
                  <n-input placeholder="请输入点赞数" v-model:value="formValue.postReach" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="索引数" path="citationIndex">
                  <n-input-number placeholder="请输入索引数" v-model:value="formValue.citationIndex" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="类型" path="type">
                  <n-input placeholder="请输入类型" v-model:value="formValue.type" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="头像url" path="avatar">
                  <n-input placeholder="请输入头像url" v-model:value="formValue.avatar" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="telegram地址" path="telegramLink">
                  <n-input placeholder="请输入telegram地址" v-model:value="formValue.telegramLink" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { Edit, View } from '@/api/tgstatChannel';
  import { State, newState, rules } from './model';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();

  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });

  // 提交表单
  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        Edit(formValue.value)
          .then((_res) => {
            message.success('操作成功');
            closeForm();
            emit('reloadTable');
          })
          .finally(() => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  // 关闭表单
  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;

    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);

      return;
    }

    // 编辑
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

<style lang="less"></style>