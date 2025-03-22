<script setup lang="ts">
import { useStore } from '../store'

const store = useStore()
</script>

<template>
  <div v-if="store.isLoading" class="top-0 left-0 flex flex-col fixed z-50 min-h-screen min-w-screen justify-center items-center gap-y-4 bg-black opacity-75 text-white">
    <div class="spinner w-16 h-16 border-2 border-[#f3f3f3] border-t-[#2c3e50] rounded-[50%]"></div>
    <div class="text-white text-xl">Catching the wind...</div>
  </div>
</template>

<style scoped>
.spinner {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style> 