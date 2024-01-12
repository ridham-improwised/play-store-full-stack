export default function useConvertData() {
  const getDate = () => {
    const days = [
      "Sunday",
      "Monday",
      "Tuesday",
      "Wednesday",
      "Thursday",
      "Friday",
      "Saturday",
    ];
    const ans = new Date();
    return `${days[ans.getDay()]} ${ans.getDate()}, ${ans.getFullYear()}`;
  };

  let pattern = /[A-Z]\B/;
  
  const convert = (formData, addingReview = false) => {
    const converted = {};
    for (let key in formData) {
      let index = key.search(pattern)
      if (index !== -1 && !addingReview) {
        let firstWord = key.slice(0, index)
        let secondWord = key.slice(index) 
        let combined = firstWord.charAt(0).toUpperCase() + firstWord.slice(1) + " " 
        combined = secondWord === "Version" ? combined + secondWord.slice(0,3) : combined + secondWord.slice(0)
        converted[combined] = formData[key]
      } else {
        converted[key.charAt(0).toUpperCase() + key.slice(1)] = formData[key];
      }
    }
    if (!addingReview) {
      converted["Last Updated"] = getDate();
    }
    return converted;
  };

  return {
    convert,
  };
}
