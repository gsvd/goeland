import { defineStore } from "pinia";
import { store } from "../../wailsjs/go/models";
import { GetAllAccounts, AddAccount, OpenAccount } from "../../wailsjs/go/main/App.js";

export const useStore = defineStore("app", {
  state: () => ({
    accounts: [] as store.Account[],
    loading: false,
    error: null as string | null,
    activeAccount: null as store.Account | null,
    uiState: {
      showAuth: false,
    },
  }),

  getters: {
    getActiveAccount(state): store.Account | null {
      return state.activeAccount;
    },
    getActiveAccountAddress(state): string | null {
      return state.activeAccount ? state.activeAccount.Address : null;
    },
    getAccounts(state): store.Account[] {
      return state.accounts;
    },
    isLoading(state): boolean {
      return state.loading;
    },
    getError(state): string | null {
      return state.error;
    }
  },

  actions: {
    async getAllAccounts(timeout: number = 2500) {
      this.loading = true;
      try {
        const raw = await GetAllAccounts();
        if (!Array.isArray(raw)) {
          throw new Error("Invalid response from GetAllAccounts");
        }

        const accounts = raw.map((acc: any) => store.Account.createFrom(acc));
        this.accounts = accounts;
        if (accounts.length > 0 && this.activeAccount === null) {
          this.setActiveAccount(accounts[0]);
        }
      } catch (err) {
        console.error("Failed to load accounts:", err);
      } finally {
        setTimeout(() => {
          this.loading = false;
        }, timeout);
      }
    },

    async login(address: string, password: string) {
      const added = await AddAccount(address, password);
      const created = store.Account.createFrom(added);
      
      this.accounts.push(created);
    },

    async setActiveAccount(account: store.Account | null) {
      this.activeAccount = account;
      this.setShowAuth(false);
    },

    setShowAuth(open: boolean) {
      this.uiState.showAuth = open;
    },

    setError(error: string | null) {
      this.error = error;
    },
  },
});
