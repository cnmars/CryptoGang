import { defineStore } from 'pinia';
import { reactive,ref } from 'vue';
import { Wallet } from '@/models/singleWallet';
import { ethers } from 'ethers';
import { useUserStore } from './userStore';

export const useWalletStore = defineStore('wallet', () => {
  const walletTable = reactive([])
  return {
    walletTable,
  }
})
