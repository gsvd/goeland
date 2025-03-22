<script lang="ts" setup>
import { reactive } from 'vue'
import { useStore } from '../store'
import { useErrorHandler } from '../composables/useErrorHandler'

const form = reactive({
  address: '',
  password: '',
})

const store = useStore()
const { errorCode, fieldErrors, handleError, clearErrors } = useErrorHandler()

async function login() {
  store.loading = true
  clearErrors()
  try {
    await store.login(form.address, form.password)
  } catch (err: any) {
    handleError(err)
  } finally {
    store.loading = false
  }
}
</script>

<template>
  <div class="flex flex-1 flex-col items-center justify-center py-12">
    <h1 class="text-xl font-medium mb-8">Login</h1>

    <form @submit.prevent="login" class="w-full max-w-sm space-y-4">
      <div>
        <label class="block text-sm font-medium mb-1">XMPP address</label>
        <input
          v-model="form.address"
          type="text"
          required
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
          required
          placeholder="********"
        />
        <div class="min-h-[1rem] mt-0.5">
          <p v-if="fieldErrors.password" class="text-red text-xs">
            {{ $t(`errors.${fieldErrors.password}`) }}
          </p>
        </div>
      </div>

      <input
        type="submit"
        :disabled="store.loading"
        value="Login"
      />
    </form>

    <div class="min-h-[1rem] mt-3">
      <p v-if="errorCode && !Object.keys(fieldErrors).length" class="text-red text-xs text-center">
        {{ $t(`errors.${errorCode}`) }}
      </p>
    </div>
  </div>
</template>