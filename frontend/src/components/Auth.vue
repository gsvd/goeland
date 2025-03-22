<script lang="ts" setup>
import { onDeactivated, onMounted } from 'vue'
import { useStore } from '../store'
import AuthLogin from './AuthLogin.vue'
import AuthJoin from './AuthJoin.vue'

import qcFlag from '../assets/images/flags/qc.png'

const store = useStore()

onMounted(() => {
  store.setShowAuth(true)
})

onDeactivated(() => {
  store.setShowAuth(false)
})
</script>

<template>
  <div
    v-if="store.uiState.showAuth"
    class="flex flex-col text-base0 min-h-screen"
  >
    <div class="flex flex-col items-center justify-center py-4">
      <h1 class="text-2xl font-bold text-green">Goeland.im</h1>
    </div>

    <div class="flex flex-row justify-center items-center divide-x-2 divide-base04 divide-dashed flex-1">
      <AuthLogin />
      <AuthJoin />
    </div>

    <div class="flex flex-row justify-center items-center gap-2 py-2 text-sm">
      <p>Built in Montréal, Meant for all</p>
      <img :src="qcFlag" class="inline-block w-6 ml-1" alt="Quebec flag" title="Quebec flag" />
    </div>
  </div>
</template>

