const API_URL = "http://localhost:8080";

export async function handleAPIRequest(url: string, options: object) {
  try {
    const response = await fetch(API_URL + url, options);
    if (response.ok) {      
      const contentType = response.headers.get("content-type");
      
      if (contentType && contentType.indexOf("application/json") !== -1) {
        return await response.json();
      } else {
        return {};
      }
      
    } else {
      throw new Error(
        response.status
          ? response.status.toString()
          : "Oops, there was a problem"
      );
    }
  } catch (error) {
    console.error(error);
    throw new Error("API request failed");
  }
}


