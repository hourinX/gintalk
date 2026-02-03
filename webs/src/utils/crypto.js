// 生成随机 AES key
export function generateAESKey() {
    const key = window.crypto.getRandomValues(new Uint8Array(32)); // 256-bit
    return key;
}
  
// AES-GCM 加密
export async function encryptWithAES(aesKey, data) {
    const iv = window.crypto.getRandomValues(new Uint8Array(12)); // 96-bit
    const alg = { name: "AES-GCM", iv };
    const key = await window.crypto.subtle.importKey(
      "raw",
      aesKey,
      alg,
      false,
      ["encrypt"]
    );
    const encoded = new TextEncoder().encode(data);
    const cipher = await window.crypto.subtle.encrypt(alg, key, encoded);
    return { cipher: new Uint8Array(cipher), iv };
}
  
// RSA-OAEP 加密 AES Key
export async function encryptAESKeyWithRSA(publicKeyPem, aesKey) {
    const pemHeader = "-----BEGIN PUBLIC KEY-----";
    const pemFooter = "-----END PUBLIC KEY-----";
    const pemContents = publicKeyPem
      .replace(pemHeader, "")
      .replace(pemFooter, "")
      .replace(/\s+/g, "");
    const binaryDer = Uint8Array.from(atob(pemContents), (c) => c.charCodeAt(0));
    const key = await window.crypto.subtle.importKey(
      "spki",
      binaryDer.buffer,
      { name: "RSA-OAEP", hash: "SHA-256" },
      false,
      ["encrypt"]
    );
    const encrypted = await window.crypto.subtle.encrypt(
      { name: "RSA-OAEP" },
      key,
      aesKey
    );
    return new Uint8Array(encrypted);
}
  
// 封装登录加密
export async function encryptPassword(publicKeyPem, password) {
    const aesKey = generateAESKey();
    const { cipher: encryptedPassword, iv } = await encryptWithAES(aesKey, password);
    const encryptedKey = await encryptAESKeyWithRSA(publicKeyPem, aesKey);
  
    return {
      encryptedPassword: btoa(String.fromCharCode(...encryptedPassword)),
      iv: btoa(String.fromCharCode(...iv)),
      encryptedKey: btoa(String.fromCharCode(...encryptedKey))
    };
}
  