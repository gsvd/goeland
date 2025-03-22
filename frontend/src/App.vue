<script lang="ts" setup>
import { onMounted } from 'vue'
import { useStore } from './store'
import Auth from './components/Auth.vue'
import Loading from './components/Loading.vue'
import Sidebar from './components/Sidebar.vue'
import { OpenAllAccounts } from '../wailsjs/go/main/App'
import OpenedAccount from './components/OpenedAccount.vue'

const store = useStore()

onMounted(() => {
  store.getAllAccounts(1000)
  OpenAllAccounts()
})
</script>

<template>
  <div class="flex min-h-screen">
    <Sidebar />

    <main class="flex-1 bg-base02">
      <Loading v-if="store.isLoading" />
      <Auth />
      <OpenedAccount />
    </main>
  </div>
</template>
