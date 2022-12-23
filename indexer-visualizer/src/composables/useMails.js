import { storeToRefs } from 'pinia';
import { useMailsStore } from '../store/mailsStore';

export const useMails = () => {
  const mailsStore = useMailsStore();
  const { mails, selectedMail, searchTerm, from } = storeToRefs(mailsStore);

  const getMails = async () => {
    const res = await fetch(
      `http://localhost:4000/documents/match?term=${searchTerm.value}&from=${from.value}`
    );
    const data = await res.json();
    mailsStore.loadMails(data['hits']['hits']);
  };

  const showMailDetail = (mail) => {
    mailsStore.selectMail(mail);
  };

  const setSearchTerm = (text) => {
    mailsStore.setSearchTerm(text);
  };

  return { mails, selectedMail, searchTerm, from, getMails, showMailDetail, setSearchTerm };
};
