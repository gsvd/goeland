import { ref, reactive } from 'vue'

type FieldErrors = Record<string, string>

type APIError = {
  code: string
  errors?: { field: string; code: string }[]
}

export function useErrorHandler() {
  const errorCode = ref<string | null>(null)
  const fieldErrors = reactive<FieldErrors>({})

  function clearErrors() {
    errorCode.value = null
    Object.keys(fieldErrors).forEach((key) => delete fieldErrors[key])
  }

  function handleError(err: unknown) {
    clearErrors()

    if (typeof err === 'string') {
      try {
        const parsed: APIError = JSON.parse(err)
        errorCode.value = parsed.code ?? 'ERR_UNKNOWN'
        parsed.errors?.forEach(({ field, code }) => {
          if (field && code) fieldErrors[field] = code
        })
      } catch {
        errorCode.value = err
      }
      return
    }

    if (typeof err === 'object' && err !== null) {
      errorCode.value = (err as any).code ?? (err as any).message ?? 'ERR_UNKNOWN'
    } else {
      errorCode.value = 'ERR_UNKNOWN'
    }
  }

  return {
    errorCode,
    fieldErrors,
    handleError,
    clearErrors,
  }
}
