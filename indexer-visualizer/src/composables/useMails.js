import { storeToRefs } from 'pinia';
import { useMailsStore } from '../store/mailsStore';

export const useMails = () => {
  const mailsStore = useMailsStore();
  const { mails, selectedMail, searchTerm, from, totalMails} = storeToRefs(mailsStore);

  const getMails = async () => {
    const res = await fetch(
      `http://localhost:4000/documents/match?term=${searchTerm.value}&from=${from.value}`
    );
    const data = await res.json();
    mailsStore.loadMails(data['items']);
    mailsStore.setTotalMails(data['total'])
  };

  const showMailDetail = (mail) => {
    mailsStore.selectMail(mail);
  };

  const setSearchTerm = (text) => {
    mailsStore.setSearchTerm(text);
  };

  const nextPage = () => {
    mailsStore.setFrom(from.value + 20);
  };

  const previousPage = () => {
    mailsStore.setFrom(from.value - 20);
  };

  return {
    mails,
    selectedMail,
    searchTerm,
    from,
    totalMails,
    getMails,
    showMailDetail,
    setSearchTerm,
    nextPage,
    previousPage,
  };
};
