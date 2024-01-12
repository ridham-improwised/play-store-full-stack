export default function useAllApiMessages() {
  const obj = {
    success: null,
    apiError: null,
    pendingStatus: null,
    categories: null,
    deleteStatus: null,
    updateStatus: null,
    addStatus: null,
    appDetails: null,
    appReviews: null,
    reviewStatus : null
  };

  const messages = useState("apiResultsData", () => obj);

  const setEmptyData = (pendingValue) => {
    messages.value.success = null;
    messages.value.pendingStatus = pendingValue;
  }
  
  const messageRemover = () => {
    messages.value.addStatus = null;
    messages.value.updateStatus = null;
    messages.value.deleteStatus = null;
    messages.value.apiError = null;
    messages.value.reviewStatus = null
  }

  return {messages, setEmptyData, messageRemover}
}
