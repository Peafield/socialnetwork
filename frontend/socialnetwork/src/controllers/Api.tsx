import { useSetUserContextAndCookie } from "./SetUserContextAndCookie";

const API_URL = "http://0.0.0.0:8080";
//const API_URL = 'http://backend-image:8080';

export async function handleAPIRequest(url: string, options: object) {
  try {
    const response = await fetch(API_URL + url, options);

    if (response.ok) {
      const data = await response.json();
      return data;
    } else {
      throw new Error(
        response.status
          ? response.status.toString()
          : "Oops, there was a problem"
      );
    }
  } catch (error) {
    if (error instanceof Error) {
      throw new Error("API request failed", {
        cause: error.message
      });
    } else {
      console.error(error);
      throw new Error("API request failed");
    }
  }
}
