export default function useCallApi() {
  const { convert } = convertData();
  const { messages, setEmptyData } = allApiMessages();
  const config = useRuntimeConfig();
  const baseUrl = config.public.base_url;
  const randomImageAccessKey = config.public.random_img_access_key;

  // fetch-data
  const fetchData = async (url, options, types = null) => {
    setEmptyData(true)
    const { data, error, pending } = await useFetch(url, options);

    messages.value.apiError = error.value;

    if (types !== "getAppDetail") {
      messages.value.pendingStatus = pending.value;
    }

    if (types === "post") {
      messages.value.addStatus = "App Added Successfully";
    } else if (types === "get") {
      messages.value.success = data.value;
    } else if (types === "delete") {
      messages.value.deleteStatus = "App Deleted Successfully";
    } else if (types === "update") {
      messages.value.updateStatus = "App Updated Successfully";
    } else if (types === "getAppDetail") {
      messages.value.appDetails = data.value
    } 
  };

  // add-app
  const addApp = async (formData) => {
    formData.installs = "0";
    formData.reviews = "0";
    formData.rating = "0";
    const options = {
      baseURL: baseUrl,
      method: "POST",
      body: formData,
      headers: {
        "Content-Type": "application/json",
      },
    };
    fetchData("/apps", options, "post");
  };

  // get-app
  const getApps = async (page, filterQuery, appId) => {
    const options = {
      baseURL: baseUrl,
      query: filterQuery
        ? Object.assign({ page: page }, filterQuery)
        : { page: page },
    };

    fetchData(appId ? `/apps/${appId}` : "/apps", options, "get");
  };
  
  // getSpecificApp
  const getSpecificApp = async (appId) => {
    const options = {
      baseURL: baseUrl
    };
    
    await fetchData(`/apps/${appId}`, options, "getAppDetail")
    messages.value.pendingStatus = null;
  }

  //get-category
  const getCategory = async () => {
    const options = {
      baseURL: baseUrl
    };
    const { data } = await useFetch("/apps/category", options);
    watch(data, () => {
      if (data.value) {
        messages.value.categories = data.value.data.category
      }
    }, {immediate:true});
  };

  //update-app
  const updateApp = async (formData, id) => {
    const options = {
      baseURL: baseUrl,
      method: "PUT",
      body: formData,
      headers: {
        "Content-Type": "application/json",
      },
    };
    fetchData(`/apps/${id}`, options, "update");
  };

  //delete-app
  const deleteApp = async (id) => {
    const deleteOptions = {
      baseURL: baseUrl,
      method: "DELETE",
    };
    await fetchData(`/apps/${id}`, deleteOptions, "delete");
  };

  //randomImage
  const fetchAppsImage = async (query) => {
    const {data} = await useFetch(
      `https://api.pexels.com/v1/search?query=${query}&per_page=1`,
      {
        headers: {
          Authorization: randomImageAccessKey,
        },
      }
    );
    return data.value.photos.length > 0 ?  data.value.photos[0].src.landscape : "https://ui-avatars.com/api/?name=RidhamParmar"
  };

  // fetch specific app details 
  const fetchAppDetails = async (id, appName) => {
    const options = {
      baseURL: baseUrl,
    };
    await fetchData(`/apps/${id}`, options, "getAppDetail")
  
    const {data} = await useFetch(baseUrl + "/reviews", {
      transform: (response) => {
        return response.filter((apps) => {
          if (apps.App === appName)
          return apps;
        });
      },
    })
    messages.value.appReviews = data.value
    messages.value.pendingStatus = null;
  }

  // add review
  const addReview = async (formData) => {
    const convertedData = convert(formData,true)
    const options = {
      baseURL : baseUrl,
      method : "POST",
      body : convertedData,
      headers : {
        "Content-Type" : "application/json",
      },
    }
    const {data, error} = await useFetch('/reviews', options) 
    messages.value.reviewStatus = "Review submitted successfully"
    messages.value.apiError = error.value
  }

  return {
    addApp,
    updateApp,
    getApps,
    getCategory,
    deleteApp,
    fetchAppsImage,
    fetchAppDetails,
    addReview,
    getSpecificApp
  };
}
