import { createI18n } from 'vue-i18n'

const messages = {
  en: {
    errors: {
      ERR_EMPTY_ADDRESS: "Address is required.",
      ERR_INVALID_ADDRESS: "Address must contain '@'.",
      ERR_PASSWORD_REQUIRED: "Password is required.",
      ERR_ACCOUNT_EXISTS: "An account with this address is already connected.",
      ERR_AUTHENTICATION_FAILED: "Authentication failed. Please check your credentials.",

      ERR_INVALID_INPUT: "Invalid form input.",
      ERR_INVALID_ADDRESS_FORMAT: "Invalid XMPP address format.",

      ERR_UNKNOWN: "An unknown error occurred."
    },
    account: {
      added_success: "Account added successfully!"
    }
  }
}

const i18n = createI18n({
  legacy: false,
  locale: 'en',
  fallbackLocale: 'en',
  messages,
})

export default i18n
