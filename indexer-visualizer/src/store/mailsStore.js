import { defineStore } from 'pinia';

export const useMailsStore = defineStore('mails', {
  state: () => ({
    mails: [],
    selectedMail: undefined,
    searchTerm: '',
    from: 0,
    totalMails: 0,
  }),
  actions: {
    selectMail(mail) {
      this.selectedMail = mail;
    },
    loadMails(mails) {
      this.mails = mails;
    },
    setSearchTerm(text) {
      this.searchTerm = text;
    },
    setFrom(value) {
      this.from = value;
    },
    setTotalMails(value) {
      this.totalMails = value;
    },
  },
});
