export const createImageURL = (imageData: string) => {
    const decodedAvatar = atob(imageData); // Decode base64-encoded avatar data
    const avatarBuffer = new ArrayBuffer(decodedAvatar.length);
    const avatarView = new Uint8Array(avatarBuffer);
    for (let i = 0; i < decodedAvatar.length; i++) {
        avatarView[i] = decodedAvatar.charCodeAt(i);
    }

    const blob = new Blob([avatarBuffer]);
    const url = URL.createObjectURL(blob);
    return url
}