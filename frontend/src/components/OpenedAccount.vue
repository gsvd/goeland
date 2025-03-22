<script lang="ts" setup>
import { useStore } from '../store'
import OpenedAccountSidebar from './OpenedAccountSidebar.vue'

const store = useStore()
</script>

<template>
  <div v-if="store.getActiveAccount && !store.uiState.showAuth" class="flex min-h-screen flex-row text-base0">
    <OpenedAccountSidebar />
  </div>
</template>
