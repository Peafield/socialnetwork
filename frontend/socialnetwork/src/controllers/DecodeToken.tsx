export function DecodePayload(token: string): any {
  const base64Url = token.split(".")[1]; // Get payload
  const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/"); // Convert Base64Url to Base64

  // Decode base64 string and parse the resulting JSON
  const payload = JSON.parse(
    decodeURIComponent(
      atob(base64)
        .split("")
        .map((c) => {
          return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
        })
        .join("")
    )
  );

  return payload;
}
