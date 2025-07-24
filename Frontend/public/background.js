console.log("background js");

chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  if (message === "LOGIN_WITH_GOOGLE") {
    console.log("go the message");

    chrome.identity.getAuthToken({ interactive: true }, (token) => {
      if (chrome.runtime.lastError) {
        console.error("Auth error:", chrome.runtime.lastError.message);
        sendResponse({ error: chrome.runtime.lastError.message });
      } else {
        console.log("Token received:", token);
        sendResponse({ token });
      }
    });

    return true; // ✅ keeps the message channel alive
  }
});
