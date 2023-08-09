const API_URL = "http://localhost:8080";

export async function handleAPIRequest(url, options) {
  try {
    const response = await fetch(API_URL + url, options);
    if (response.ok) {
      const data = await response.json();
      return data;
    } else {
      throw new Error(response.status || "Oops, there was a problem");
    }
  } catch (error) {
    console.error(error);
    throw new Error("API request failed");
  }
}
