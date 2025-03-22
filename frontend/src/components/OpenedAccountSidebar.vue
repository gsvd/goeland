<script lang="ts" setup>
import { useStore } from '../store'

const store = useStore()
</script>

<template>
  <aside class="min-w-[300px] flex flex-col py-3 text-base0 bg-base02 shadow-base04 border-r-1 border-base04 px-3 gap-y-2">
    <div class="flex flex-row gap-3 bg-base04 p-2 rounded-md text-sm cursor-pointer hover:bg-base05 transition">
      <div class="w-12 h-12 rounded-full bg-base02 flex items-center justify-center">PP</div>
      <div class="flex flex-1 flex-col justify-around">
        <div class="flex flex-row justify-between items-center">
            <div>RandomDude</div>
            <div class="text-xs text-base0">3:15pm</div>
        </div>
        <div class="text-xs">Me: Hello lol</div>
      </div>
    </div>
  </aside>
</template>
