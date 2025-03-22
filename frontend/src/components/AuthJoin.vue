<script lang="ts" setup>
import { reactive } from 'vue'
import { useStore } from '../store'
import { useErrorHandler } from '../composables/useErrorHandler'

const disabled = true

const form = reactive({
  address: '',
  password: '',
  confirmPassword: '',
})

const store = useStore()
const { errorCode, fieldErrors, handleError, clearErrors } = useErrorHandler()

async function join() {}
</script>

<template>
  <div class="flex flex-1 flex-col items-center justify-center py-12">
    <h1 class="text-xl font-medium mb-8">Join</h1>

    <form @submit.prevent="join" class="w-full max-w-sm space-y-4">
      <div>
        <label class="block text-sm font-medium mb-1">XMPP address</label>
        <input
          v-model="form.address"
          type="text"
          :disabled="disabled"
          placeholder="gsvd@goeland.im"
        />
        <div class="min-h-[1rem] mt-0.5">
          <p v-if="fieldErrors.address" class="text-red text-xs">
            {{ $t(`errors.${fieldErrors.address}`) }}
          </p>
        </div>
      </div>

      <div>
        <label class="block text-sm font-medium mb-1">Password</label>
        <input
          v-model="form.password"
          type="password"
          :disabled="disabled"
          placeholder="********"
        />
        <div class="min-h-[1rem] mt-0.5">
          <p v-if="fieldErrors.password" class="text-red text-xs">
            {{ $t(`errors.${fieldErrors.password}`) }}
          </p>
        </div>
      </div>

      <div>
        <label class="block text-sm font-medium mb-1">Confirm Password</label>
        <input
          v-model="form.confirmPassword"
          type="password"
          class=""
          :disabled="disabled"
          placeholder="********"
        />
        <div class="min-h-[1rem] mt-0.5">
          <p v-if="fieldErrors.confirmPassword" class="text-red text-xs">
            {{ $t(`errors.${fieldErrors.confirmPassword}`) }}
          </p>
        </div>
      </div>

      <input
        type="submit"
        :disabled="disabled || store.loading"
        value="Join"
      />
    </form>

    <div class="min-h-[1rem] mt-3">
      <p v-if="disabled" class="text-red text-xs text-center">
        Not available yet.
      </p>
      <p v-else-if="errorCode && !Object.keys(fieldErrors).length" class="text-red text-xs text-center">
        {{ $t(`errors.${errorCode}`) }}
      </p>
    </div>
  </div>
</template>
